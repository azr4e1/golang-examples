package firstset

import (
	"bytes"
	"slices"
)

const Precision = 0.05

var mostFrequentLetters = []byte{'e', 'a', 'i', 'o', ' '}

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

func FrequencyXORCypher(message []byte) []byte {
	for char := 0; char < 128; char++ {
		key := bytes.Repeat([]byte{byte(char)}, len(message))
		decrypted := bytes.ToLower(XOR(message, key))
		isContained := true
		for _, c := range mostFrequentLetters {
			if !slices.Contains(decrypted, c) {
				isContained = false
				break
			}
		}

		if isContained {
			return key
		}
	}

	return nil
}
