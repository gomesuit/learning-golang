package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	fmt.Println("vim-go")
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
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
