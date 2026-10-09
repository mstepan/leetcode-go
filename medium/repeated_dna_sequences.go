package medium

const mask uint32 = (1 << 20) - 1

func FindRepeatedDnaSequences(s string) []string {

	if len(s) < 10 {
		return nil
	}

	var hash uint32 = 0

	for _, ch := range s[:10] {
		hash = (hash << 2) | encodeChar(ch)
	}

	hashToStr := make(map[uint32]string)
	hashToStr[hash] = s[0:10]

	duplicatedSubstrings := make(map[string]struct{})

	for start, last := 1, 10; last < len(s); start, last = start+1, last+1 {

		hash = ((hash << 2) | encodeChar(rune(s[last]))) & mask

		cur_str := s[start : last+1]

		if _, exists := hashToStr[hash]; exists {
			duplicatedSubstrings[cur_str] = struct{}{}
		} else {
			hashToStr[hash] = cur_str
		}

	}

	var result []string

	for single_sub := range duplicatedSubstrings {
		result = append(result, single_sub)
	}

	return result
}

// A => 00, 0
// C => 01, 1
// T => 10, 2
// G => 11, 3
func encodeChar(ch rune) uint32 {
	switch ch {
	case 'A':
		return 0
	case 'C':
		return 1
	case 'T':
		return 2
	case 'G':
		return 3
	default:
		panic("Unexpected character found")
	}
}

// func decode(hash uint32) string {

// 	var decoded_str []rune

// 	for i := 0; i < 10; i++ {
// 		decoded_str = append(decoded_str, decodeChar(hash&0x03))
// 		hash >>= 2
// 	}

// 	reverse(decoded_str)

// 	return string(decoded_str)
// }

// func decodeChar(value uint32) rune {
// 	switch value {
// 	case 0:
// 		return 'A'
// 	case 1:
// 		return 'C'
// 	case 2:
// 		return 'T'
// 	case 3:
// 		return 'G'
// 	default:
// 		panic("Can't decode value to characeter")
// 	}
// }

// func reverse(chars []rune) {
// 	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
// 		chars[i], chars[j] = chars[j], chars[i]

// 	}
// }
