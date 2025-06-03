package crypto

import "testing"

func TestAtbashEncrypt(t *testing.T) {
	tests := []struct {
		plaintext string
		expected  string
	}{
		{plaintext: "hello world", expected: "svool dliow"},
		{plaintext: "the quick brown fox jumps over the lazy dog", expected: "gsv jfrxp yildm ulc qfnkh levi gsv ozab wlt"},
		{plaintext: "meet me at 5pm behind the shed", expected: "nvvg nv zg 5kn yvsrmw gsv hsvw"},
		{plaintext: "we are discovered save yourself", expected: "dv ziv wrhxlevivw hzev blfihvou"},
		{plaintext: "xqzmtplkjhgfdsayuirewnbvc", expected: "cjangkopqstuwhzbfrivdmyex"},
	}

	atbash := NewAtbash()

	for _, test := range tests {
		ciphertext := atbash.Encrypt(test.plaintext)

		if ciphertext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, ciphertext)
		}
	}

}

func TestAtbashDecrypt(t *testing.T) {
	tests := []struct {
		ciphertext string
		expected   string
	}{
		{ciphertext: "svool dliow", expected: "hello world"},
		{ciphertext: "gsv jfrxp yildm ulc qfnkh levi gsv ozab wlt", expected: "the quick brown fox jumps over the lazy dog"},
		{ciphertext: "nvvg nv zg 5kn yvsrmw gsv hsvw", expected: "meet me at 5pm behind the shed"},
		{ciphertext: "dv ziv wrhxlevivw hzev blfihvou", expected: "we are discovered save yourself"},
		{ciphertext: "cjangkopqstuwhzbfrivdmyex", expected: "xqzmtplkjhgfdsayuirewnbvc"},
	}

	atbash := NewAtbash()

	for _, test := range tests {
		plaintext := atbash.Decrypt(test.ciphertext)

		if plaintext != test.expected {
			t.Errorf("expected: %s, output: %s", test.expected, plaintext)
		}

	}

}
