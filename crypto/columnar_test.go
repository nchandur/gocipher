package crypto

import "testing"

func TestColumnarEncrypt(t *testing.T) {
	tests := []struct {
		plaintext string
		key       string
		expected  string
	}{
		{plaintext: "defend the east wall of the castle", key: "german", expected: "nalcxehwttdttfseeleedsoaxfeahl"},
		{plaintext: "the quick brown fox jumps over the lazy dog", key: "cipher", expected: "tcnmrzuojvlgqrxoeohkfptyeboshdiwueax"},
		{plaintext: "we are discovered save yourself", key: "head", expected: "ecoomvhzgqkwxpeeyxhirfuotaotubnjsrld"},
		{plaintext: "xqzmtplkjhgfdsayuirewnbvc", key: "tail", expected: "qphsinxzlgarbxmkfyevxxtjduwc"},
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
		{expected: "defendtheeastwallofthecastle", key: "german", ciphertext: "nalcxehwttdttfseeleedsoaxfeahl"},
		{expected: "thequickbrownfoxjumpsoverthelazydog", key: "cipher", ciphertext: "tcnmrzuojvlgqrxoeohkfptyeboshdiwueax"},
		{expected: "wearediscoveredsaveyourself", key: "head", ciphertext: "ecoomvhzgqkwxpeeyxhirfuotaotubnjsrld"},
		{expected: "xqzmtplkjhgfdsayuirewnbvc", key: "tail", ciphertext: "qphsinxzlgarbxmkfyevxxtjduwc"},
	}

	for _, test := range tests {
		col := NewColumnar(test.key)

		plaintext := col.Decrypt(test.ciphertext)

		if plaintext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, plaintext)
		}

	}
}
