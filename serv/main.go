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
	ScorePerso   int
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
	sameuser := false
	long := 0
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
			if name == "" {
				name = r.FormValue("Username")
			}
			if !win && listwords == "" {
				listword := r.FormValue("Difficulty")
				listwords = listword
			}
			letter := r.FormValue("letter")
			if (!min && !maj) || win || lose {
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
				tableauX = nil
				http.Redirect(w, r, "/win", 301)
			}
			if attempts <= 0 {
				attempts = 10
				rep = ""
				lose = true
				score = 0
				tableauX = nil
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
				ScorePerso:   score,
			}
			tmpl.Execute(w, data)
		} else {
			letterused := hangmanweb.PrintTableEspace(tableauX)
			data := TodoPageData{
				PageTitle:    rep,
				Attemptsleft: attempts,
				LetterUsed:   letterused,
				UserName:     name,
				LWU:          listwords,
				ScorePerso:   score,
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
				if score < 10 {
					long = 6
				} else {
					long = 7
				}
				cpt := 0
				if len(oldchrp) >= len(name)+long {
					for k := 0; k < len(name); k++ {
						if oldchrp[k] == name[k] {
							cpt++
						}
					}
					if cpt == len(name) {
						oldchrp = oldchrp[len(name)+long:]
						sameuser = true
					}

				}
				file, _ := os.Create("save.txt")
				scorestr := strconv.FormatInt(int64(score), 10)
				chr = name + " " + scorestr + "<br>" + oldchrp
				cptfin := 0
				for k := 0; k < len(chr); k++ {
					if chr[k] == '<' {
						cptfin++
					}
				}
				var invchr string
				var invchr2 string
				if cptfin == 11 && !sameuser {
					for _, v := range chr {
						invchr = string(v) + invchr
					}
					invchr = invchr[1:]
					cptlim := 0
					for k := 0; k < len(invchr); k++ {
						if invchr[k] != '>' {
							cptlim++
						} else {
							break
						}
					}
					invchr = invchr[cptlim:]
					for _, v := range invchr {
						invchr2 = string(v) + invchr2
					}
					chr = invchr2
				}
				file.WriteString(chr)
				usersaved = true
				file.Close()
			}
		}
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
		oldchr, _ := ioutil.ReadFile("save.txt")
		oldchrp := hangmanweb.PrintTable(oldchr)
		resp := hangmanweb.PrintTable(a)
		dyti := LoseData{
			WordFind:    resp,
			Scoreboards: oldchrp,
		}
		tmplLose.Execute(w, dyti)
	})

	fs := http.FileServer(http.Dir("css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	fs2 := http.FileServer(http.Dir("images/"))
	http.Handle("/images/", http.StripPrefix("/images/", fs2))
	http.ListenAndServe("localhost:8080", nil)
}
