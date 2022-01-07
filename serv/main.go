package main

import (
	"fmt"
	"hangmanweb"
	"net/http"
	"text/template"
)

type TodoPageData struct {
	PageTitle    string
	Attemptsleft int
	TextDeco     string
}

func main() {
	TextDeco := ""
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
			fmt.Println("repasse")
			if attempts == 11 {
				b = a
				TextDeco = "You win"
			}
			if attempts <= 0 {
				TextDeco = "You lose"
			}
			fmt.Println(attempts)
			//fmt.Print(letter)
		}
		rep := hangmanweb.PrintTable(b)
		if attempts == 11 || attempts <= 0 {
			data := TodoPageData{
				PageTitle: rep,
				TextDeco:  TextDeco,
			}
			tmpl.Execute(w, data)
		} else {
			data := TodoPageData{
				PageTitle:    rep,
				Attemptsleft: attempts,
			}
			tmpl.Execute(w, data)
		}

	})
	fs := http.FileServer(http.Dir("css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.ListenAndServe(":8080", nil)
}
