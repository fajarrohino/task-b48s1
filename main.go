package main

import (
	"context"
	"day10/connection"
	"day10/middleware"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)
type DataProject struct{
	    Id              int
        ProjectName     string
        StartDate       time.Time
        EndDate         time.Time
        Duration        string
        Description     string
        Image           string
        NodeJs          bool
        ReactJs         bool
        NextJs          bool
        TypoScript      bool
        Author          string
}

type User struct{
    Id int
    Name string
    Email string
    HashedPassword string
}

type SessionData struct{
    IsLogin bool
    Name string
}

// slice of struc ->mirip array of objec
var sessionUser = SessionData{}

// var dataUsers = []User{}
var Projects = []DataProject{}
// routes
func main() {
    e := echo.New()
    // session
    e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))
    // connection database
    connection.DatabaseConnect()

    // aset
    e.Static("/aset", "aset")
    // aset
    e.Static("/uploads", "uploads")

    // Routes Get
    e.GET("/", home)
    e.GET("/form-project", formproject)
    e.GET("/testimonial", testimonial)
    e.GET("/contact", contact)
    e.GET("/listproject", project)
    e.GET("/project-detail/:id", projectdetail)
    e.GET("/project-edit/:id", projectedit)

    // auth
    e.GET("/form-register", formregister)
    e.GET("/form-login", formlogin)
    e.POST("/login", login)
    e.POST("/register", register)
    e.POST("/logout", logout)

    // Post
    e.POST("/", middleware.UploadFile(addproject))
    // e.POST("/", addproject)
    e.POST("/project-delete/:id", deleteProject)
    e.POST("/project-edit/:Id", middleware.UploadFile(submitEditProject))
    // e.POST("/project-edit/:Id", submitEditProject)
   

    e.Logger.Fatal(e.Start("localhost:5000"))
}

// handlers

    // GET("/", home)
func home(c echo.Context) error{
    
    // utk mengambil data di database
    Querys,errQuery:= connection.Conn.Query(context.Background(), "SELECT tb_project.id, tb_user.name AS author, tb_project.project_name, tb_project.start_date, tb_project.end_date, tb_project.description, tb_project.image, tb_project.nodejs, tb_project.reactjs, tb_project.nextjs, tb_project.typoscript, tb_project.duration FROM tb_project LEFT JOIN tb_user ON tb_project.author_id = tb_user.id")
    if errQuery != nil {
        return c.JSON(http.StatusInternalServerError, errQuery.Error())
    }
    var inputProject []DataProject
    // var inputUser []User
    for Querys.Next(){
        var each = DataProject{}
        // var data = User{}
        // var tempAuthor sql.NullString
        err:= Querys.Scan(&each.Id, &each.Author, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Image, &each.NodeJs, &each.ReactJs, &each.NextJs, &each.TypoScript, &each.Duration)

        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"home" : err.Error()})
        }
        inputProject = append(inputProject, each)
        // inputUser = append(inputUser, data)
    }
    //session
    sess, _:= session.Get("session", c)
    
    if sess.Values["isLogin"] != true {
        sessionUser.IsLogin = false
    } else {
        sessionUser.IsLogin = sess.Values["isLogin"].(bool)
        sessionUser.Name = sess.Values["name"].(string)
    }
    // utk pemanggilan di html index
    myProject:=map[string]interface{}{
        "DataSessUser" :sessionUser,
        "myproject" : inputProject,
        // "myauthor" : inputUser,
        "MessageFlash": sess.Values["message"], // regis berhasil
        "StatusFlash": sess.Values["status"], // true
    }
    delete(sess.Values, "message")
    delete(sess.Values, "status")
    // save
    errSess:= sess.Save(c.Request(), c.Response())
    if errSess != nil{
        return c.JSON(http.StatusInternalServerError, map[string]string{"sess": errSess.Error()})
    }
    tmpl, err := template.ParseFiles("views/index.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    return tmpl.Execute(c.Response(),myProject)
}
    // GET("/form-project", formproject)
func formproject(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/add-project.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    // session
    sess, _:= session.Get("session", c)
    
    if sess.Values["isLogin"] != true {
        sessionUser.IsLogin = false
    } else {
        sessionUser.IsLogin = sess.Values["isLogin"].(bool)
        sessionUser.Name = sess.Values["name"].(string)
    }
    data := map[string]interface{}{
		"DataSessUser" :sessionUser,
	}
    return tmpl.Execute(c.Response(),data)
}
    // GET("/testimonial", testimonial)
