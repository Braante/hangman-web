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

type StartData struct {
	PageLevel string
}

func main() {
	tmpl1 := template.Must(template.ParseFiles("html/start.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			level := r.FormValue("gender")
			fmt.Println(level)
		}
		dota := StartData{}
		tmpl1.Execute(w, dota)
	})
	TextDeco := ""
	var a []byte
	var b []byte
	attempts := 10
	var min bool
	var maj bool
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		if TextDeco != "You win" && TextDeco != "You lose" {

			switch r.Method {
			case "GET":
				http.Redirect(w, r, "/", 301)

			case "POST":
				listwords := r.FormValue("Difficulty")
				letter := r.FormValue("letter")
				if len(listwords) != 0 {
					a = hangmanweb.WordChoose(listwords)
					/* --- cheat code --- */
					fmt.Println(string((a)))
					b = hangmanweb.PlusALea(a)
					min, maj = hangmanweb.Initialisation(a)
					//fmt.Print(min, maj)
				}

				attempts = hangmanweb.CheckAccents(min, maj, b, a, attempts, letter)
				if attempts == 11 {
					b = a
					TextDeco = "You win"
				}
				if attempts <= 0 {
					TextDeco = "You lose"
				}
				//fmt.Println(attempts)
				//fmt.Print(letter)
			}
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
