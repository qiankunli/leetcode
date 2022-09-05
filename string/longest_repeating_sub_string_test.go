package string

import (
	"fmt"
	"testing"
)

func TestLongestRepeatingSubstring(t *testing.T) {
	//fmt.Println(LongestRepeatingSubstring("abcd"))
	fmt.Println(LongestRepeatingSubstring("abbaba"))

	//fmt.Println(LongestRepeatingSubstring("aaaaa"))
}

func TestLongestRepeatingSubstringWithPrefix(t *testing.T) {
	fmt.Println(LongestRepeatingSubstringWithPrefix("aaaaa", 3))
}
