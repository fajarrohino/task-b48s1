package main

import (
	"context"
	"day10/connection"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
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
}
// slice of struc ->mirip array of objec
var Projects = []DataProject{
    // {
    //     ProjectName     :"Project 1",
    //     StartDate        :"14-05-2023",
    //     EndDate         :"14-07-2023",
    //     Duration        :"2 Month",
    //     Description     :"Maap mata lumayan sakit",
    //     Image           :"code.jpg",
    //     NodeJs          :true,
    //     ReactJs         :true,
    //     NextJs          :true,
    //     TypoScript      :true,
    // },
    // {
    //     ProjectName     :"Project 2",
    //     StartDate        :"15-06-2023",
    //     EndDate         :"15-07-2023",
    //     Duration        :"1 Month",
    //     Description     :"Maap mata lumayan sakit",
    //     Image           :"code1.jpg",
    //     NodeJs          :true,
    //     ReactJs         :true,
    //     NextJs          :true,
    //     TypoScript      :true,
    // },
    // {
    //     ProjectName     :"Project 3",
    //     StartDate        :"15-07-2023",
    //     EndDate         :"15-08-2023",
    //     Duration        :"1 Month",
    //     Description     :"Maap mata lumayan sakit",
    //     Image           :"code.jpg",
    //     NodeJs          :true,
    //     ReactJs         :true,
    //     NextJs          :true,
    //     TypoScript      :true,
    // },
}
// routes
func main() {
    e := echo.New()

    // connection database
    connection.DatabaseConnect()

    // connecting aset
    e.Static("/aset", "aset")

    // Routes Get
    e.GET("/", home)
    e.GET("/form-project", formproject)
    e.GET("/testimonial", testimonial)
    e.GET("/contact", contact)
    e.GET("/project-detail/:id", projectdetail)
    e.GET("/project-edit/:id", projectedit)

    // auth
    e.GET("/form-register", formregister)
    e.GET("/form-login", formlogin)

    // Routes Post
    e.POST("/", addproject)
    e.POST("/project-delete/:id", deleteProject)
    e.POST("/project-edit/:Id", submitEditProject)

    e.Logger.Fatal(e.Start("localhost:5500"))
}

// handlers

    // GET("/", home)
func home(c echo.Context) error{
    tmpl, err := template.ParseFiles("views/index.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    
    // utk mengambil data di database
    Querys,errQuery:= connection.Conn.Query(context.Background(), "SELECT * FROM tb_project")

    if errQuery != nil {
        return c.JSON(http.StatusInternalServerError, errQuery.Error())
    }
    var inputProject []DataProject
    for Querys.Next(){
        var each = DataProject{}
        err:= Querys.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Image, &each.NodeJs, &each.ReactJs, &each.NextJs, &each.TypoScript, &each.Duration)

        if err != nil {
            return c.JSON(http.StatusInternalServerError,err.Error())
        }
        inputProject = append(inputProject, each)
    }
    
    // utk pemanggilan di html index
    myProject:=map[string]interface{}{
        "myproject" : inputProject,
    }
    // fmt.Println("ini data index", myProject)
    return tmpl.Execute(c.Response(),myProject)
}
    // GET("/form-project", formproject)
func formproject(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/add-project.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    return tmpl.Execute(c.Response(),nil)
}
    // GET("/testimonial", testimonial)
func testimonial(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/testimonial.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    return tmpl.Execute(c.Response(),nil)
}
    // GET("/contact", contact)
func contact(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/contact.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    return tmpl.Execute(c.Response(),nil)
}
    // GET("/project-detail/:id", projectdetail)
