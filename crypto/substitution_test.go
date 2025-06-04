package crypto

import "testing"

func TestSubstitutionEncrypt(t *testing.T) {
	tests := []struct {
		plaintext string
		key       string
		expected  string
	}{
		{plaintext: "defend the east wall of the castle", key: "phqgiumeaylnofdxjkrcvstzwb", expected: "giuifg cei iprc tpnn du cei qprcni"},
		{plaintext: "the quick brown fox jumps over the lazy dog", key: "jzvswkbxpmndloahqrtyuicfge", expected: "yxw qupvn zraco kaf mulht aiwr yxw djeg sab"},
		{plaintext: "we are discovered save yourself", key: "euoqhmnwrgjacktfidyszlbxvp", expected: "bh edh qryotlhdhq yelh vtzdyham"},
		{plaintext: "xqzmtplkjhgfdsayuirewnbvc", key: "dnfkpqzsjcbroiwtvxyuhaegml", expected: "gvloutrbcszqkydmhjxpeinaf"},
	}

	for _, test := range tests {
		sub := NewSubstitution(test.key)

		ciphertext := sub.Encrypt(test.plaintext)

		if ciphertext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, ciphertext)
		}

	}
}

func TestSubstitutionDecrypt(t *testing.T) {
	tests := []struct {
		ciphertext string
		key        string
		expected   string
	}{
		{expected: "defend the east wall of the castle", key: "phqgiumeaylnofdxjkrcvstzwb", ciphertext: "giuifg cei iprc tpnn du cei qprcni"},
		{expected: "the quick brown fox jumps over the lazy dog", key: "jzvswkbxpmndloahqrtyuicfge", ciphertext: "yxw qupvn zraco kaf mulht aiwr yxw djeg sab"},
		{expected: "we are discovered save yourself", key: "euoqhmnwrgjacktfidyszlbxvp", ciphertext: "bh edh qryotlhdhq yelh vtzdyham"},
		{expected: "xqzmtplkjhgfdsayuirewnbvc", key: "dnfkpqzsjcbroiwtvxyuhaegml", ciphertext: "gvloutrbcszqkydmhjxpeinaf"},
	}

	for _, test := range tests {
		sub := NewSubstitution(test.key)

		plaintext := sub.Decrypt(test.ciphertext)

		if plaintext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, plaintext)
		}

	}

}
