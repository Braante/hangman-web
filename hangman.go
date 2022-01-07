package hangmanweb

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Message struct {
	Attempts int
	Word     []byte
	Tableau  []byte
	Stock    []byte
}

func Hangman() {
	var word []byte
	maj := false
	min := false
	var attempts int
	var tableau []byte
	var stockLetter []byte
	if CheckSave() {
		save, err := os.Open(os.Args[2])
		if err != nil {
			fmt.Println("This save doesn't exist !")
			os.Exit(0)
		}
		info, _ := os.Stat(os.Args[2])
		size := info.Size()
		arr := make([]byte, size)
		save.Read(arr)
		save.Close()
		var store Message
		json.Unmarshal(arr, &store)
		attempts = store.Attempts
		word = store.Word
		tableau = store.Tableau
	} else {
		attempts = 10
		word = (WordChoose())
		tableau = []byte{}
	}
	fmt.Println("Good Luck, you have", attempts, "attempts.")
	lword := len(word)
	if !CheckSave() {
		for i := 0; i < lword; i++ {
			tableau = append(tableau, '_')
		}
	}
	var letter byte
	var compteur1 int
	if (word[0] < 91 && word[0] > 64) || (word[0] < 215 && word[0] > 191) || (word[0] < 221 && word[0] > 216) {
		maj = true
	} else if (word[0] < 123 && word[0] > 96) || (word[0] > 223 && word[0] < 247) || (word[0] > 248 && word[0] <= 255) {
		min = true
	}
	if !CheckSave() {
		for i := 0; i < (len(word)/2)-1; i++ {
			compteur1 = 0
			letter = LetterAlea(word)
			for j := 0; j < len(stockLetter); j++ {
				if letter == stockLetter[j] {
					i--
					compteur1++
					break
				} else {
					continue
				}
			}
			if compteur1 > 0 {
				continue
			} else {
				for j := 0; j < lword; j++ {
					if word[j] == letter {
						tableau[j] = letter
						stockLetter = append(stockLetter, letter)
						break
					}
				}
			}
		}
	}
	if ChooseFile() {
		PrintArtTable(tableau, min)
	} else {
		PrintTable(tableau)
		fmt.Println()
	}
	CheckAccents(min, maj, tableau, word, attempts, string(letter))
}

