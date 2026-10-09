package medium

import (
	"slices"
	"testing"
)

func TestFindRepeatedDnaSequencesNormalCase(t *testing.T) {
	assertDnaResults(t, "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"})
	assertDnaResults(t, "AAAAAAAAAAAAA", []string{"AAAAAAAAAA"})
}

func TestFindRepeatedDnaSequencesSmallString(t *testing.T) {
	assertDnaResults(t, "AT", nil)
}

func TestFindRepeatedDnaSequencesEmptyString(t *testing.T) {
	assertDnaResults(t, "", nil)
}

func assertDnaResults(t *testing.T, str string, expected []string) {

	actual := FindRepeatedDnaSequences(str)
	slices.Sort(actual)
	slices.Sort(expected)

	if !slices.Equal(actual, expected) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
