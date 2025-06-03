package crypto

import "testing"

func TestROT13Encrypt(t *testing.T) {
	tests := []struct {
		plaintext string
		expected  string
	}{
		{plaintext: "hello world", expected: "uryyb jbeyq"},
		{plaintext: "the quick brown fox jumps over the lazy dog", expected: "gur dhvpx oebja sbk whzcf bire gur ynml qbt"},
		{plaintext: "meet me at 5pm behind the shed", expected: "zrrg zr ng 5cz oruvaq gur furq"},
		{plaintext: "we are discovered save yourself", expected: "jr ner qvfpbirerq fnir lbhefrys"},
		{plaintext: "xqzmtplkjhgfdsayuirewnbvc", expected: "kdmzgcyxwutsqfnlhverjaoip"},
	}

	rot := NewROT13()

	for _, test := range tests {
		ciphertext := rot.Encrypt(test.plaintext)

		if ciphertext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, ciphertext)
		}
	}
}

func TestROT13Decrypt(t *testing.T) {
	tests := []struct {
		ciphertext string
		expected   string
	}{
		{ciphertext: "uryyb jbeyq", expected: "hello world"},
		{ciphertext: "gur dhvpx oebja sbk whzcf bire gur ynml qbt", expected: "the quick brown fox jumps over the lazy dog"},
		{ciphertext: "zrrg zr ng 5cz oruvaq gur furq", expected: "meet me at 5pm behind the shed"},
		{ciphertext: "jr ner qvfpbirerq fnir lbhefrys", expected: "we are discovered save yourself"},
		{ciphertext: "kdmzgcyxwutsqfnlhverjaoip", expected: "xqzmtplkjhgfdsayuirewnbvc"},
	}

	rot := NewROT13()

	for _, test := range tests {
		plaintext := rot.Decrypt(test.ciphertext)

		if plaintext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, plaintext)
		}
	}
}
