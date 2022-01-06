package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("standard.txt")

	var test string
	var table [8][]string
	letter := 95
	compteur := 0
	begin := 298
	cpt := 0
	cpt2 := -1
	scanner := bufio.NewScanner(file)
	if letter > 64 && letter < 91 {
		for i := 64; i < 91; i++ {
			if letter == i {
				break
			}
			compteur++
		}
		fmt.Println(compteur)
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
		begin = 116
		end := (begin + 6)
		cpt2++
		table[cpt2] = append(table[cpt2], "        ")
		cpt2++
		table[cpt2] = append(table[cpt2], "        ")
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
	for i := 0; i < len(table); i++ {
		for k := 0; k < len(table[i]); k++ {
			fmt.Println(string(table[i][k]))
		}
	}
}
