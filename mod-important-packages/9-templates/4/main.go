package main

import (
	"html/template"
	"net/http"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {

	http.HandleFunc("/", showCourses)
	http.ListenAndServe(":8080", nil)
}

func showCourses(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(w, Courses{
		{"GO", 40},
		{"Java", 50},
		{"Python", 60},
		{"NodeJS", 70},
	})
	if err != nil {
		panic(err)
	}
}
