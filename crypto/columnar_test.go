package crypto

import "testing"

func TestColumnarEncrypt(t *testing.T) {
	tests := []struct {
		plaintext string
		key       string
		expected  string
	}{
		{plaintext: "defend the east wall of the castle", key: "german", expected: "n wfc etslhtd altsee o edea a fht el"},
		{plaintext: "the quick brown fox jumps over the lazy dog", key: "cipher", expected: "tioxs agqbfme d    uve  hcw  tz eknjohy uroprlo "},
		{plaintext: "we are discovered save yourself", key: "head", expected: "  crs rfadoeays eese eulwrivdvoe"},
		{plaintext: "xqzmtplkjhgfdsayuirewnbvc", key: "tail", expected: "qphsin zlgarb mkfyev xtjduwc"},
	}

	for _, test := range tests {
		col := NewColumnar(test.key)

		ciphertext := col.Encrypt(test.plaintext)

		if ciphertext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, ciphertext)
		}

	}
}

func TestColumnarDecrypt(t *testing.T) {
	tests := []struct {
		ciphertext string
		key        string
		expected   string
	}{
		{ciphertext: "n wfc etslhtd altsee o edea a fht el", key: "german", expected: "defend the east wall of the castle"},
		{ciphertext: "tioxs agqbfme d    uve  hcw  tz eknjohy uroprlo ", key: "cipher", expected: "the quick brown fox jumps over the lazy dog"},
		{ciphertext: "  crs rfadoeays eese eulwrivdvoe", key: "head", expected: "we are discovered save yourself"},
		{ciphertext: "qphsin zlgarb mkfyev xtjduwc", key: "tail", expected: "xqzmtplkjhgfdsayuirewnbvc"},
	}

	for _, test := range tests {
		col := NewColumnar(test.key)

		plaintext := col.Decrypt(test.ciphertext)

		if plaintext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, plaintext)
		}

	}
}
