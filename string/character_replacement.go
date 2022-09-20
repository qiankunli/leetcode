package string

// 找一个字符串，除了k个其它字符外，其它字符都一样。
// 暴力法 就是所有子串，找符合这个要求的
// 窗口，k个长度，如果右边的字符在 窗口里，就扩，如果不在，就缩，记录窗口滑动过程中的最大长度
// 看了题解之后
// 依然没有脱离寻找连续子串的范畴
// 424
func characterReplacement(s string, k int) int {

	n := len(s)
	if n == 0 {
		return 0
	}
	if n <= k {
		return n
	}
	l := 0
	r := 0
	maxC := 0
	cache := make(map[string]int, 0)
	for l <= r && r < n {
		sr := s[r : r+1]
		cache[sr]++
		maxC = max(maxC, cache[sr]) // 窗口里数量最多的字符 的数量
		// 窗口里只能有k-1 个不一样的 其它的字符都是一样的
		if r-l+1-maxC > k { // 窗口只会变大或不变，不缩小，就不用一个 max 值记录窗口长度了
			sl := s[l : l+1]
			cache[sl]--
			l++
		}
		r++ // 最后多加了一次
	}
	return r - l
}
