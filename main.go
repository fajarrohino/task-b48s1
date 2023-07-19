package main

import (
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

// routes
func main() {
    e := echo.New()
    e.Static("/aset", "aset")

    e.GET("/index", home)
    e.GET("/project", project)
    e.GET("/testimonial", testimonial)
    e.GET("/contact", contact)
    e.GET("/project-detail", projectdetail)
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