func testimonial(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/testimonial.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    // session
    sess, _:= session.Get("session", c)
    
    if sess.Values["isLogin"] != true {
        sessionUser.IsLogin = false
    } else {
        sessionUser.IsLogin = sess.Values["isLogin"].(bool)
        sessionUser.Name = sess.Values["name"].(string)
    }
    data := map[string]interface{}{
		"DataSessUser" :sessionUser,
	}
    return tmpl.Execute(c.Response(),data)
}
    // GET("/contact", contact)
func contact(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/contact.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"Contact" :err.Error()})
    }
    //session
    sess, _:= session.Get("session", c)
    
    if sess.Values["isLogin"] != true {
        sessionUser.IsLogin = false
    } else {
        sessionUser.IsLogin = sess.Values["isLogin"].(bool)
        sessionUser.Name = sess.Values["name"].(string)
    }
    data := map[string]interface{}{
		"DataSessUser" :sessionUser,
	}
    return tmpl.Execute(c.Response(),data)
}
    // GET("/list-project", project)
func project(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/list-project.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    // utk mengambil data di database
    Querys,errQuery:= connection.Conn.Query(context.Background(),"SELECT tb_project.id, tb_user.name AS author, tb_project.project_name, tb_project.start_date, tb_project.end_date, tb_project.description, tb_project.image, tb_project.nodejs, tb_project.reactjs, tb_project.nextjs, tb_project.typoscript, tb_project.duration FROM tb_project LEFT JOIN tb_user ON tb_project.author_id = tb_user.id")

    if errQuery != nil {
        return c.JSON(http.StatusInternalServerError, errQuery.Error())
    }
    var inputProject []DataProject
    for Querys.Next(){
        var each = DataProject{}
        err:= Querys.Scan(&each.Id, &each.Author, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Image, &each.NodeJs, &each.ReactJs, &each.NextJs, &each.TypoScript, &each.Duration)

        if err != nil {
            return c.JSON(http.StatusInternalServerError,err.Error())
        }
        inputProject = append(inputProject, each)
    }
    // session
    sess, _:= session.Get("session", c)
    
    if sess.Values["isLogin"] != true {
        sessionUser.IsLogin = false
    } else {
        sessionUser.IsLogin = sess.Values["isLogin"].(bool)
        sessionUser.Name = sess.Values["name"].(string)
    }
    // utk pemanggilan di html index
    myProject:=map[string]interface{}{
        "myproject" : inputProject,
        "DataSessUser" :sessionUser,
    }
    delete(sess.Values, "message")
    delete(sess.Values, "status")
    // save
    errSess:= sess.Save(c.Request(), c.Response())
    if errSess != nil{
        return c.JSON(http.StatusInternalServerError, map[string]string{"sess": errSess.Error()})
    }
    return tmpl.Execute(c.Response(),myProject)
}   
    // GET("/project-detail/:id", projectdetail)
func projectdetail(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/project-detail.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    id,_:= strconv.Atoi(c.Param("id"))

    projectdetail := DataProject{}
    
    errQuery :=connection.Conn.QueryRow(context.Background(), "SELECT tb_project.id, tb_user.name AS author, tb_project.project_name, tb_project.start_date, tb_project.end_date, tb_project.description, tb_project.image, tb_project.nodejs, tb_project.reactjs, tb_project.nextjs, tb_project.typoscript, tb_project.duration FROM tb_project LEFT JOIN tb_user ON tb_project.author_id = tb_user.id WHERE tb_project.id=$1",id).Scan(&projectdetail.Id, &projectdetail.Author, &projectdetail.ProjectName, &projectdetail.StartDate, &projectdetail.EndDate, &projectdetail.Description, &projectdetail.Image, &projectdetail.NodeJs, &projectdetail.ReactJs, &projectdetail.NextJs, &projectdetail.TypoScript, &projectdetail.Duration)

    if errQuery != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"dimana error" :errQuery.Error()})
    }
    data :=map[string]interface{}{
        "Id" :id,
        "dp":   projectdetail,
        // dp ->DataProject
    }
    return tmpl.Execute(c.Response(),data)
}
// GET("/project-edit/:id", projectedit)
func projectedit(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/edit-project.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,map[string]string{"disini" : err.Error()})
    }
    id, _ := strconv.Atoi(c.Param("id"))

    editDp := DataProject{}

    errEDP :=connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE Id=$1",id).Scan(&editDp.Id, &editDp.ProjectName, &editDp.StartDate, &editDp.EndDate, &editDp.Description, &editDp.Image, &editDp.NodeJs, &editDp.ReactJs, &editDp.NextJs, &editDp.TypoScript, &editDp.Duration, &editDp.Author)
    if errEDP != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error dimana" : errEDP.Error()})
    }
    sess, _:= session.Get("session", c)
    
    if sess.Values["isLogin"] != true {
        sessionUser.IsLogin = false
    } else {
        sessionUser.IsLogin = sess.Values["isLogin"].(bool)
        sessionUser.Name = sess.Values["name"].(string)
    }
    edit :=map[string]interface{}{
        "Id"  : id,
        "data":   editDp,
        "DataSessUser" :sessionUser,
        // dp ->DataProject
    }
    return tmpl.Execute(c.Response(),edit)
}

    // POST("/project-delete/:id", deleteProject)
