package main

import (
	"log"
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	err := temp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = temp.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = temp.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = temp.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
