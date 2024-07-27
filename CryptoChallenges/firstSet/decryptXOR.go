package firstset

import (
	"slices"
)

// Let KEYSIZE be the guessed length of the key; try values from 2 to (say) 40.
// 1. Write a function to compute the edit distance/Hamming distance between two strings. The Hamming distance is just the number of differing bits. The distance between:
//
//	this is a test
//
// 	and
//
// 	wokka wokka!!!
//
// 	is 37. Make sure your code agrees before you proceed.
// 2. For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second KEYSIZE worth of bytes, and find the edit distance between them. Normalize this result by dividing by KEYSIZE.
// 3. The KEYSIZE with the smallest normalized edit distance is probably the key. You could proceed perhaps with the smallest 2-3 KEYSIZE values. Or take 4 KEYSIZE blocks instead of 2 and average the distances.
// 4. Now that you probably know the KEYSIZE: break the ciphertext into blocks of KEYSIZE length.
// 5. Now transpose the blocks: make a block that is the first byte of every block, and a block that is the second byte of every block, and so on.
// 6. Solve each block as if it was single-character XOR. You already have code to do this.
// 7. For each block, the single-byte XOR key that produces the best looking histogram is the repeating-key XOR key byte for that block. Put them together and you have the key.

func getOnesInByte(b byte) int {
	total := 0
	for {
		total += int(b) % 2
		b /= 2

		if b == 0 {
			break
		}
	}

	return total
}

func EditDistance(s1, s2 []byte) int {
	minLen := min(len(s1), len(s2))
	diff := 0
	for i := 0; i < minLen; i++ {
		b := s1[i] ^ s2[i]
		diff += getOnesInByte(b)
	}
	diff += (len(s1) - minLen) * 8
	diff += (len(s2) - minLen) * 8

	return diff
}

func FindTop5KeySize(message []byte) []int {
	maxKeyLength := len(message)
	distances := map[int]float64{}
	keys := []int{}
	for i := 1; i < maxKeyLength; i++ {
		keys = append(keys, i)
		x, y := message[:i], message[i:min(2*i, len(message)-1)]
		distance := EditDistance(x, y)
		distances[i] = float64(distance) / float64(i*8)
	}

	slices.SortFunc(keys, func(a, b int) int {
		if distances[a] > distances[b] {
			return 1
		}
		if distances[a] == distances[b] {
			return 0
		}
		return -1
	})

	return keys[:min(5, maxKeyLength)]
}

func GetBlocks(message []byte, keyLen int) [][]byte {
	blocks := make([][]byte, keyLen)
	for i, el := range message {
		index := i % keyLen
		s := blocks[index]
		if s == nil {
			s = []byte{}
		}
		s = append(s, el)
		blocks[index] = s
	}

	return blocks
}

func GetKey(blocks [][]byte) []byte {
	keys := []byte{}
	for _, b := range blocks {
		_, key := FrequencyXORCypher(b)
		if key == nil {
			return nil
		}
		keys = append(keys, key[0])
	}

	return keys
}
