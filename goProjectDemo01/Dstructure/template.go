package main

import (
	"os"
	"text/template"
)

const myTemplate = `{{$variable := .}}{{$variable | printf "%d"}}`

/* func add(a int, b int) int {
	return a + b
} */

func add(a int) int {
	return a + 1
}

func main() {
	/* 	person := Person{
		Name: "kyrie",
		Age:  27,
	} */

	/* 	mymap := map[string]bool{
		"kyrie": true,
		"james": true,
		"kobe":  false,
	} */
	/* 	array := [9]int{7, 8, 7, 1, 4, 8, 9, 2, 8} */
	var num int = -10
	template, _ := template.New("myTemplate").Funcs(template.FuncMap{"add": add}).Parse(myTemplate)
	template.Execute(os.Stdout, num)
}
