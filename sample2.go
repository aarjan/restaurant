package main

import (
	"html/template"
	"os"
)

type s struct {
	A
	B int
}

type A struct {
	I int
}

func main() {
	s1 := s{A{2}, 3}
	// slic := []int{1, 3, 34}
	templ, _ := template.New("test").Parse(`{{if isset .B 3}}"hello"{{end}}`)
	templ.Execute(os.Stdout, s1)
}
