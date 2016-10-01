package main

import (
	"os"
	"text/template"
)

func main() {
	sweaters := Inventory{"wool", 17}
	//tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	tmpl, err := template.ParseFiles("sample.txt")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

type Inventory struct {
	Material string
	Count    uint
}
