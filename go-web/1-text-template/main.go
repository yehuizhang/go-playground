package main

import (
	"log"
	"os"
	"text/template"
)

var temp *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	// temp = template.Must(template.ParseGlob("templates/*.gohtml"))
	temp = template.Must(template.ParseGlob("templates/*.*html"))
}

func main() {

	// err := temp.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = temp.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = temp.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = temp.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// names := []string{"Adam", "Becky", "Chris", "David"}
	// err = temp.ExecuteTemplate(os.Stdout, "array.html", names)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cities := map[string]string{
	// 	"China":  "Beijing",
	// 	"France": "Paris",
	// 	"Japan":  "Tokyo",
	// }
	// err = temp.ExecuteTemplate(os.Stdout, "map.html", cities)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// goblin := sage{"Goblin", "Time is money, frind!"}
	// err := temp.ExecuteTemplate(os.Stdout, "struct.html", goblin)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	sages := []sage{
		sage{"Goblin", "Time is money, frind!"},
		sage{"Human", "My strength is weak"},
	}
	err := temp.ExecuteTemplate(os.Stdout, "structs.html", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
