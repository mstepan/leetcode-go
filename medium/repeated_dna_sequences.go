package medium

const mask uint32 = (1 << 20) - 1

// 187. Repeated DNA Sequences
//
// https://leetcode.com/problems/repeated-dna-sequences/description/
func FindRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}

	var hash uint32 = 0

	for i := range 10 {
		hash = (hash << 2) | encodeByte(s[i])
	}

	hashToStr := make(map[uint32]string)
	hashToStr[hash] = s[0:10]

	duplicatedSubstrings := make(map[string]struct{})

	for start, last := 1, 10; last < len(s); start, last = start+1, last+1 {

		hash = ((hash << 2) | encodeByte(s[last])) & mask

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
func encodeByte(ch byte) uint32 {
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
