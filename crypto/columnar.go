package crypto

import (
	"sort"
	"strings"
)

type Columnar struct {
	key string
}

type keyCharacter struct {
	key   rune
	index int
}

func NewColumnar(key string) Columnar {
	col := Columnar{}
	key = col.removeDuplicates(key)
	col.key = key

	return col

}

func (c *Columnar) padder(text string) string {

	for len(text)%len(c.key) != 0 {
		text += " "
	}

	return text
}

func (c *Columnar) removeDuplicates(text string) string {
	seen := make(map[rune]bool)

	res := make([]rune, 0, len(text))

	for _, char := range text {
		if !seen[char] {
			seen[char] = true
			res = append(res, char)
		}
	}
	return string(res)

}

func (c *Columnar) Encrypt(plaintext string) string {
	kLen := len(c.key)

	plaintext = c.padder(plaintext)

	rows := len(plaintext) / kLen

	mat := make([][]rune, rows)

	for i := range rows {
		mat[i] = []rune(plaintext[i*kLen : (i+1)*kLen])
	}

	keyChars := make([]keyCharacter, kLen)

	for idx, char := range c.key {
		keyChars[idx] = keyCharacter{key: char, index: idx}
	}

	sort.Slice(keyChars, func(i, j int) bool {
		return keyChars[i].key < keyChars[j].key
	})

	var ciphertext strings.Builder

	for _, kc := range keyChars {
		for row := range rows {
			ciphertext.WriteRune(mat[row][kc.index])
		}
	}

	return ciphertext.String()

}

func (c *Columnar) Decrypt(ciphertext string) string {
	keyLen := len(c.key)

	rows := len(ciphertext) / keyLen

	keyChars := make([]keyCharacter, keyLen)
	for i, c := range c.key {
		keyChars[i] = keyCharacter{c, i}
	}
	sort.Slice(keyChars, func(i, j int) bool {
		return keyChars[i].key < keyChars[j].key
	})

	columns := make([][]rune, keyLen)
	start := 0
	for _, kc := range keyChars {
		columns[kc.index] = []rune(ciphertext[start : start+rows])
		start += rows
	}

	var plaintext strings.Builder
	for row := range rows {
		for col := range keyLen {
			plaintext.WriteRune(columns[col][row])
		}
	}

	return strings.TrimSpace(plaintext.String())
}
