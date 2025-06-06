package crypto

import "strings"

type Columnar struct {
	key string
}

func NewColumnar(key string) Columnar {
	c := Columnar{}

	key = strings.ReplaceAll(key, " ", "")
	_, c.key = c.removeDuplicates(key)

	return c
}

func (c *Columnar) removeDuplicates(word string) (map[rune]bool, string) {
	letterMap := make(map[rune]bool)
	str := ""

	for _, w := range word {
		if _, val := letterMap[w]; !val {
			letterMap[w] = true
			str += string(w)
		}
	}

	return letterMap, str
}

func (c *Columnar) Encrypt(plaintext string) string {
	plaintext = strings.ToLower(plaintext)
	plaintext = strings.ReplaceAll(plaintext, " ", "")

	var res string

	return res
}

func (c *Columnar) Decrypt(ciphertext string) string {
	ciphertext = strings.ToLower(ciphertext)
	ciphertext = strings.ReplaceAll(ciphertext, " ", "")

	var res string

	return res
}
