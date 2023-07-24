package main

import (
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)
type DataProject struct{
	    Id              int
        ProjectName     string
        StartDate        string
        EndDate         string
        Duration        string
        Description     string
        Image           string
        NodeJs          bool
        ReactJs         bool
        NextJs          bool
        TypoScript      bool
}

var Projects = []DataProject{
    {
        ProjectName     :"Project 1",
        StartDate        :"14-05-2023",
        EndDate         :"14-07-2023",
        Duration        :"2 Month",
        Description     :"Maap mata lumayan sakit",
        Image           :"code.jpg",
        NodeJs          :true,
        ReactJs         :true,
        NextJs          :true,
        TypoScript      :true,
    },
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
    e.Static("/aset", "aset")

    // Routes Get
    e.GET("/", home)
    e.GET("/form-project", formproject)
    e.GET("/testimonial", testimonial)
    e.GET("/contact", contact)
    e.GET("/project-detail/:id", projectdetail)
    e.GET("/project-edit/:id", projectedit)

    // Routes Post
    e.POST("/", addproject)
    e.POST("/project-delete/:id", deleteProject)
    e.POST("/project-edit/:id", submitEditProject)

    e.Logger.Fatal(e.Start("localhost:5000"))
}

// handlers
func home(c echo.Context) error{
    tmpl, err := template.ParseFiles("views/index.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    // utk pemanggilan di html index
    Projects:=map[string]interface{}{
        "Projects" : Projects,
    }
    return tmpl.Execute(c.Response(),Projects)
}
func formproject(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/add-project.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    return tmpl.Execute(c.Response(),nil)
}
func testimonial(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/testimonial.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    return tmpl.Execute(c.Response(),nil)
}
func contact(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/contact.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    return tmpl.Execute(c.Response(),nil)
}
func projectdetail(c echo.Context) error  {
    // id := c.Param("id")
    tmpl, err := template.ParseFiles("views/project-detail.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }

    id, _ := strconv.Atoi(c.Param("id"))

    projectdetail := DataProject{}
    for i, data := range Projects{
        if id == i {
            projectdetail = DataProject{
                ProjectName:    data.ProjectName,
                StartDate:       data.StartDate,
                EndDate:        data.EndDate,
                Duration:       data.Duration,
                Description:    data.Description,
                Image:           data.Image,
                NodeJs:         data.NodeJs,
                ReactJs:        data.ReactJs,
                NextJs:         data.NextJs,
                TypoScript:     data.TypoScript,
            }
        }
    }
    data :=map[string]interface{}{
        "dp":   projectdetail,
        // dp ->DataProject
    }
    return tmpl.Execute(c.Response(),data)
}
func projectedit(c echo.Context) error  {
    // id := c.Param("id")
    tmpl, err := template.ParseFiles("views/edit-project.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,map[string]string{"pesan": err.Error()})
    }
    id, _ := strconv.Atoi(c.Param("id"))

    projectdetail := DataProject{}
    for i, data := range Projects{
        if id == i {
            projectdetail = DataProject{
                ProjectName:    data.ProjectName,
                StartDate:       data.StartDate,
                EndDate:        data.EndDate,
                Duration:       data.Duration,
                Description:    data.Description,
                Image:           data.Image,
                NodeJs:         data.NodeJs,
                ReactJs:        data.ReactJs,
                NextJs:         data.NextJs,
                TypoScript:     data.TypoScript,
            }
        }
    }
    data :=map[string]interface{}{
        "dp":   projectdetail,
        "Id": id,
        // dp ->DataProject
    }
    return tmpl.Execute(c.Response(),data)
}
func deleteProject(c echo.Context) error  {
    id, _:= strconv.Atoi(c.Param("id"))

    Projects = append(Projects[:id], Projects[id+1:]...)
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func submitEditProject(c echo.Context) error  {
    // mengambil id dr parms
    id,_:=strconv.Atoi(c.Param("id"))

    nameProject:=c.FormValue("input-project-name")
    startDate:=c.FormValue("input-start-date")
    endDate:=c.FormValue("input-end-date")
    description:=c.FormValue("input-description")
    image:=c.FormValue("input-img")
    nodejs:=c.FormValue("input-nodejs")
    reactjs:=c.FormValue("input-reactjs")
    nextjs:=c.FormValue("input-nextjs")
    typoscript:=c.FormValue("input-typoscript")

    var projectedit = DataProject{
                ProjectName:    nameProject,
                StartDate:      startDate,
                EndDate:        endDate,
                Duration:       coutDuration(startDate, endDate),
                Description:    description,
                Image:          image,
                NodeJs:         (nodejs=="nodeJs"),
                ReactJs:        (reactjs=="reactJs"),
                NextJs:         (nextjs=="nextJs"),
                TypoScript:     (typoscript=="typoscript"),
            }
    Projects[id] = projectedit
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
func addproject(c echo.Context) error  {
    // input data
    nameProject:=c.FormValue("input-project-name")
    sDate:=c.FormValue("input-start-date")
    eDate:=c.FormValue("input-end-date")
    description:=c.FormValue("input-description")
    image:=c.FormValue("input-img")
    nodejs:=c.FormValue("input-nodejs")
    reactjs:=c.FormValue("input-reactjs")
    nextjs:=c.FormValue("input-nextjs")
    typoscript:=c.FormValue("input-typoscript")

    var newProject = DataProject{
                ProjectName:    nameProject,
                StartDate:      sDate,
                EndDate:        eDate,
                Duration:       coutDuration(sDate, eDate),
                Description:    description,
                Image:          image,
                NodeJs:         (nodejs=="nodeJs"),
                ReactJs:        (reactjs=="reactJs"),
                NextJs:         (nextjs=="nextJs"),
                TypoScript:     (typoscript=="typoscript"),
    }
    Projects = append(Projects, newProject)
    return c.Redirect(http.StatusMovedPermanently, "/")
}