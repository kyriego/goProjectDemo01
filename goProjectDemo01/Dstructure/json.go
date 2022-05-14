package main

import (
	"flag"
	"log"
	"os"
	"text/template"
	"time"

	"gopl.io/ch4/github"
)

const (
	MonthHours = 30 * 24
	YearHours  = 365 * 24
)

//	item.Number, item.User.Login, item.Title)
/* const myTemplate = `{{$Count := 0}}{{$MonthHours := .MonthHours}}{{$YearHours := .YearHours}}Issues Count:{{len .Items}}
{{range .Items}}{{$CurrHours := hoursLater .CreatedAt}}{{if le $CurrHours $YearHours}}Number:{{.Number}}
{{if .Title}}Title:{{.Title}}
{{else}}{{end}}{{if .User.Login}}Login:{{.User.Login}}
{{else}}{{end}}
{{else}}{{end}}
{{end}}
Total={{$Count}}` */

const myTemplate = `{{$Count := 0}}{{$MonthHours := .MonthHours}}{{$YearHours := .YearHours}}Issues Count:{{len .Items}}

{{range .Items}}{{$CurrHours := hoursLater .CreatedAt}}{{if ge $CurrHours $YearHours}}{{$Count = add $Count}}{{if .Number}}Number:{{.Number}}
{{else}}{{end}}{{if .Title}}Title:{{.Title}}
{{else}}{{end}}{{if .User.Login}}UserLogin:{{.User.Login}}
{{else}}{{end}}
{{else}}{{end}}{{end}}
TotalCount: {{$Count}}`

var t = flag.Int("t", 0, "choose the time Limit in \"in a month\"(0), \"after a month and in a year\"(1), \"after a year\"(2)")

type templateData struct {
	Items      []*github.Issue
	MonthHours int
	YearHours  int
}

func hoursLater(t time.Time) int {
	d := time.Since(t)
	f := d.Hours()
	return int(f)
}

func add(a int) int {
	return a + 1
}

func main() {
	flag.Parse()
	result, err := github.SearchIssues(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
	/* 	fmt.Printf("%d result.Items:\n", len(result.Items))
	   	count := 0
	   	for _, item := range result.Items {
	   		d := time.Since(item.CreatedAt)
	   		hours := d.Hours()
	   		switch *t {
	   		case 0:
	   			if hours <= MonthHours {
	   				fmt.Printf("#%-5d %9.9s %.55s\n",
	   					item.Number, item.User.Login, item.Title)
	   			}
	   		case 1:
	   			if hours > MonthHours && hours <= YearHours {
	   				fmt.Printf("#%-5d %9.9s %.55s\n",
	   					item.Number, item.User.Login, item.Title)
	   			}
	   		case 2:
	   			if hours > YearHours {
	   				count++
	   				fmt.Printf("#%-5d %9.9s %.55s\n",
	   					item.Number, item.User.Login, item.Title)
	   			}
	   		}
	   	}
	   	fmt.Printf("TotalCount = %d\n", count) */
	data := templateData{result.Items, MonthHours, YearHours}
	template := template.Must(template.New("myTemplate").Funcs(template.FuncMap{"hoursLater": hoursLater, "add": add}).Parse(myTemplate))
	template.Execute(os.Stdout, data)
}
