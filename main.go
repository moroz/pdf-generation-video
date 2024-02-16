package main

import (
	"os"
	"text/template"
)

var PDFTemplate = template.Must(template.ParseFiles("main.tmpl.tex"))

func main() {
	PDFTemplate.Execute(os.Stdout, nil)
}
