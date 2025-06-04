package crypto

import (
	"strings"
	"unicode"
)

type Polybius struct {
	alphabet string
	key      string
	chars    string
}

func NewPolybius(alphabet, key, chars string) Polybius {
	p := Polybius{}
	_, p.alphabet = p.removeDuplicates(alphabet)
	_, p.chars = p.removeDuplicates(chars)
	p.key = p.keygen(key)
	return p
}

func (p *Polybius) removeDuplicates(word string) (map[rune]bool, string) {
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

func (p *Polybius) keygen(word string) string {
	letterMap, key := p.removeDuplicates(word)

	for _, alpha := range p.alphabet {
		if unicode.IsLetter(alpha) {
			if ok := letterMap[alpha]; !ok {
				key += string(alpha)
			}
		}
	}

	return key
}

func (p *Polybius) Encrypt(plaintext string) string {

	plaintext = strings.ToLower(plaintext)

	res := ""

	charLen := len(p.chars)

	for _, plain := range plaintext {
		if unicode.IsLetter(plain) {
			idx := strings.Index(p.key, string(plain))
			if idx >= 0 {
				res += string(p.chars[idx/charLen]) + string(p.chars[idx%charLen])
			}
		}
	}

	return res
}

func (p *Polybius) Decrypt(ciphertext string) string {
	res := ""

	n := len(ciphertext)

	for i := 0; i < n-1; i = i + 2 {
		xIdx := strings.IndexByte(p.chars, ciphertext[i])
		yIdx := strings.IndexByte(p.chars, ciphertext[i+1])
		idx := (xIdx * 5) + yIdx
		res += string(p.key[idx])
	}

	return res
}
