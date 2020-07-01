package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
)

// Template for html
type Template struct {
	templates *template.Template
}

// Render func for html template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World@@")
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("./public/views/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":80"))
}