func deleteProject(c echo.Context) error  {
    id, _:= strconv.Atoi(c.Param("id"))

    _, errDelete:=connection.Conn.Exec(context.Background(),"DELETE FROM tb_project WHERE Id=$1", id)
    if errDelete != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"apakah disini":errDelete.Error()})
    }

    // Projects = append(Projects[:id], Projects[id+1:]...)
	return c.Redirect(http.StatusMovedPermanently, "/")
}
    // POST("/project-edit/:id", submitEditProject)
func submitEditProject(c echo.Context) error  {
    id:=c.FormValue("id")
    // author:=c.FormValue("input-author")
    nameProject:=c.FormValue("input-project-name")
    startDate:=c.FormValue("input-start-date")
    endDate:=c.FormValue("input-end-date")
    description:=c.FormValue("input-description")
    // image:=c.FormValue("input-img")
    nodejs:=c.FormValue("input-nodejs")
    reactjs:=c.FormValue("input-reactjs")
    nextjs:=c.FormValue("input-nextjs")
    typoscript:=c.FormValue("input-typoscript")

    image := c.Get("dataFile").(string)
    Duration := coutDuration(startDate, endDate)

    // connection to database
    update, errupdate :=connection.Conn.Exec(context.Background(), "UPDATE tb_project SET project_name=$1, start_date=$2, end_date=$3, description=$4, image=$5, nodejs=$6, reactjs=$7, nextjs=$8, typoscript=$9, duration=$10 WHERE id=$11", nameProject, startDate, endDate, description, image, nodejs != "", reactjs != "", nextjs != "", typoscript != "",Duration, id)
    if errupdate != nil {
        return c.JSON(http.StatusInternalServerError,map[string]string{"update error " :errupdate.Error()})
    }
    fmt.Println("data masuk", update.RowsAffected())
    return c.Redirect(http.StatusMovedPermanently, "/listproject")
}
// POST("/", addproject) => masih error
func addproject(c echo.Context) error  {
    // input data
    nameProject:=c.FormValue("input-project-name")
    startDate:=c.FormValue("input-start-date")
    endDate:=c.FormValue("input-end-date")
    description:=c.FormValue("input-description")
    nodejs:=c.FormValue("input-nodejs")
    reactjs:=c.FormValue("input-reactjs")
    nextjs:=c.FormValue("input-nextjs")
    typoscript:=c.FormValue("input-typoscript")
    // get file
    image:=c.Get("dataFile").(string)
    
    // session
    sess, errSess := session.Get("session", c)
    if errSess != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"di sess add": errSess.Error() })
    }
    // sess.Save(c.Request(), c.Response())
    // if errSess != nil{
    //     return c.JSON(http.StatusInternalServerError, map[string]string{"sess": errSess.Error()})
    // }

    authorId:=sess.Values["Id"].(int)
    // variabel duration
    Duration := coutDuration(startDate,endDate)

    // fmt.Println(nameProject,startDate,endDate,description,image,nodejs,reactjs,nextjs,typoscript)
    // return c.JSON(http.StatusOK, "berhasil")


    // connection database
    _,err:=connection.Conn.Exec(context.Background(),
    "INSERT INTO tb_project (project_name, start_date, end_date, description, image, nodejs, reactjs, nextjs, typoscript,  duration, author_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", nameProject, startDate, endDate, description, image, (nodejs=="nodeJs"), (reactjs=="reactJs"), (nextjs=="nextJs"), (typoscript=="typoscript"), Duration, authorId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"di add":err.Error()} )
    }
    return c.Redirect(http.StatusMovedPermanently,"/listproject")
}

