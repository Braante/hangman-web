package main

import (
	"fmt"
	"hangmanweb"
	"net/http"
	"text/template"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	a := hangmanweb.WordChoose()
	b := hangmanweb.PlusALea(a)
	tmpl := template.Must(template.ParseFiles("html/test.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			letter := r.FormValue("letter")
			fmt.Print(letter)
		}

		data := TodoPageData{
			PageTitle: string(b),
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)

}
