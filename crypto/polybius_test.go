package crypto

import "testing"

func TestPolybiusEncrypt(t *testing.T) {
	tests := []struct {
		alphabet  string
		key       string
		chars     string
		plaintext string
		expected  string
	}{
		{alphabet: "abcdefghiklmnopqrstuvwxyz", key: "gvyztbraocfmexduhplkniwsq", chars: "ABCDE", plaintext: "defend the east castle of the wall", expected: "CECCCACCEACEAEDBCCCCBCEDAEBEBCEDAEDDCCBDCAAEDBCCECBCDDDD"},
		{alphabet: "abcdefghiklmnopqrstuvwxyz", key: "zoqmyxuihabewntrdclvgkfps", chars: "01234", plaintext: "far over the misty mountains cold", expected: "42143001342130241321031244240403011123241412234432013331"},
		{alphabet: "abcdefghiklmnopqrstuvwxyz", key: "tlworqgvmhepskaxzydcinufb", chars: "PQRST", plaintext: "hello world", expected: "QTRPPQPQPSPRPSPTPQSS"},
	}

	for _, test := range tests {
		poly := NewPolybius(test.alphabet, test.key, test.chars)

		ciphertext := poly.Encrypt(test.plaintext)

		if ciphertext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, ciphertext)
		}
	}

}

func TestPolybiusDecrypt(t *testing.T) {
	tests := []struct {
		alphabet   string
		key        string
		chars      string
		ciphertext string
		expected   string
	}{
		{alphabet: "abcdefghiklmnopqrstuvwxyz", key: "gvyztbraocfmexduhplkniwsq", chars: "ABCDE", expected: "defendtheeastcastleofthewall", ciphertext: "CECCCACCEACEAEDBCCCCBCEDAEBEBCEDAEDDCCBDCAAEDBCCECBCDDDD"},
		{alphabet: "abcdefghiklmnopqrstuvwxyz", key: "zoqmyxuihabewntrdclvgkfps", chars: "01234", expected: "faroverthemistymountainscold", ciphertext: "42143001342130241321031244240403011123241412234432013331"},
		{alphabet: "abcdefghiklmnopqrstuvwxyz", key: "tlworqgvmhepskaxzydcinufb", chars: "PQRST", expected: "helloworld", ciphertext: "QTRPPQPQPSPRPSPTPQSS"},
	}

	for _, test := range tests {
		poly := NewPolybius(test.alphabet, test.key, test.chars)

		plaintext := poly.Decrypt(test.ciphertext)

		if plaintext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, plaintext)
		}

	}

}