func coutDuration(d1 string, d2 string) string {
    date1, _:= time.Parse("2006-01-2", d1)
    date2, _:= time.Parse("2006-01-2", d2)

    // selisih
    distance :=date2.Sub(date1)
    days := int(distance.Hours()/24)
    weeks := days/7
    months := weeks/30

    if months > 12 {
        return strconv.Itoa(months/12) + "Year"
    }
    if months > 0 {
		return strconv.Itoa(months) + " Month"
	}
	if weeks > 0 {
		return strconv.Itoa(weeks) + " Week"
	}
	return strconv.Itoa(days) + " Day"
}

    // GET("/form-register/", register)
func formregister(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/form-register.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    sess,errSess:=session.Get("session", c)
    if errSess != nil {
        return c.JSON(http.StatusInternalServerError, errSess.Error())
    }
    mFlash := map[string]interface{}{
        "MessageFlash": sess.Values["message"], // regis berhasil
        "StatusFlash": sess.Values["status"],
    }
    delete(sess.Values, "message")
    delete(sess.Values, "status")
    sess.Save(c.Request(), c.Response())

    return tmpl.Execute(c.Response(),mFlash)
}

    // e.POST("/register", register)
func register(c echo.Context) error{
    name := c.FormValue("input-name")
    email := c.FormValue("input-email")
    password := c.FormValue("input-password")

    hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(password),10) //hashedPassword -> merubah password -> acak
    if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
    // check input data 
    fmt.Println(name, email, password)
    regisQuery, regisErr:=connection.Conn.Exec(context.Background(),"INSERT INTO tb_user (name, email, password) VALUES($1, $2, $3)", name, email, hashedPassword)
    fmt.Println("data masuk : ", regisQuery)
    if regisErr != nil {
        return c.JSON(http.StatusInternalServerError, regisErr.Error())
    }
    if err != nil {
		return redirectWithMessage(c, "Register Failed!", false, "/form-register")
	}
    return redirectWithMessage(c, "Register Sucsses!", true ,"/form-login")
}

    // e.GET("/form-login/", login)
func formlogin(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/form-login.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    sess,errSess:=session.Get("session", c)
    if errSess != nil {
        return c.JSON(http.StatusInternalServerError, errSess.Error())
    }
    mFlash := map[string]interface{}{
        "MessageFlash": sess.Values["message"], // regis berhasil
        "StatusFlash": sess.Values["status"], // regis berhasil
    }
    delete(sess.Values, "message")
    delete(sess.Values, "status")
    sess.Save(c.Request(), c.Response())
    
    return tmpl.Execute(c.Response(),mFlash)
}

    // e.POST("/login", login)
func login(c echo.Context) error {
    email := c.FormValue("input-email")
    password := c.FormValue("input-password")

    // check
    // fmt.Println(email, password)
    // return c.JSON(http.StatusInternalServerError, map[string]string{
    //     "email" : email,
    //     "password" : password,
    // })

    user := User{}
    // connection to database and check a email true or false
    err:=connection.Conn.QueryRow(context.Background(), "SELECT id, name, email, password FROM tb_user WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword)
    if err != nil {
        return redirectWithMessage(c, "Email Incorrect!", false,  "/form-login")
    }
    errPassword := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
    if errPassword !=nil {
        return redirectWithMessage(c, "Password Incorrect!", false, "/form-login")
    }
    // return c.JSON(http.StatusOK, "Login Berhasil!")
    sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM MASA BERLAKU LOGIN
	sess.Values["message"] = "Login success"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["Id"] = user.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(),  c.Response())

    fmt.Println("Author ID : ",sess.Values["Id"]) 
    
    return c.Redirect(http.StatusMovedPermanently, "/")
}

    // e.POST("/logout", logout)
func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge= -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, Path string) error {
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, Path)
}