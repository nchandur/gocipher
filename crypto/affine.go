package crypto

import (
	"regexp"
	"strings"
	"unicode"
)

type Affine struct {
	plain  string
	cipher string
}

func NewAffine(scale, shift int) Affine {
	plain := "abcdefghijklmnopqrstuvwxyz"

	cipher := ""

	for idx := range plain {
		newIdx := ((scale * idx) + shift) % 26
		cipher += string(plain[newIdx])
	}

	return Affine{plain: plain, cipher: cipher}

}

func (a *Affine) Encrypt(plaintext string) string {
	text := a.clean(plaintext)
	res := ""

	for _, t := range text {
		if unicode.IsLetter(rune(t)) {
			idx := strings.Index(a.plain, string(t))
			res += string(a.cipher[idx])
		} else {
			res += string(t)
		}
	}

	return res
}

func (a *Affine) Decrypt(ciphertext string) string {
	text := a.clean(ciphertext)
	res := ""

	for _, t := range text {
		if unicode.IsLetter(rune(t)) {
			idx := strings.Index(a.cipher, string(t))
			res += string(a.plain[idx])
		} else {
			res += string(t)
		}
	}

	return res
}

func (a *Affine) clean(text string) string {
	text = strings.TrimSpace(text)

	text = strings.ToLower(text)

	re := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)

	text = re.ReplaceAllString(text, "")

	return text

}
