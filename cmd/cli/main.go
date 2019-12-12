package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"github.com/theantichris/github"
)

const issueListTemplate = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------------
Number: {{.Number}}
User:	{{.User.Login}}
Title:	{{.Title | printf "%.64s"}}
Age:	{{.CreatedAt | daysAgo}} days
{{end}}`

var issueList = template.Must(template.New("issueList").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(issueListTemplate))

func main() {
	issues, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := issueList.Execute(os.Stdout, issues); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
