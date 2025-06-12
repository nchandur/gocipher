package crypto

import (
	"regexp"
	"strings"
)

type Autokey struct {
	alphabet string
	key      string
}

func (a *Autokey) removeSpecialCharacters(text string) string {
	text = strings.ToLower(text)
	re := regexp.MustCompile(`[^a-z]+`)

	text = re.ReplaceAllString(text, "")
	return text

}

func NewAutokey(key string) Autokey {
	a := Autokey{alphabet: "abcdefghijklmnopqrstuvwxyz"}
	key = a.removeSpecialCharacters(key)
	a.key = key
	return a

}

func (a *Autokey) Encrypt(plaintext string) string {
	plaintext = a.removeSpecialCharacters(plaintext)
	n := len(plaintext)
	rest := n - len(a.key)

	key := a.key + plaintext[:rest]

	var ciphertext strings.Builder

	for idx, plain := range plaintext {
		p := plain - 'a'
		c := rune(key[idx] - 'a')
		k := (p + c) % 26
		ciphertext.WriteRune(k + 'a')
	}

	return ciphertext.String()
}

func (a *Autokey) Decrypt(ciphertext string) string {
	ciphertext = a.removeSpecialCharacters(ciphertext)
	var plaintext strings.Builder

	key := a.key

	for idx, cipher := range ciphertext {
		k := rune(key[idx] - 'a')
		c := cipher - 'a'
		p := (c - k + 26) % 26
		plaintext.WriteRune(p + 'a')
		key += string(p + 'a')

	}

	return plaintext.String()
}
