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
	/* --- cheat code --- */
	fmt.Println(string((a)))
	b := hangmanweb.PlusALea(a)
	min, maj, attempts := hangmanweb.Initialisation(a)
	//fmt.Print(min, maj)
	tmpl := template.Must(template.ParseFiles("html/test.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			letter := r.FormValue("letter")
			attempts = hangmanweb.CheckAccents(min, maj, b, a, attempts, letter)
			if attempts == 11 {
				b = a
			}
			fmt.Println(attempts)
			//fmt.Print(letter)
		}
		rep := hangmanweb.PrintTable(b)
		data := TodoPageData{
			PageTitle: rep,
		}
		tmpl.Execute(w, data)
	})
	fs := http.FileServer(http.Dir("css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.ListenAndServe(":8080", nil)
}
