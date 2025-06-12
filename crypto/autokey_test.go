package crypto

import "testing"

func TestAuokeyEncrypt(t *testing.T) {
	tests := []struct {
		plaintext string
		key       string
		expected  string
	}{
		{plaintext: "defend the east wall of the castle", key: "fortification", expected: "iswxvibjexiggzeqpbimoigakmhe"},
		{plaintext: "the quick brown fox jumps over the lazy dog", key: "fortification", expected: "yvvjcnkmbkwkayvbzourcpmsngmsijtksgu"},
		{plaintext: "we are discovered save yourself", key: "fortification", expected: "bsrkmiquchdseahsrzhggwfnicj"},
		{plaintext: "xqzmtplkjhgfdsayuirewnbvc", key: "fortification", expected: "ceqfbutmjaotqpqxgbgpgwibh"},
	}

	for _, test := range tests {
		auto := NewAutokey(test.key)
		ciphertext := auto.Encrypt(test.plaintext)

		if ciphertext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, ciphertext)
		}

	}

}

func TestAutokeyDecrypt(t *testing.T) {
	tests := []struct {
		ciphertext string
		key        string
		expected   string
	}{
		{expected: "defendtheeastwallofthecastle", key: "fortification", ciphertext: "iswxvibjexiggzeqpbimoigakmhe"},
		{expected: "thequickbrownfoxjumpsoverthelazydog", key: "fortification", ciphertext: "yvvjcnkmbkwkayvbzourcpmsngmsijtksgu"},
		{expected: "wearediscoveredsaveyourself", key: "fortification", ciphertext: "bsrkmiquchdseahsrzhggwfnicj"},
		{expected: "xqzmtplkjhgfdsayuirewnbvc", key: "fortification", ciphertext: "ceqfbutmjaotqpqxgbgpgwibh"},
	}

	for _, test := range tests {
		auto := NewAutokey(test.key)
		plaintext := auto.Decrypt(test.ciphertext)

		if plaintext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, plaintext)
		}

	}

}
