package string

import (
	"fmt"
	"testing"
)

func TestWordBreak2(t *testing.T) {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	sentences := WordBreak2(s, wordDict)
	for _, v := range sentences {
		fmt.Println(v)
	}
}
