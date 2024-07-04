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
		out = byte(int('a') + ib - 10)
	} else {
		return 0, errors.New("invalid byte")
	}

	return out, nil
}

func decodeHex(b byte) (byte, error) {
	var out byte
	if ib := int(b); ib >= '0' && ib <= '9' {
		out = byte(ib - '0')
	} else if ib >= 'a' && ib <= 'f' {
		out = byte(ib - 'a' + NumeralLength)
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

func encodeBytesTripletsToBase64(input []byte) ([]byte, error) {
	out := []byte{}
	switch len(input) {
	case 0:
		return input, nil
	case 1:
		b := input[0]
		first, err := encodeBase64(b >> 2)
		if err != nil {
			return nil, err
		}
		second, err := encodeBase64(b & 0b00000011)
		if err != nil {
			return nil, err
		}
		out = append(out, first, second)
	case 2:
		b1, b2 := input[0], input[1]
		first, err := encodeBase64(b1 >> 2)
		if err != nil {
			return nil, err
		}
		second, err := encodeBase64(((b1 & 0b00000011) << 4) + (b2 >> 4))
		if err != nil {
			return nil, err
		}
		third, err := encodeBase64(b2 & 0b00001111)
		if err != nil {
			return nil, err
		}
		out = append(out, first, second, third)
	case 3:
		b1, b2, b3 := input[0], input[1], input[2]
		first, err := encodeBase64(b1 >> 2)
		if err != nil {
			return nil, err
		}
		second, err := encodeBase64(((b1 & 0b00000011) << 4) + (b2 >> 4))
		if err != nil {
			return nil, err
		}
		third, err := encodeBase64((b2&0b00001111)<<2 + (b3 >> 6))
		if err != nil {
			return nil, err
		}
		fourth, err := encodeBase64(b3 & 0b00111111)
		if err != nil {
			return nil, err
		}
		out = append(out, first, second, third, fourth)
	default:
		return nil, errors.New("Too many bytes, max three")
	}

	return out, nil
}

func ToBase64(input []byte) ([]byte, error) {
	var toEncode []byte
	out := []byte{}
	for len(input) > 2 {
		toEncode, input = input[:3], input[3:]
		encodedBytes, err := encodeBytesTripletsToBase64(toEncode)
		if err != nil {
			return out, err
		}
		out = append(out, encodedBytes...)
	}
	encodedBytes, err := encodeBytesTripletsToBase64(input)
	if err != nil {
		return out, err
	}
	out = append(out, encodedBytes...)

	return out, nil
}
