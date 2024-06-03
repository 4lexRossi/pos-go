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
	t := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} - Workload: {{.Workload}}"))
	// tmp := template.New("CourseTemplate")
	// tmp, err := tmp.Parse("Course: {{.Name}} - Workload: {{.Workload}}")
	// if err != nil {
	// 	panic(err)
	// }
	err := t.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
