package firstset

import (
	"bytes"
	"math"
)

const Precision = 0.8

// var mostFrequentLetters = []byte{'e', 'a', 'i', 'o', ' '}

var mostFrequentLetters = map[byte]float32{
	' ': 0.5,
	'e': 0.111607,
	'a': 0.084966,
	'r': 0.075809,
	'i': 0.075448,
	'o': 0.071635,
	't': 0.069509,
	'n': 0.066544,
	's': 0.057351,
	'l': 0.054893,
	'c': 0.045388,
	'u': 0.036308,
	'd': 0.033844,
	'p': 0.031671,
	'm': 0.030129,
	'h': 0.030034,
	'g': 0.024705,
	'b': 0.020720,
	'f': 0.018121,
	'y': 0.017779,
	'w': 0.012899,
	'k': 0.011016,
	'v': 0.010074,
	'x': 0.002902,
	'z': 0.002722,
	'j': 0.001965,
	'q': 0.001962,
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

func getSimilarity(freqs map[byte]float32) float32 {
	var totalRatio float32 = 1
	count := 0
	length := 0
	for b, f := range freqs {
		length++
		// val := freqs[b]
		val, ok := mostFrequentLetters[b]
		if !ok {
			continue
		}
		count += 1
		ratio := f / val
		totalRatio *= ratio
	}
	if count < min(15, length) {
		return float32(math.Inf(1))
	}
	return abs(float32(1) - totalRatio)
}

func FrequencyXORCypher(message []byte) byte {
	var mostSimilar byte
	var mostSimilarVal float32 = float32(math.Inf(1))
	for key := 0; key < 128; key++ {
		fullKey := bytes.Repeat([]byte{byte(key)}, len(message))
		decrypted := XOR(message, fullKey)

		freq := GetFrequency(decrypted)
		similarity := getSimilarity(freq)
		// fmt.Println(string(byte(key)), similarity)
		if similarity < mostSimilarVal {
			mostSimilar = byte(key)
			mostSimilarVal = similarity
		}
	}

	return mostSimilar
}

func RepeatingKeyXOREncrypt(key []byte, message []byte) []byte {
	if len(message)*len(key) == 0 {
		return message
	}
	out := []byte{}
	for i, b := range message {
		bKey := key[i%len(key)]
		out = append(out, b^bKey)
	}

	return out
}

func abs[N int | float32](x N) N {
	if x < 0 {
		return -x
	}
	return x
}

func getNumOnes(b byte) int {
	total := 0
	for b != 0 {
		total += int(b & 0b00000001)
		b = b >> 1
	}

	return total
}

func EditDistance(input1, input2 []byte) int {
	length := min(len(input1), len(input2))
	distance := 0
	for i := 0; i < length; i++ {
		b1, b2 := input1[i], input2[i]
		diff := b1 ^ b2
		distance += getNumOnes(diff)
	}
	distance += (len(input1) - length) * 8
	distance += (len(input2) - length) * 8

	return distance
}

func GetKeyLengths(message []byte) map[int]float32 {
	maxKeyLength := len(message)
	distances := make(map[int]float32)
	for keyL := 1; keyL < maxKeyLength; keyL++ {
		k1, k2 := message[:keyL], message[keyL:min(2*keyL, len(message)-1)]
		distance := float32(EditDistance(k1, k2)) / float32(8*keyL)
		distances[keyL] = distance
	}

	return distances
}

func GetBlocks(message []byte, keyLen int) [][]byte {
	blocks := make([][]byte, keyLen)
	for i, b := range message {
		index := i % keyLen
		s := blocks[index]
		if s == nil {
			s = []byte{}
		}
		s = append(s, b)
		blocks[index] = s
	}

	return blocks
}
