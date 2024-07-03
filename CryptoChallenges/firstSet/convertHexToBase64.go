package firstset

import (
	"errors"
)

const AlphabetLength = 26
const NumeralLength = 10

func encodeHex(b byte) (byte, error) {
	var out byte
	if ib := int(b); ib < 10 {
		out = byte(int('0') + ib)
	} else if int(b) < 16 {
		out = byte(int('a') + ib)
	} else {
		return 0, errors.New("invalid byte")
	}

	return out, nil
}

func decodeHex(b byte) (byte, error) {
	var out byte
	if ib := int(b); ib >= int('0') && ib <= int('9') {
		out = byte(ib - int('0'))
	} else if ib >= int('a') && ib <= int('z') {
		out = byte(ib - int('a'))
	} else {
		return 0, errors.New("invalid hex")
	}

	return out, nil
}

func ToHex(input []byte) []byte {
	if len(input) == 0 {
		return input
	}
	out := []byte{}
	for _, b := range input {
		first, _ := encodeHex(b >> 4)
		second, _ := encodeHex(b & 0b00001111)
		out = append(out, first, second)
	}
	return out
}

func FromHex(hx []byte) ([]byte, error) {
	if len(hx) == 0 {
		return hx, nil
	}
	if len(hx) == 1 {
		decoded, err := decodeHex(hx[0])
		if err != nil {
			return nil, err
		}
		return []byte{decoded}, nil
	}
	out := []byte{}
	for i := 0; i < len(hx)-1; i += 2 {
		first, err := decodeHex(hx[i])
		if err != nil {
			return out, err
		}
		second, err := decodeHex(hx[i+1])
		if err != nil {
			return out, err
		}
		newByte := (first << 4) | (second)
		out = append(out, newByte)
	}
	return out, nil
}

func encodeBase64(b byte) (byte, error) {
	if int(b) >= 64 {
		return 0, errors.New("invalid byte")
	}

	var out byte
	if ib := int(b); ib < AlphabetLength {
		out = byte('A' + ib)
	} else if ib < AlphabetLength+AlphabetLength {
		out = byte('a' + ib - AlphabetLength)
	} else if ib < AlphabetLength+AlphabetLength+NumeralLength {
		out = byte('0' + ib - AlphabetLength - AlphabetLength)
	} else if ib == 64-2 { // last two
		out = '+'
	} else {
		out = '/'
	}

	return out, nil
}

func buildExtractor(n int) byte {
	if n <= 0 {
		return 0
	}
	var b byte = 1
	for i := 0; i < n; i++ {
		b = (b << 1) + 1
	}
	b = b << byte(8-n)

	return b
}

func assembleBase64Byte(remainder, b byte) (byte, byte) {
	if b == 0 {
		return remainder, 0
	}

	lengthB1 := byteLen(remainder)
	remainingLen := 6 - lengthB1
	firstB := remainder << byte(remainingLen)
	extractor := buildExtractor(remainingLen)
	secondB := (b & extractor) >> byte(8-remainingLen)

	remainingBits := b & (extractor ^ buildExtractor(8))
	final := firstB + secondB

	return final, remainingBits
}

func ToBase64(input []byte) ([]byte, error) {
	if len(input) == 0 {
		return input, nil
	}
	out := []byte{}
	remainder := byte(0)
	var b1 byte
	for _, b := range input {
		b1, remainder = assembleBase64Byte(remainder, b)
		encodedByte, err := encodeBase64(b1)
		if err != nil {
			return out, err
		}
		out = append(out, encodedByte)
	}
	return out, nil
}

func FromBase64(b64 []byte) []byte {
	return nil
}

func byteLen(b byte) int {
	length := 0
	for b != 0b00000000 {
		b = b >> 1
		length++
	}

	return length
}
