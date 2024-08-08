package configs

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
}
