package crypto

import (
	"strings"
	"unicode"
)

type Substitution struct {
	alphabet string
	key      string
}

func NewSubstitution(key string) Substitution {
	s := Substitution{alphabet: "abcdefghijklmnopqrstuvwxyz"}
	s.key = s.keygen(key)
	return s
}

func (s *Substitution) removeDuplicates(word string) (map[rune]bool, string) {
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

func (s *Substitution) keygen(word string) string {
	letterMap, key := s.removeDuplicates(word)

	for _, alpha := range s.alphabet {
		if unicode.IsLetter(alpha) {
			if ok := letterMap[alpha]; !ok {
				key += string(alpha)
			}
		}
	}

	return key
}

func (s *Substitution) Encrypt(plaintext string) string {
	plaintext = strings.ToLower(plaintext)

	var res string

	for _, plain := range plaintext {
		if unicode.IsLetter(plain) {
			idx := strings.IndexRune(s.alphabet, plain)
			res += string(s.key[idx])
		} else {
			res += string(plain)
		}
	}

	return res
}

func (s *Substitution) Decrypt(ciphertext string) string {

	ciphertext = strings.ToLower(ciphertext)

	var res string

	for _, cip := range ciphertext {
		if unicode.IsLetter(cip) {
			idx := strings.IndexRune(s.key, cip)
			res += string(s.alphabet[idx])
		} else {
			res += string(cip)
		}
	}

	return res
}
