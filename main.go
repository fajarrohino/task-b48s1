package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

// routes
func main() {
    e := echo.New()
    e.Static("/aset", "aset")

    e.GET("/home", home)
    e.GET("/project", project)
    e.GET("/testimonial", testimonial)
    e.GET("/contact", contact)
    e.GET("/project-detail", projectdetail)
    e.POST("/project-add", addproject)

    e.Logger.Fatal(e.Start("localhost:5000"))
}

// handlers
func home(c echo.Context) error{
    tmpl, err := template.ParseFiles("views/index.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    return tmpl.Execute(c.Response(),nil)
}
func project(c echo.Context) error  {
    tmpl, err := template.ParseFiles("views/project.html")
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
    tmpl, err := template.ParseFiles("views/project-detail.html")
    if err != nil {
        return c.JSON(http.StatusInternalServerError,err.Error())
    }
    return tmpl.Execute(c.Response(),nil)
}
func addproject(c echo.Context) error  {
    nameProject:=c.FormValue("input-project-name")
    startDate:=c.FormValue("input-start-date")
    endDate:=c.FormValue("input-end-date")
    description:=c.FormValue("input-description")
    nodejs:=c.FormValue("input-nodejs")
    reactjs:=c.FormValue("input-reactjs")
    nextjs:=c.FormValue("input-nextjs")
    typoscript:=c.FormValue("input-typoscript")

    fmt.Println("Nama Project : ",nameProject)
    fmt.Println("Start Date : ",startDate)
    fmt.Println("End Date : ",endDate)
    fmt.Println("Description : ",description)
    fmt.Println("Reactjs : ",reactjs)
    fmt.Println("Nextjs : ",nextjs)
    fmt.Println("Typescript : ",typoscript)
    fmt.Println("Nodejs : ",nodejs)

    return c.Redirect(http.StatusMovedPermanently, "/project")
}