package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	course := Course{"GO", 40}
	tmp := template.New("CourseTemplate")
	tmp, err := tmp.Parse("Course: {{.Name}} - Workload: {{.Workload}}")
	if err != nil {
		panic(err)
	}
	err = tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
