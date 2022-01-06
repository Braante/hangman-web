package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func main() {
	word := (WordChoose())
	maj := false
	min := false
	/* --- cheat code --- */
	fmt.Println(string((word)))
	if ChooseFile() {
		fmt.Println("Vous avez choisi l'édition deluxe.")
	} else if !ChooseFile() {
		fmt.Println("Vous avez choisi l'édition standard.")
	}
	fmt.Println("Good Luck, you have 10 attempts.")
	attempts := 10
	lword := len(word)
	tableau := []byte{}
	for i := 0; i < lword; i++ {
		tableau = append(tableau, '_')
	}
	var letter byte
	var stockLetter []byte
	var compteur1 int
	if (word[0] < 91 && word[0] > 64) || (word[0] < 215 && word[0] > 191) || (word[0] < 221 && word[0] > 216) {
		maj = true
	} else if (word[0] < 123 && word[0] > 96) || (word[0] > 223 && word[0] < 247) || (word[0] > 248 && word[0] <= 255) {
		min = true
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
	ABC2(attempts, letter)
	PrintTable(tableau)
	fmt.Println()
	CheckAccents(min, maj, tableau, word, attempts, string(letter))

}

func WordChoose() []byte {
	var n int
	cpt := 0
	cptmot := 0
	index := 0
	args := os.Args[1]
	var word []byte
	listword, _ := ioutil.ReadFile(args)
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

func EnterLetter(tableauX []byte, lucky int) (byte, []byte, int, bool, []byte, bool) {
	isALetter := false
	stopgame := false
	var sentence []byte
	var repsentence []byte
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose: ")
	letter, _ := reader.ReadString('\n')
	tempstr := string(letter)
	var rep []byte
	for _, wordabc := range tempstr {
		rep = append(rep, byte(wordabc))
	}
	bytestop := []byte("STOP")
	if len(rep) > 3 {
		isALetter = false
		for i := 0; i < len(letter)-2; i++ {
			sentence = append(sentence, letter[i])
		}
		tempsentence := string(sentence)
		for _, wordabc := range tempsentence {
			repsentence = append(repsentence, byte(wordabc))
		}
		for i := 0; i < len(bytestop); i++ {
			if bytestop[i] == repsentence[i] {
				stopgame = true
			} else {
				break
			}
		}
	}
	for i := 0; i < len(tableauX); i++ {
		if rep[0] == tableauX[i] {
			fmt.Println("You have already tried this letter.")
			lucky = 1
			//fmt.Println(tableauX)
			return rep[0], tableauX, lucky, isALetter, repsentence, stopgame
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
				tableauX = append(tableauX, rep[0])
				break
			}
		}
	}
	if len(tableauX) == 0 {
		tableauX = append(tableauX, rep[0])
	}
	fmt.Println(tableauX)
	return rep[0], tableauX, lucky, isALetter, repsentence, stopgame
}

func CheckAccents(min bool, maj bool, tableau []byte, word []byte, attempts int, letter string) {
	tableauX := []byte{}
	var lettertest byte
	var isALetter bool
	var sentence []byte
	var stopgame bool
	for {
		lucky := 0
		lettertest, tableauX, lucky, isALetter, sentence, stopgame = EnterLetter(tableauX, lucky)
		if stopgame {
			var save []byte
			asave, _ := json.Marshal([]byte(tableau))
			for i := 0; i < len(asave); i++ {
				save = append(save, asave[i])
			}
			asave, _ = json.Marshal([]byte(word))
			for i := 0; i < len(asave); i++ {
				save = append(save, asave[i])
			}
			nbattempts := 48
			if attempts != 10 {
				for i := attempts; i > 0; i-- {
					nbattempts++
				}
			}
			//nbattemps 48 = 10 attempts left
			asave, _ = json.Marshal(byte(nbattempts))
			save = append(save, asave[0])
			ioutil.WriteFile("save.txt", save, 0777)
			os.Exit(0)
		}
		if !isALetter {
			if len(sentence) > len(word) {
				attempts = abc(attempts)
				attempts = abc(attempts)
				continue
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
			PrintTable(word)
			fmt.Println("Congrats !")
			os.Exit(0)
		} else if len(sentence) > 1 {
			if attempts == 1 {
				abc(attempts)
				os.Exit(0)
			} else if attempts == 2 {
				attempts = abc(attempts)
				attempts = abc(attempts)
				os.Exit(0)
			} else {
				attempts = abc(attempts)
				attempts = abc(attempts)
				continue
			}
		}
		if lucky == 1 {
			continue
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
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
					attempts = abc(attempts)
				} else {
					PrintTable(tableau)
				}
			}
		}
		if lettertest != 'e' && lettertest != 'a' && lettertest != 'i' && lettertest != 'o' && lettertest != 'u' && (lettertest < 123 && lettertest > 96) || (lettertest < 91 && lettertest > 64) {
			compteur := -5
			tableau, compteur = Check(tableau, word, lettertest, compteur)
			if compteur == -1 {
				PrintTable(tableau)
			} else {
				attempts = abc(attempts)
			}
		}
		if CheckFin(tableau) {
			os.Exit(0)
		}
		if attempts == 0 {
			os.Exit(0)
		}
	}
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
	if !ChooseFile() {
		for i := 0; i < len(tableau); i++ {
			fmt.Print(string(tableau[i]))
			fmt.Print(string(' '))
		}
		fmt.Println()
		fmt.Println()
	} else {

	}
}

func PrintHang(attempts int) {
	hang, _ := os.Open("hangman.txt")
	info, _ := os.Stat("hangman.txt")
	size := info.Size()
	arr := make([]byte, size)
	hang.Read(arr)
	hang.Close()
	nb := (attempts - 10) * -1
	fmt.Println(string(arr[0+80*(nb-1)-(1*(nb-1)) : 77+80*(nb-1)-(1*(nb-1))]))
}

func abc(attempts int) int {
	attempts--
	fmt.Print("Not present in the word, ")
	fmt.Print(attempts)
	fmt.Println(" attempts remaining")
	PrintHang(attempts)
	return attempts
}

func ChooseFile() bool {
	main := os.Args

	if len(main) > 3 {
		if main[2] == "--letterFile" {
			return true
		}
	}
	return false

}

func UseFile(attempts int, letter byte) {
	var count int
	// var begin int
	// var last int
	art, _ := os.Open(os.Args[3])
	info, _ := os.Stat(os.Args[3])
	size := info.Size()
	arr := make([]byte, size)
	art.Read(arr)
	art.Close()
	if letter > 64 && letter < 91 {
		for i := 64; i < 91; i++ {
			if letter == byte(i) {
				break
			}
			count++
		}
	} else if letter > 96 && letter < 123 {
		for i := 96; i < 123; i++ {
			if letter == byte(i) {
				break
			}
			count++
		}
	}
	fmt.Println(count)
	fmt.Println(string(arr[2440+(70*count) : 2520+(70*count)]))
	fmt.Println('\n')
	count = 0
}

func ABC2(attempts int, letter byte) int {
	UseFile(attempts, letter)
	return attempts
}
