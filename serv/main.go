package main

import (
	"fmt"
	"hangmanweb"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

type WinData struct {
	WordFind    string
	Scoreboards string
}

type LoseData struct {
	WordFind    string
	Scoreboards string
}

type SaveStruct struct {
	NameSave string
	Score    int
}

type TodoPageData struct {
	PageTitle    string
	Attemptsleft int
	LWU          string
	LetterUsed   string
	UserName     string
	Restart      string
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
	chr := ""
	lose := false
	usersaved := false
	var rep string
	win := false
	var a []byte
	var b []byte
	attempts := 10
	min := false
	maj := false
	var tableauX []byte
	name := ""
	var listwords string
	score := 0
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			http.Redirect(w, r, "/", 301)
		case "POST":
			fmt.Println("namia:", name)
			if name == "" {
				name = r.FormValue("Username")
			}
			fmt.Println("hfeu:", win)
			if !win && listwords == "" {
				fmt.Println("aieaieaie")
				listword := r.FormValue("Difficulty")
				listwords = listword
			}
			letter := r.FormValue("letter")
			fmt.Println("listorsd:", listwords)
			if (!min && !maj) || win || lose {
				fmt.Println("peteuurtur")
				a = hangmanweb.WordChoose(listwords)
				/* --- cheat code --- */
				fmt.Println(string((a)))
				b = hangmanweb.PlusALea(a)
				min, maj = hangmanweb.Initialisation(a)
				win = false
				lose = false
			}

			attempts, tableauX = hangmanweb.CheckAccents(min, maj, b, a, attempts, letter, tableauX)
			if attempts == 11 {
				b = a
				attempts = 10
				rep = ""
				win = true
				usersaved = false
				http.Redirect(w, r, "/win", 301)
			}
			if attempts <= 0 {
				attempts = 10
				rep = ""
				lose = true
				score = 0
				http.Redirect(w, r, "/lose", 301)
			}
		}
		rep = hangmanweb.PrintTable(b)
		if attempts == 11 || attempts <= 0 {
			data := TodoPageData{
				PageTitle:    rep,
				LWU:          listwords,
				Attemptsleft: attempts,
				UserName:     name,
			}
			tmpl.Execute(w, data)
		} else {
			letterused := hangmanweb.PrintTableEspace(tableauX)
			fmt.Println("passagelaici")
			data := TodoPageData{
				PageTitle:    rep,
				Attemptsleft: attempts,
				LetterUsed:   letterused,
				UserName:     name,
				LWU:          listwords,
			}
			tmpl.Execute(w, data)
		}
	})

	tmplWin := template.Must(template.ParseFiles("html/win.html"))
	http.HandleFunc("/win", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if !win {
				http.Redirect(w, r, "/", 301)
			}
		case "POST":
			if !usersaved && win {
				oldchr, _ := ioutil.ReadFile("save.txt")
				oldchrp := hangmanweb.PrintTable(oldchr)
				score++
				file, _ := os.Create("save.txt")
				scorestr := strconv.FormatInt(int64(score), 10)
				chr = name + " " + scorestr + "<br>" + oldchrp
				file.WriteString(chr)
				usersaved = true
				file.Close()
			}
		}
		fmt.Println("scoreboardfjier:0", chr)
		doti := WinData{
			WordFind:    rep,
			Scoreboards: chr,
		}
		tmplWin.Execute(w, doti)
	})

	tmplLose := template.Must(template.ParseFiles("html/lose.html"))
	http.HandleFunc("/lose", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if !lose {
				http.Redirect(w, r, "/", 301)
			}
		}
		resp := hangmanweb.PrintTable(a)
		dyti := LoseData{
			WordFind: resp,
		}
		tmplLose.Execute(w, dyti)
	})

	fs := http.FileServer(http.Dir("css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	fs2 := http.FileServer(http.Dir("images/"))
	http.Handle("/images/", http.StripPrefix("/images/", fs2))
	http.ListenAndServe(":8080", nil)
}
