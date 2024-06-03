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
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
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
