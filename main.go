package main

import (
	"fmt"
	"sort"
	"strings"
)

type keyCharacter struct {
	key   rune
	index int
}

func padder(text, key string) string {

	for len(text)%len(key) != 0 {
		text += "X"
	}

	return text
}

func removeDuplicates()

func main() {

	text := "defend the east wall of the castle"
	key := "german"
	kLen := len(key)

	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "")

	text = padder(text, key)

	rows := len(text) / kLen

	mat := make([][]rune, rows)

	for i := range rows {
		mat[i] = []rune(text[i*kLen : (i+1)*kLen])
	}

	keyChars := make([]keyCharacter, kLen)

	for idx, char := range key {
		keyChars[idx] = keyCharacter{key: char, index: idx}
	}

	sort.Slice(keyChars, func(i, j int) bool {
		return keyChars[i].key < keyChars[j].key
	})

	ciphertext := ""

	for _, kc := range keyChars {
		for row := range rows {
			ciphertext += string(mat[row][kc.index])
		}
	}

	fmt.Println(ciphertext)
}