func WordChoose() []byte {
	var n int
	cpt := 0
	cptmot := 0
	index := 0
	args := os.Args[1]
	var word []byte
	listword, err := ioutil.ReadFile(args)
	if err != nil {
		fmt.Println("This ListWord doesn't exist !")
		os.Exit(0)
	}
	info, _ := os.Stat(args)
	size := info.Size()
	arr := make([]byte, size)
	arr = listword
	var res []byte
	for i := 0; i < len(arr); i++ {
		if arr[i] == 13 {
			cptmot++
		}
	}
	for i := 0; i < len(arr); i++ {
		res = append(res, byte(arr[i]))
	}
	if cptmot == 0 {
		return res
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n = r1.Intn(cptmot + 1)
	if n == 0 {
		n++
	}
	for i := 0; i < len(res); i++ {
		if arr[i] == 13 {
			cpt++
			if n == 1 {
				word = res[0:i]
				tempstr := string(word)
				var word []byte
				for _, wordabc := range tempstr {
					word = append(word, byte(wordabc))
				}
				return word
			}
			if cpt == n {
				word = res[index+2 : i]
				tempstr := string(word)
				var word []byte
				for _, wordabc := range tempstr {
					word = append(word, byte(wordabc))
				}
				return word
			}
			if cpt == n-1 && cptmot+1 == n {
				var temp []byte
				word = res[index+2:]
				for k := 0; k < len(word); k++ {
					if word[k] == 10 {
						temp = word[k+1:]
					}
				}
				return temp
			}
			index = i
		}
	}
	return word
}

func LetterAlea(word []byte) byte {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := r1.Intn(len(word))
	letter := word[n]
	return letter
}

func PlusALea(word []byte) []byte {
	compteur1 := 0
	var letter byte
	var stockLetter []byte
	lword := len(word)
	var tableau []byte

	for i := 0; i < lword; i++ {
		tableau = append(tableau, '_')
	}

	for i := 0; i < (len(word)/2)-1; i++ {
		compteur1 = 0
		letter = LetterAlea(word)
		for j := 0; j < len(stockLetter); j++ {
			if letter == stockLetter[j] {
				i--
				compteur1++
				break
			} else {
				continue
			}
		}
		if compteur1 > 0 {
			continue
		} else {
			for j := 0; j < lword; j++ {
				if word[j] == letter {
					tableau[j] = letter
					stockLetter = append(stockLetter, letter)
					break
				}
			}
		}
	}
	return tableau
}

func EnterLetter(letter string, tableauX []byte, lucky int) (byte, []byte, int, bool, []byte, bool) {
	isALetter := false
	isInvalid := false
	var sentence []byte
	var repsentence []byte
	var rep []byte
	for _, wordabc := range letter {
		rep = append(rep, byte(wordabc))
	}

	if len(rep) > 3 {
		isALetter = false
		for i := 0; i < len(letter)-2; i++ {
			sentence = append(sentence, letter[i])
		}
		tempsentence := string(sentence)
		for _, wordabc := range tempsentence {
			repsentence = append(repsentence, byte(wordabc))
		}
	} else {
		if rep[0] >= 32 && rep[0] <= 64 {
			isInvalid = true
		}
		isALetter = true
	}

	if isALetter && !isInvalid {
		for i := 0; i < len(tableauX); i++ {
			if rep[0] == tableauX[i] {
				//fmt.Println("You have already tried this letter.")
				lucky = 1
				//fmt.Println(tableauX)
				return rep[0], tableauX, lucky, isALetter, repsentence, isInvalid
			}
		}
		for i := 0; i <= len(tableauX); i++ {
			if i == len(tableauX) {
				//e
				if (rep[0] > 231 && rep[0] < 236) || (rep[0] > 199 && rep[0] < 204) || rep[0] == 101 || rep[0] == 69 {

					tableauX = append(tableauX, 232)
					tableauX = append(tableauX, 233)
					tableauX = append(tableauX, 234)
					tableauX = append(tableauX, 235)
					tableauX = append(tableauX, 200)
					tableauX = append(tableauX, 201)
					tableauX = append(tableauX, 202)
					tableauX = append(tableauX, 203)
					tableauX = append(tableauX, 101)
					tableauX = append(tableauX, 69)
					break
				} else if (rep[0] > 223 && rep[0] < 230) || (rep[0] > 191 && rep[0] < 198) || rep[0] == 97 || rep[0] == 65 {
					tableauX = append(tableauX, 224)
					tableauX = append(tableauX, 225)
					tableauX = append(tableauX, 226)
					tableauX = append(tableauX, 227)
					tableauX = append(tableauX, 228)
					tableauX = append(tableauX, 229)
					tableauX = append(tableauX, 192)
					tableauX = append(tableauX, 193)
					tableauX = append(tableauX, 194)
					tableauX = append(tableauX, 195)
					tableauX = append(tableauX, 196)
					tableauX = append(tableauX, 197)
					tableauX = append(tableauX, 97)
					tableauX = append(tableauX, 65)
					break
				} else if (rep[0] > 235 && rep[0] < 240) || (rep[0] > 203 && rep[0] < 208) || rep[0] == 105 || rep[0] == 73 {
					tableauX = append(tableauX, 236)
					tableauX = append(tableauX, 237)
					tableauX = append(tableauX, 238)
					tableauX = append(tableauX, 239)
					tableauX = append(tableauX, 204)
					tableauX = append(tableauX, 205)
					tableauX = append(tableauX, 206)
					tableauX = append(tableauX, 207)
					tableauX = append(tableauX, 105)
					tableauX = append(tableauX, 73)
					break
				} else if (rep[0] > 241 && rep[0] < 247) || (rep[0] > 209 && rep[0] < 215) || rep[0] == 111 || rep[0] == 79 {
					tableauX = append(tableauX, 242)
					tableauX = append(tableauX, 243)
					tableauX = append(tableauX, 244)
					tableauX = append(tableauX, 245)
					tableauX = append(tableauX, 246)
					tableauX = append(tableauX, 210)
					tableauX = append(tableauX, 211)
					tableauX = append(tableauX, 212)
					tableauX = append(tableauX, 213)
					tableauX = append(tableauX, 214)
					tableauX = append(tableauX, 111)
					tableauX = append(tableauX, 79)
					break
				} else if (rep[0] > 248 && rep[0] < 253) || (rep[0] > 216 && rep[0] < 221) || rep[0] == 117 || rep[0] == 85 {
					tableauX = append(tableauX, 249)
					tableauX = append(tableauX, 250)
					tableauX = append(tableauX, 251)
					tableauX = append(tableauX, 252)
					tableauX = append(tableauX, 217)
					tableauX = append(tableauX, 218)
					tableauX = append(tableauX, 219)
					tableauX = append(tableauX, 220)
					tableauX = append(tableauX, 117)
					tableauX = append(tableauX, 85)
					break
				} else {
					if letter[0] > 64 && letter[0] < 91 {
						tableauX = append(tableauX, letter[0])
						tableauX = append(tableauX, letter[0]+32)
						break
					} else if letter[0] > 96 && letter[0] < 123 {
						tableauX = append(tableauX, letter[0])
						tableauX = append(tableauX, letter[0]-32)
						break
					}
					tableauX = append(tableauX, rep[0])
					break
				}
			}
		}
		if len(tableauX) == 0 {
			tableauX = append(tableauX, rep[0])
		}
	}
	//fmt.Println(tableauX)
	return rep[0], tableauX, lucky, isALetter, repsentence, isInvalid
}

func CheckAccents(min bool, maj bool, tableau []byte, word []byte, attempts int, letter string) int {
	tableauX := []byte{}
	var lettertest byte
	var isALetter bool
	var sentence []byte
	var isInvalid bool
	//for {
	lucky := 0
	lettertest, tableauX, lucky, isALetter, sentence, isInvalid = EnterLetter(letter, tableauX, lucky)
	if isInvalid {
		attempts = wrong(attempts)
		fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	}
	if !isALetter {
		if len(sentence) > len(word) {
			fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			attempts = wrong(attempts)
			attempts = wrong(attempts)
			//continue
		}
		for i := 0; i < len(sentence); i++ {
			if sentence[i]-32 == word[i] {
				sentence[i] = word[i]
			}
			if sentence[i]+32 == word[i] {
				sentence[i] = word[i]
			}
			//e
			if ((sentence[i] < 236 && sentence[i] > 231) || (sentence[i] < 204 && sentence[i] > 199) || sentence[i] == 101 || sentence[i] == 69) && ((word[i] < 236 && word[i] > 231) || (word[i] < 204 && word[i] > 199) || word[i] == 101 || word[i] == 69) {
				sentence[i] = word[i]
			}
			//a
			if ((sentence[i] < 232 && sentence[i] > 223) || (sentence[i] < 200 && sentence[i] > 191) || sentence[i] == 97 || sentence[i] == 65) && ((word[i] < 232 && word[i] > 223) || (word[i] < 200 && word[i] > 191) || word[i] == 97 || word[i] == 65) {
				sentence[i] = word[i]
			}
			//i
			if ((sentence[i] < 240 && sentence[i] > 235) || (sentence[i] < 208 && sentence[i] > 203) || sentence[i] == 105 || sentence[i] == 73) && ((word[i] < 240 && word[i] > 235) || (word[i] < 208 && word[i] > 203) || word[i] == 105 || word[i] == 73) {
				sentence[i] = word[i]
			}
			//o
			if ((sentence[i] < 247 && sentence[i] > 239) || (sentence[i] < 215 && sentence[i] > 209) || sentence[i] == 111 || sentence[i] == 79) && ((word[i] < 247 && word[i] > 239) || (word[i] < 215 && word[i] > 209) || word[i] == 111 || word[i] == 79) {
				sentence[i] = word[i]
			}
			//u
			if ((sentence[i] < 253 && sentence[i] > 248) || (sentence[i] < 221 && sentence[i] > 216) || sentence[i] == 117 || sentence[i] == 85) && ((word[i] < 253 && word[i] > 248) || (word[i] < 221 && word[i] > 216) || word[i] == 117 || word[i] == 85) {
				sentence[i] = word[i]
			}
		}
	}
	if string(sentence) == string(word) {
		if ChooseFile() {
			PrintArtTable(word, min)
		} else {
			PrintTable(word)
			fmt.Println()
		}
		fmt.Println("Congrats !")
		os.Exit(0)
	} else if len(sentence) > 1 {
		fmt.Println("ddddddddddddddddddddddddddddddd")
		if attempts == 1 {
			wrong(attempts)
			os.Exit(0)
		} else if attempts == 2 {
			attempts = wrong(attempts)
			attempts = wrong(attempts)
			os.Exit(0)
		} else {
			attempts = wrong(attempts)
			attempts = wrong(attempts)
			fmt.Println("ccccccccccccccccccccccccccccccccccccccccccccccccc")
			//continue
		}
	}
	if lucky == 1 {
		//continue
	}
	if maj && (lettertest < 123 && lettertest > 96) {
		lettertest = lettertest - 32
	}
	if min && (lettertest < 91 && lettertest > 64) {
		lettertest = lettertest + 32
	}
	if min {
		// accent de 'e'
		if lettertest == 'e' || lettertest == 'è' || lettertest == 'é' || lettertest == 'ê' || lettertest == 'ë' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'e', compteur)
			tableau, compteur = Check(tableau, word, 'è', compteur)
			tableau, compteur = Check(tableau, word, 'é', compteur)
			tableau, compteur = Check(tableau, word, 'ê', compteur)
			tableau, compteur = Check(tableau, word, 'ë', compteur)
			if compteur == 5 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'E' en minuscune
		if lettertest == 'E' || lettertest == 'È' || lettertest == 'É' || lettertest == 'Ê' || lettertest == 'Ë' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'e', compteur)
			tableau, compteur = Check(tableau, word, 'è', compteur)
			tableau, compteur = Check(tableau, word, 'é', compteur)
			tableau, compteur = Check(tableau, word, 'ê', compteur)
			tableau, compteur = Check(tableau, word, 'ë', compteur)
			if compteur == 5 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'a'
		if lettertest == 'a' || lettertest == 'à' || lettertest == 'á' || lettertest == 'â' || lettertest == 'ã' || lettertest == 'ä' || lettertest == 'å' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'a', compteur)
			tableau, compteur = Check(tableau, word, 'à', compteur)
			tableau, compteur = Check(tableau, word, 'á', compteur)
			tableau, compteur = Check(tableau, word, 'â', compteur)
			tableau, compteur = Check(tableau, word, 'ã', compteur)
			tableau, compteur = Check(tableau, word, 'ä', compteur)
			tableau, compteur = Check(tableau, word, 'å', compteur)
			if compteur == 7 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					if ChooseFile() {
						PrintArtTable(tableau, min)
					} else {
						PrintTable(tableau)
						fmt.Println()
					}
				}
			}
		}
		// accent de 'A' en minuscule
		if lettertest == 'A' || lettertest == 'À' || lettertest == 'Á' || lettertest == 'Â' || lettertest == 'Ã' || lettertest == 'Ä' || lettertest == 'Å' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'a', compteur)
			tableau, compteur = Check(tableau, word, 'à', compteur)
			tableau, compteur = Check(tableau, word, 'á', compteur)
			tableau, compteur = Check(tableau, word, 'â', compteur)
			tableau, compteur = Check(tableau, word, 'ã', compteur)
			tableau, compteur = Check(tableau, word, 'ä', compteur)
			tableau, compteur = Check(tableau, word, 'å', compteur)
			if compteur == 7 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'i'
		if lettertest == 'i' || lettertest == 'ì' || lettertest == 'í' || lettertest == 'î' || lettertest == 'ï' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'i', compteur)
			tableau, compteur = Check(tableau, word, 'ì', compteur)
			tableau, compteur = Check(tableau, word, 'í', compteur)
			tableau, compteur = Check(tableau, word, 'î', compteur)
			tableau, compteur = Check(tableau, word, 'ï', compteur)
			if compteur == 5 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'I' en minuscule
		if lettertest == 'I' || lettertest == 'Ì' || lettertest == 'Í' || lettertest == 'Î' || lettertest == 'Ï' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'i', compteur)
			tableau, compteur = Check(tableau, word, 'ì', compteur)
			tableau, compteur = Check(tableau, word, 'í', compteur)
			tableau, compteur = Check(tableau, word, 'î', compteur)
			tableau, compteur = Check(tableau, word, 'ï', compteur)
			if compteur == 5 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'o'
		if lettertest == 'o' || lettertest == 'ò' || lettertest == 'ó' || lettertest == 'ô' || lettertest == 'õ' || lettertest == 'ö' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'o', compteur)
			tableau, compteur = Check(tableau, word, 'ò', compteur)
			tableau, compteur = Check(tableau, word, 'ó', compteur)
			tableau, compteur = Check(tableau, word, 'ô', compteur)
			tableau, compteur = Check(tableau, word, 'õ', compteur)
			tableau, compteur = Check(tableau, word, 'ö', compteur)
			if compteur == 6 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'O' en minuscule
		if lettertest == 'O' || lettertest == 'Ò' || lettertest == 'Ó' || lettertest == 'Ô' || lettertest == 'Õ' || lettertest == 'Ö' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'o', compteur)
			tableau, compteur = Check(tableau, word, 'ò', compteur)
			tableau, compteur = Check(tableau, word, 'ó', compteur)
			tableau, compteur = Check(tableau, word, 'ô', compteur)
			tableau, compteur = Check(tableau, word, 'õ', compteur)
			tableau, compteur = Check(tableau, word, 'ö', compteur)
			if compteur == 6 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'u'
		if lettertest == 'u' || lettertest == 'ù' || lettertest == 'ú' || lettertest == 'û' || lettertest == 'ü' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'u', compteur)
			tableau, compteur = Check(tableau, word, 'ù', compteur)
			tableau, compteur = Check(tableau, word, 'ú', compteur)
			tableau, compteur = Check(tableau, word, 'û', compteur)
			tableau, compteur = Check(tableau, word, 'ü', compteur)
			if compteur == 5 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'U' en minuscule
		if lettertest == 'U' || lettertest == 'Ù' || lettertest == 'Ú' || lettertest == 'Û' || lettertest == 'Ü' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'u', compteur)
			tableau, compteur = Check(tableau, word, 'ù', compteur)
			tableau, compteur = Check(tableau, word, 'ú', compteur)
			tableau, compteur = Check(tableau, word, 'û', compteur)
			tableau, compteur = Check(tableau, word, 'ü', compteur)
			if compteur == 5 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
	} else if maj {
		// accent de 'E'
		if lettertest == 'E' || lettertest == 'È' || lettertest == 'É' || lettertest == 'Ê' || lettertest == 'Ë' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'E', compteur)
			tableau, compteur = Check(tableau, word, 'È', compteur)
			tableau, compteur = Check(tableau, word, 'É', compteur)
			tableau, compteur = Check(tableau, word, 'Ê', compteur)
			tableau, compteur = Check(tableau, word, 'Ë', compteur)
			if compteur == 5 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'e' en majuscule
		if lettertest == 'e' || lettertest == 'è' || lettertest == 'é' || lettertest == 'ê' || lettertest == 'ë' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'E', compteur)
			tableau, compteur = Check(tableau, word, 'È', compteur)
			tableau, compteur = Check(tableau, word, 'É', compteur)
			tableau, compteur = Check(tableau, word, 'Ê', compteur)
			tableau, compteur = Check(tableau, word, 'Ë', compteur)
			if compteur == 5 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'A'
		if lettertest == 'A' || lettertest == 'À' || lettertest == 'Á' || lettertest == 'Â' || lettertest == 'Ã' || lettertest == 'Ä' || lettertest == 'Å' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'A', compteur)
			tableau, compteur = Check(tableau, word, 'À', compteur)
			tableau, compteur = Check(tableau, word, 'Á', compteur)
			tableau, compteur = Check(tableau, word, 'Â', compteur)
			tableau, compteur = Check(tableau, word, 'Ã', compteur)
			tableau, compteur = Check(tableau, word, 'Ä', compteur)
			tableau, compteur = Check(tableau, word, 'Å', compteur)
			if compteur == 7 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'a' en majuscule
		if lettertest == 'a' || lettertest == 'à' || lettertest == 'á' || lettertest == 'â' || lettertest == 'ã' || lettertest == 'ä' || lettertest == 'å' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'A', compteur)
			tableau, compteur = Check(tableau, word, 'À', compteur)
			tableau, compteur = Check(tableau, word, 'Á', compteur)
			tableau, compteur = Check(tableau, word, 'Â', compteur)
			tableau, compteur = Check(tableau, word, 'Ã', compteur)
			tableau, compteur = Check(tableau, word, 'Ä', compteur)
			tableau, compteur = Check(tableau, word, 'Å', compteur)
			if compteur == 7 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'I'
		if lettertest == 'I' || lettertest == 'Ì' || lettertest == 'Í' || lettertest == 'Î' || lettertest == 'Ï' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'I', compteur)
			tableau, compteur = Check(tableau, word, 'Ì', compteur)
			tableau, compteur = Check(tableau, word, 'Í', compteur)
			tableau, compteur = Check(tableau, word, 'Î', compteur)
			tableau, compteur = Check(tableau, word, 'Ï', compteur)
			if compteur == 5 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'i' en majuscule
		if lettertest == 'i' || lettertest == 'ì' || lettertest == 'í' || lettertest == 'î' || lettertest == 'ï' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'I', compteur)
			tableau, compteur = Check(tableau, word, 'Ì', compteur)
			tableau, compteur = Check(tableau, word, 'Í', compteur)
			tableau, compteur = Check(tableau, word, 'Î', compteur)
			tableau, compteur = Check(tableau, word, 'Ï', compteur)
			if compteur == 5 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'O'
		if lettertest == 'O' || lettertest == 'Ò' || lettertest == 'Ó' || lettertest == 'Ô' || lettertest == 'Õ' || lettertest == 'Ö' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'O', compteur)
			tableau, compteur = Check(tableau, word, 'Ò', compteur)
			tableau, compteur = Check(tableau, word, 'Ó', compteur)
			tableau, compteur = Check(tableau, word, 'Ô', compteur)
			tableau, compteur = Check(tableau, word, 'Õ', compteur)
			tableau, compteur = Check(tableau, word, 'Ö', compteur)
			if compteur == 6 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'o' en majuscule
		if lettertest == 'o' || lettertest == 'ò' || lettertest == 'ó' || lettertest == 'ô' || lettertest == 'õ' || lettertest == 'ö' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'O', compteur)
			tableau, compteur = Check(tableau, word, 'Ò', compteur)
			tableau, compteur = Check(tableau, word, 'Ó', compteur)
			tableau, compteur = Check(tableau, word, 'Ô', compteur)
			tableau, compteur = Check(tableau, word, 'Õ', compteur)
			tableau, compteur = Check(tableau, word, 'Ö', compteur)
			if compteur == 6 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'U'
		if lettertest == 'U' || lettertest == 'Ù' || lettertest == 'Ú' || lettertest == 'Û' || lettertest == 'Ü' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'U', compteur)
			tableau, compteur = Check(tableau, word, 'Ù', compteur)
			tableau, compteur = Check(tableau, word, 'Ú', compteur)
			tableau, compteur = Check(tableau, word, 'Û', compteur)
			tableau, compteur = Check(tableau, word, 'Ü', compteur)
			if compteur == 5 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
		// accent de 'u' en majuscule
		if lettertest == 'u' || lettertest == 'ù' || lettertest == 'ú' || lettertest == 'û' || lettertest == 'ü' {
			compteur := 0
			tableau, compteur = Check(tableau, word, 'U', compteur)
			tableau, compteur = Check(tableau, word, 'Ù', compteur)
			tableau, compteur = Check(tableau, word, 'Ú', compteur)
			tableau, compteur = Check(tableau, word, 'Û', compteur)
			tableau, compteur = Check(tableau, word, 'Ü', compteur)
			if compteur == 5 {
				attempts = wrong(attempts)
			} else {
				if ChooseFile() {
					PrintArtTable(tableau, min)
				} else {
					PrintTable(tableau)
					fmt.Println()
				}
			}
		}
	}
	if lettertest != 'e' && lettertest != 'a' && lettertest != 'i' && lettertest != 'o' && lettertest != 'u' && (lettertest < 123 && lettertest > 96) || (lettertest < 91 && lettertest > 64) {
		compteur := -5
		tableau, compteur = Check(tableau, word, lettertest, compteur)
		if compteur == -1 && lettertest != 'E' && lettertest != 'A' && lettertest != 'I' && lettertest != 'O' && lettertest != 'U' {
			if ChooseFile() {
				PrintArtTable(tableau, min)
			} else {
				PrintTable(tableau)
				//fmt.Println("tttttttteeeeeeessssssssstttttttttt")
				fmt.Println(letter)
			}
		} else if !maj {
			attempts = wrong(attempts)
			fmt.Println("ffffffffffffffffffffffffffffffffffffffffffff")
		}
	}

	if CheckFin(tableau) {
		os.Exit(0)
	}
	if attempts == 0 {
		os.Exit(0)
	}
	return attempts
}

