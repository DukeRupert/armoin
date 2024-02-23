package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// path parameters
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// query parameters
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// html templates
func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}

func main() {
	e := echo.New()
	// Parse templates
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	// Pass templates to Renderer
	e.Renderer = t
	e.GET("/hello", Hello)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":1323"))
}
