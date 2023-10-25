package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"lessonProgram/db"
	"lessonProgram/models"
	"lessonProgram/routes"
)

func main() {
	db.InitDB()

	err := db.Db.AutoMigrate(&models.Student{}, &models.Todo{})

	if err != nil {
		panic("Tablo oluşturma hatası: " + err.Error())
	}

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	routes.InitRoutes(e)

	e.Start(":8080")
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
