package crypto

import (
	"regexp"
	"strings"
	"unicode"
)

type ROT13 struct {
	plain  string
	cipher string
}

func NewROT13() ROT13 {
	plain := "abcdefghijklmnopqrstuvwxyz"
	cipher := "nopqrstuvwxyzabcdefghijklm"

	return ROT13{plain: plain, cipher: cipher}
}

func (r *ROT13) Encrypt(plaintext string) string {
	text := r.clean(plaintext)

	res := ""

	for _, t := range text {
		if unicode.IsLetter(rune(t)) {
			idx := strings.Index(r.plain, string(t))
			res += string(r.cipher[idx])
		} else {
			res += string(t)
		}
	}

	return res

}

func (r *ROT13) Decrypt(ciphertext string) string {
	text := r.clean(ciphertext)
	res := ""

	for _, t := range text {
		if unicode.IsLetter(rune(t)) {
			idx := strings.Index(r.cipher, string(t))
			res += string(r.plain[idx])
		} else {
			res += string(t)
		}
	}

	return res
}

func (r *ROT13) clean(text string) string {
	text = strings.TrimSpace(text)

	text = strings.ToLower(text)

	re := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)

	text = re.ReplaceAllString(text, "")

	return text

}
