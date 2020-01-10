package main

import (
	"log"
	"os"
	"text/template"

)

var templates *template.Template

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml") // does nothing but gets reads vales of tpl.gohtml

	if err != nil{
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html") // Creates an html file
	if err != nil{
		log.Println("Error containing file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil) // puts value og gpl.gohtml into index.html
	if err != nil{
		log.Fatalln(err)
	}

// adding some more files in tpl to parse
	tpl, err = template.ParseFiles("one.gmao")
	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err!= nil{
		log.Fatalln(err)
	}

	tpl,err = tpl.ParseFiles("two.gmao", "vespa.gmao")
	if err!= nil{
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err!= nil{
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err!= nil{
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err!= nil{
		log.Fatalln(err)
	}
}
