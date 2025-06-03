package crypto

import "testing"

func TestRailFenceEncrypt(t *testing.T) {
	tests := []struct {
		plaintext string
		rails     int
		expected  string
	}{
		{plaintext: "defend the east wall of the castle", rails: 3, expected: "dnhaw tcleedtees alo h atef  tlfes"},
		{plaintext: "meet me at 5pm behind the shed", rails: 5, expected: "maeee tbhh ee  itstm5mn hd pde"},
		{plaintext: "the quick brown fox jumps over the lazy dog", rails: 4, expected: "tioxs aghucrwo p rtlzoeqkbnfjmoeh yd   uve "},
	}

	for _, test := range tests {
		rail := NewRailFence(test.rails)

		ciphertext := rail.Encrypt(test.plaintext)

		if ciphertext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, ciphertext)
		}

	}

}

func TestRailFenceDecrypt(t *testing.T) {
	tests := []struct {
		ciphertext string
		rails      int
		expected   string
	}{
		{ciphertext: "dnhaw tcleedtees alo h atef  tlfes", rails: 3, expected: "defend the east wall of the castle"},
		{ciphertext: "maeee tbhh ee  itstm5mn hd pde", rails: 5, expected: "meet me at 5pm behind the shed"},
		{ciphertext: "tioxs aghucrwo p rtlzoeqkbnfjmoeh yd   uve ", rails: 4, expected: "the quick brown fox jumps over the lazy dog"},
	}

	for _, test := range tests {
		rail := NewRailFence(test.rails)

		plaintext := rail.Decrypt(test.ciphertext)

		if plaintext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, plaintext)
		}

	}

}