func projectdetail(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/project-detail.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    id,_:= strconv.Atoi(c.Param("id"))

    projectdetail := DataProject{}
    
    errQuery :=connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE id=$1",id).Scan(&projectdetail.Id, &projectdetail.ProjectName, &projectdetail.StartDate, &projectdetail.EndDate, &projectdetail.Description, &projectdetail.Image, &projectdetail.NodeJs, &projectdetail.ReactJs, &projectdetail.NextJs, &projectdetail.TypoScript, &projectdetail.Duration)

    if errQuery != nil {
        return c.JSON(http.StatusInternalServerError, errQuery.Error())
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
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    id, _ := strconv.Atoi(c.Param("id"))

    editDp := DataProject{}

    errEDP :=connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE Id=$1",id).Scan(&editDp.Id, &editDp.ProjectName, &editDp.StartDate, &editDp.EndDate, &editDp.Description, &editDp.Image, &editDp.NodeJs, &editDp.ReactJs, &editDp.NextJs, &editDp.TypoScript, &editDp.Duration)
    if errEDP != nil {
        return c.JSON(http.StatusInternalServerError, errEDP.Error())
    }
    
    edit :=map[string]interface{}{
        "Id"  : id,
        "data":   editDp,
        // dp ->DataProject
    }
    return tmpl.Execute(c.Response(),edit)
}
    // GET("/form-register/", register)
func formregister(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/form-register.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    return tmpl.Execute(c.Response(),nil)
}
    // e.GET("/form-login/", login)
func formlogin(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/form-login.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    return tmpl.Execute(c.Response(),nil)
}

    // POST("/project-delete/:id", deleteProject)
func deleteProject(c echo.Context) error  {
    id, _:= strconv.Atoi(c.Param("id"))

    _, errDelete:=connection.Conn.Exec(context.Background(),"DELETE FROM tb_project WHERE Id=$1", id)
    if errDelete != nil {
        return c.JSON(http.StatusInternalServerError, errDelete.Error())
    }

    // Projects = append(Projects[:id], Projects[id+1:]...)
	return c.Redirect(http.StatusMovedPermanently, "/")
}
    // POST("/project-edit/:id", submitEditProject)
func submitEditProject(c echo.Context) error  {
    // mengambil id dr parms
    // id,_:=strconv.Atoi(c.Param("id"))
    id:=c.FormValue("id")
    nameProject:=c.FormValue("input-project-name")
    startDate:=c.FormValue("input-start-date")
    endDate:=c.FormValue("input-end-date")
    description:=c.FormValue("input-description")
    image:=c.FormValue("input-img")
    nodejs:=c.FormValue("input-nodejs")
    reactjs:=c.FormValue("input-reactjs")
    nextjs:=c.FormValue("input-nextjs")
    typoscript:=c.FormValue("input-typoscript")

    // check input data
    // fmt.Println(id, "ini iddddd")
    // fmt.Println(nameProject)
    // fmt.Println(startDate)
    // fmt.Println(endDate)
    // fmt.Println(description)
    // fmt.Println(image)
    // fmt.Println(nodejs)
    // fmt.Println(reactjs)
    // fmt.Println(nextjs)
    // fmt.Println(typoscript)

    Duration := coutDuration(startDate, endDate)

    // connection to database
    update, errupdate :=connection.Conn.Exec(context.Background(), "UPDATE tb_project SET project_name=$1, start_date=$2, end_date=$3, description=$4, image=$5, nodejs=$6, reactjs=$7, nextjs=$8, typoscript=$9, duration=$10 WHERE id=$11", nameProject, startDate, endDate, description, image, nodejs != "", reactjs != "", nextjs != "", typoscript != "",Duration, id)
    if errupdate != nil {
        return c.JSON(http.StatusInternalServerError, errupdate.Error())
    }
    fmt.Println("data masuk", update.RowsAffected())
    return c.Redirect(http.StatusMovedPermanently, "/")
}
// POST("/", addproject)
func addproject(c echo.Context) error  {
    // input data
    nameProject:=c.FormValue("input-project-name")
    startDate:=c.FormValue("input-start-date")
    endDate:=c.FormValue("input-end-date")
    description:=c.FormValue("input-description")
    image:=c.FormValue("input-img")
    nodejs:=c.FormValue("input-nodejs")
    reactjs:=c.FormValue("input-reactjs")
    nextjs:=c.FormValue("input-nextjs")
    typoscript:=c.FormValue("input-typoscript")

    // variabel duration
    Duration := coutDuration(startDate,endDate)

    // connection database
    _,err:=connection.Conn.Exec(context.Background(),
    "INSERT INTO tb_project (project_name, start_date, end_date, description, image, nodejs, reactjs, nextjs, typoscript,  duration) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", nameProject, startDate, endDate, description, image,(nodejs=="nodeJs"), (reactjs=="reactJs"), (nextjs=="nextJs"), (typoscript=="typoscript"), Duration)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.Redirect(http.StatusMovedPermanently, "/")
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