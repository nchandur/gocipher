package crypto

import "testing"

func TestAffineEncrypt(t *testing.T) {
	tests := []struct {
		plaintext string
		scale     int
		shift     int
		expected  string
	}{
		{plaintext: "hello world", scale: 1, shift: 0, expected: "hello world"},
		{plaintext: "Defend the east wall of the castle", scale: 3, shift: 10, expected: "twzwxt pfw wkmp ykrr az pfw qkmprw"},
		{plaintext: "meet me at 5pm behind the shed", scale: 5, shift: 2, expected: "kwwt kw ct 5zk hwlqpr tlw olwr"},
	}

	for _, test := range tests {
		affine := NewAffine(test.scale, test.shift)

		ciphertext := affine.Encrypt(test.plaintext)

		if test.expected != ciphertext {
			t.Errorf("expected: %s, output: %s", test.expected, ciphertext)
		}

	}

}

func TestAffineDecrypt(t *testing.T) {
	tests := []struct {
		ciphertext string
		scale      int
		shift      int
		expected   string
	}{
		{ciphertext: "hello world", scale: 1, shift: 0, expected: "hello world"},
		{ciphertext: "twzwxt pfw wkmp ykrr az pfw qkmprw", scale: 3, shift: 10, expected: "defend the east wall of the castle"},
		{ciphertext: "kwwt kw ct 5zk hwlqpr tlw olwr", scale: 5, shift: 2, expected: "meet me at 5pm behind the shed"},
	}

	for _, test := range tests {
		affine := NewAffine(test.scale, test.shift)

		plaintext := affine.Decrypt(test.ciphertext)

		if plaintext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, plaintext)
		}

	}

}