func Check(tableauV []byte, word []byte, letter byte, compteur int) ([]byte, int) {
	pres := false
	for i := 0; i < len(word); i++ {
		if letter == word[i] {
			compteur = -1
			tableauV[i] = letter
			pres = true
		}
	}

	if !pres {
		compteur += 1
	}
	return tableauV, compteur
}

func CheckFin(tableau []byte) bool {
	for i := 0; i < len(tableau); i++ {
		if tableau[i] == '_' {
			return false
		}
	}
	fmt.Println("Congrats !")
	return true
}

func PrintTable(tableau []byte) {
	for i := 0; i < len(tableau); i++ {
		fmt.Print(string(tableau[i]))
		fmt.Print(string(' '))
	}
	fmt.Println()
	fmt.Println()
}

func PrintHang(attempts int) {
	hang, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Println("No hangman.txt file ! Please make one !")
		os.Exit(0)
	}
	info, _ := os.Stat("hangman.txt")
	size := info.Size()
	arr := make([]byte, size)
	hang.Read(arr)
	hang.Close()
	nb := (attempts - 10) * -1
	fmt.Println(string(arr[0+80*(nb-1)-(1*(nb-1)) : 77+80*(nb-1)-(1*(nb-1))]))
}

func wrong(attempts int) int {
	attempts--
	fmt.Print("Not present in the word, ")
	fmt.Print(attempts)
	fmt.Println(" attempts remaining")
	PrintHang(attempts)
	if attempts <= 0 {
		os.Exit(0)
	}
	return attempts
}

