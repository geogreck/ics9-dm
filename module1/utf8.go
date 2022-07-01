package main

func encodeRune(utf32Rune rune) []byte {
	if utf32Rune < 128 {
		return []byte{byte(utf32Rune)}
	} else if utf32Rune < 2048 {
		return []byte{byte(utf32Rune>>6 | 192), byte(utf32Rune&63 | 128)}
	} else if utf32Rune < 65535 {
		return []byte{byte(utf32Rune>>12 | 224), byte(utf32Rune>>6&63 | 128), byte(utf32Rune&63 | 128)}
	} else {
		return []byte{byte(utf32Rune>>18&7 | 240), byte(utf32Rune>>12&63 | 128),
			byte(utf32Rune>>6&63 | 128), byte(utf32Rune&63 | 128)}
	}
}

func encode(utf32 []rune) []byte {
	utf8 := make([]byte, 0)
	for _, rune := range utf32 {
		rune_utf8 := encodeRune(rune)
		utf8 = append(utf8, rune_utf8...)
	}
	return utf8
}

func decodeByte(utf8 []byte, i int) (rune, int) {
	if utf8[i] < 128 {
		return rune(utf8[i]), i + 1
	} else if utf8[i] <= 223 {
		return rune(utf8[i]&31)<<6 | rune(utf8[i+1]&63), i + 2
	} else if utf8[i] <= 239 {
		return rune(utf8[i]&15)<<12 | rune(utf8[i+1]&63)<<6 | rune(utf8[i+2]&63), i + 3
	} else {
		return rune(utf8[i]&7)<<18 | rune(utf8[i+1]&63)<<12 | rune(utf8[i+2]&63)<<6 | rune(utf8[i+3]&63), i + 4
	}
}

func decode(utf8 []byte) []rune {
	utf32 := make([]rune, 0)
	var rune rune
	for i := 0; i < len(utf8); {
		rune, i = decodeByte(utf8, i)
		utf32 = append(utf32, rune)
	}
	return utf32
}

func main() {
}
