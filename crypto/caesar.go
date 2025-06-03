package crypto

import (
	"regexp"
	"strings"
	"unicode"
)

type Caesar struct {
	plain  string
	cipher string
}

func NewCaesar(shift int) Caesar {
	plain := "abcdefghijklmnopqrstuvwxyz"

	cipher := ""

	for idx := range plain {
		newIdx := (idx + shift) % 26
		cipher += string(plain[newIdx])
	}

	return Caesar{plain: plain, cipher: cipher}

}

func (c *Caesar) Encrypt(plaintext string) string {

	text := c.clean(plaintext)

	res := ""

	for _, t := range text {
		if unicode.IsLetter(rune(t)) {
			idx := strings.Index(c.plain, string(t))
			res += string(c.cipher[idx])
		} else {
			res += string(t)
		}
	}

	return res

}

func (c *Caesar) Decrypt(ciphertext string) string {
	text := c.clean(ciphertext)
	res := ""

	for _, t := range text {
		if unicode.IsLetter(rune(t)) {
			idx := strings.Index(c.cipher, string(t))
			res += string(c.plain[idx])
		} else {
			res += string(t)
		}
	}

	return res
}

func (c *Caesar) clean(text string) string {
	text = strings.TrimSpace(text)

	text = strings.ToLower(text)

	re := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)

	text = re.ReplaceAllString(text, "")

	return text

}