func CheckSave() bool {
	if os.Args[1] == "--startWith" && len(os.Args) == 3 {
		return true
	}
	return false
}

func ChooseFile() bool {
	main := os.Args

	if len(main) > 3 {
		if main[2] == "--letterFile" && main[3] != "" {
			return true
		}
	}
	return false

}

func AsciiArt(letter byte, table [8][]string, min bool) [8][]string {
	file, err := os.Open(os.Args[3])
	if err != nil {
		fmt.Println("No Ascii Art ! Please make one !")
		os.Exit(0)
	}
	var test string
	compteur := 0
	begin := 298
	if min {
		begin = 586
	}
	cpt := 0
	cpt2 := -1
	scanner := bufio.NewScanner(file)
	if letter > 64 && letter < 91 {
		for i := 64; i < 91; i++ {
			if letter == byte(i) {
				break
			}
			compteur++
		}
		//fmt.Println(compteur)
		begin = (begin + (8 * compteur)) + 1*compteur
		end := (begin + 8)
		//fmt.Println(begin)
		for scanner.Scan() {
			cpt++
			test = (scanner.Text())

			if cpt > begin {
				cpt2++
				table[cpt2] = append(table[cpt2], test)
				//fmt.Println(test)
			}

			if cpt == end {
				file.Close()
				break
			}
		}
	}
	//a
	if letter == 64 {
		begin = 298
		if min {
			begin = 586
		}
		end := (begin + 8)
		for scanner.Scan() {
			cpt++
			test = (scanner.Text())

			if cpt > begin {
				cpt2++
				table[cpt2] = append(table[cpt2], test)
			}

			if cpt == end {
				file.Close()
				break
			}
		}
	}
	//tiret
	if letter == 95 {
		begin = 118
		end := 122
		for i := 0; i < 4; i++ {
			cpt2++
			table[cpt2] = append(table[cpt2], "         ")
		}
		for scanner.Scan() {
			cpt++
			test = (scanner.Text())

			if cpt > begin {
				cpt2++
				table[cpt2] = append(table[cpt2], test)
			}

			if cpt == end {
				file.Close()
				break
			}
		}
	}
	return table
}

func PrintArtTable(tableau []byte, min bool) {
	var table [8][]string
	for _, elem := range tableau {
		if elem == 95 {
			table = AsciiArt(95, table, min)
		} else if (elem > 231 && elem < 236) || (elem > 199 && elem < 204) {
			table = AsciiArt(68, table, min)
		} else {
			if elem > 64 && elem < 91 {
				table = AsciiArt(elem-1, table, min)
			} else if elem > 95 && elem < 123 {
				table = AsciiArt(elem-33, table, min)
			}
		}

	}
	for i := 0; i < len(table); i++ {
		for k := 0; k < len(table[i]); k++ {
			fmt.Print(table[i][k])
		}
		fmt.Println()
	}
}

func Initialisation(word []byte) (bool, bool, int) {
	min := false
	maj := false
	attempts := 10

	if (word[0] < 91 && word[0] > 64) || (word[0] < 215 && word[0] > 191) || (word[0] < 221 && word[0] > 216) {
		maj = true
	} else if (word[0] < 123 && word[0] > 96) || (word[0] > 223 && word[0] < 247) || (word[0] > 248 && word[0] <= 255) {
		min = true
	}

	return min, maj, attempts
}
