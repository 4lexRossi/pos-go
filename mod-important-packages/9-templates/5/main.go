package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Courses{
		{"GO", 40},
		{"Java", 50},
		{"Python", 60},
		{"NodeJS", 70},
	})
	if err != nil {
		panic(err)
	}
}
