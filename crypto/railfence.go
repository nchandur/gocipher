package crypto

import "strings"

type RailFence struct {
	rails int
}

func NewRailFence(rails int) RailFence {
	return RailFence{rails: rails}
}

func (r *RailFence) Encrypt(plaintext string) string {
	if r.rails <= 1 {
		return plaintext
	}

	fence := make([][]rune, r.rails)
	for i := range fence {
		fence[i] = make([]rune, len(plaintext))
	}

	row, down := 0, false
	for col, char := range plaintext {
		fence[row][col] = char
		if row == 0 || row == r.rails-1 {
			down = !down
		}
		if down {
			row++
		} else {
			row--
		}
	}

	var result strings.Builder
	for _, r := range fence {
		for _, c := range r {
			if c != 0 {
				result.WriteRune(c)
			}
		}
	}
	return result.String()
}

func (r *RailFence) Decrypt(ciphertext string) string {
	if r.rails <= 1 {
		return ciphertext
	}

	length := len(ciphertext)
	fence := make([][]bool, r.rails)
	for i := range fence {
		fence[i] = make([]bool, length)
	}

	row, down := 0, false
	for col := range length {
		fence[row][col] = true
		if row == 0 || row == r.rails-1 {
			down = !down
		}
		if down {
			row++
		} else {
			row--
		}
	}

	index := 0
	railMatrix := make([][]rune, r.rails)
	for i := range railMatrix {
		railMatrix[i] = make([]rune, length)
		for j := 0; j < length; j++ {
			if fence[i][j] && index < length {
				railMatrix[i][j] = rune(ciphertext[index])
				index++
			}
		}
	}

	var result strings.Builder
	row, down = 0, false
	for col := range length {
		result.WriteRune(railMatrix[row][col])
		if row == 0 || row == r.rails-1 {
			down = !down
		}
		if down {
			row++
		} else {
			row--
		}
	}
	return result.String()
}
