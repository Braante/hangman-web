package main

import (
	"encoding/json"
	"fmt"
	"hangmanweb"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type WinData struct {
	PageWin string
}

type SaveStruct struct {
	Score    int
	NameSave string
}

type TodoPageData struct {
	PageTitle    string
	Attemptsleft int
	TextDeco     string
	LetterUsed   string
	UserName     string
}

type StartData struct {
	PageLevel string
}

func main() {
	tmpl1 := template.Must(template.ParseFiles("html/start.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dota := StartData{}
		tmpl1.Execute(w, dota)
	})
	tmplWin := template.Must(template.ParseFiles("html/win.html"))
	http.HandleFunc("/win", func(w http.ResponseWriter, r *http.Request) {
		doti := WinData{}
		tmplWin.Execute(w, doti)
	})
	TextDeco := ""
	var a []byte
	var b []byte
	attempts := 10
	var min bool
	var maj bool
	var tableauX []byte
	name := ""
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		if TextDeco != "You win" && TextDeco != "You lose" {

			switch r.Method {
			case "GET":
				http.Redirect(w, r, "/", 301)

			case "POST":
				if name == "" {
					name = r.FormValue("Username")
				}
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

				attempts, tableauX = hangmanweb.CheckAccents(min, maj, b, a, attempts, letter, tableauX)
				if attempts == 11 {
					saveos, _ := os.Open("save.txt")
					info, _ := os.Stat("save.txt")
					size := info.Size()
					old := make([]byte, size)
					saveos.Read(old)
					saveos.Close()
					m := SaveStruct{attempts, name}
					saved, _ := json.Marshal(m)
					for i := 0; i < len(old); i++ {
						saved = append(saved, old[i])
					}
					ioutil.WriteFile("save.txt", saved, 0777)
					b = a
					http.Redirect(w, r, "/win", 301)
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
				PageTitle:    rep,
				TextDeco:     TextDeco,
				Attemptsleft: attempts,
				UserName:     name,
			}
			tmpl.Execute(w, data)
		} else {
			letterused := hangmanweb.PrintTableEspace(tableauX)
			data := TodoPageData{
				PageTitle:    rep,
				Attemptsleft: attempts,
				LetterUsed:   letterused,
				UserName:     name,
			}
			tmpl.Execute(w, data)
		}

	})
	fs := http.FileServer(http.Dir("css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	fs2 := http.FileServer(http.Dir("images/"))
	http.Handle("/images/", http.StripPrefix("/images/", fs2))
	http.ListenAndServe(":8080", nil)
}
