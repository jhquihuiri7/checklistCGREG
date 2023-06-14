package functions

import "html/template"

func ParseTemplate() *template.Template {
	var tmp *template.Template
	tmp = template.Must(template.ParseGlob("./templates/*.html"))
	return tmp
}

