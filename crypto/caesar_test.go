package crypto

import "testing"

func TestCaesarEncrypt(t *testing.T) {
	tests := []struct {
		plaintext string
		shift     int
		expected  string
	}{
		{plaintext: "attack at dawn", shift: 5, expected: "fyyfhp fy ifbs"},
		{plaintext: "the quick brown fox jumps over the lazy dog", shift: 2, expected: "vjg swkem dtqyp hqz lworu qxgt vjg ncba fqi"},
		{plaintext: "meet me at 5pm behind the shed", shift: 0, expected: "meet me at 5pm behind the shed"},
	}

	for _, test := range tests {
		caesar := NewCaesar(test.shift)

		ciphertext := caesar.Encrypt(test.plaintext)

		if ciphertext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, ciphertext)
		}

	}

}

func TestCaesarDecrypt(t *testing.T) {
	tests := []struct {
		ciphertext string
		shift      int
		expected   string
	}{
		{ciphertext: "fyyfhp fy ifbs", shift: 5, expected: "attack at dawn"},
		{ciphertext: "vjg swkem dtqyp hqz lworu qxgt vjg ncba fqi", shift: 2, expected: "the quick brown fox jumps over the lazy dog"},
		{ciphertext: "meet me at 5pm behind the shed", shift: 0, expected: "meet me at 5pm behind the shed"},
	}

	for _, test := range tests {
		caesar := NewCaesar(test.shift)

		plaintext := caesar.Decrypt(test.ciphertext)

		if plaintext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, plaintext)
		}

	}

}
