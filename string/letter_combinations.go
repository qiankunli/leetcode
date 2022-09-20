package string

func LetterCombinations(digits string) []string {
	return letterCombinations(digits)
}

// 17
// 回溯法
var hash map[string][]string
var path []string
var result []string

func letterCombinations(digits string) []string {

	n := len(digits)
	if n == 0 {
		return nil
	}
	hash = make(map[string][]string)
	hash["2"] = []string{"a", "b", "c"}
	hash["3"] = []string{"d", "e", "f"}
	hash["4"] = []string{"g", "h", "i"}
	hash["5"] = []string{"j", "k", "l"}
	hash["6"] = []string{"m", "n", "o"}
	hash["7"] = []string{"p", "q", "r", "s"}
	hash["8"] = []string{"t", "u", "v"}
	hash["9"] = []string{"w", "x", "y", "z"}
	path = make([]string, 0)
	result = make([]string, 0)
	doLetterCombinations(digits, 0)
	return result
}
func doLetterCombinations(digits string, start int) {

	if start >= len(digits) {
		s := join(path)
		result = append(result, s)
		return
	}
	s := digits[start : start+1]
	chars := hash[s]
	for _, v := range chars {
		path = append(path, v)
		doLetterCombinations(digits, start+1)
		path = path[:len(path)-1]
	}
}
func join(path []string) string {
	sum := ""
	for _, s := range path {
		sum += s
	}
	return sum
}
