package firstset

import (
	"bytes"
	"slices"
)

const Precision = 0.05

var mostFrequentLetters = []byte{'e', 'a', 'i', 'o', 't', 'n', 's', 'l', 'c', ' '}

var mostFrequentLetters = map[byte]float32{
	'e': 0.111607,
	'a': 0.08496,
	'r': 0.075809,
	'i': 0.075448,
	'o': 0.071635,
	't': 0.069509,
	'n': 0.066544,
	's': 0.057351,
	'l': 0.054893,
	'c': 0.045388,
}

func XOR(input1, input2 []byte) []byte {
	maxLen := max(len(input1), len(input2))
	out := []byte{}

	for i := 0; i < maxLen; i++ {
		if i >= len(input1) {
			out = append(out, input2[i])
		} else if i >= len(input2) {
			out = append(out, input1[i])
		} else {
			b := input1[i] ^ input2[i]
			out = append(out, b)
		}
	}
	return out
}

func GetFrequency(message []byte) map[byte]float32 {
	freqs := make(map[byte]float32)
	for _, b := range message {
		freqs[b]++
	}

	length := len(message)
	for b, val := range freqs {
		freqs[b] = val / float32(length)
	}

	return freqs
}

func Top5Chars(freqs map[byte]float32) []byte {
	chars := []byte{}
	for b := range freqs {
		chars = append(chars, b)
	}
	if len(chars) == 0 {
		return chars
	}
	slices.SortStableFunc(chars, func(a, b byte) int {
		val1, val2 := freqs[a], freqs[b]
		if val1 > val2 {
			return -1
		} else if val1 == val2 {
			return 0
		} else {
			return 1
		}
	})
	sliceLen := min(len(chars), 5)

	return chars[:sliceLen]
}

func IsEnglish(input []byte) bool {
	for _, c := range mostFrequentLetters {
		if !slices.Contains(input, c) {
			return false
		}
	}
	return true
}

func FrequencyXORCypher(message []byte) ([]byte, []byte) {
	for char := 0; char < 128; char++ {
		key := bytes.Repeat([]byte{byte(char)}, len(message))
		decrypted := XOR(message, key)

		if IsEnglish(bytes.ToLower(decrypted)) {
			return decrypted, key
		}
	}

	return message, nil
}
