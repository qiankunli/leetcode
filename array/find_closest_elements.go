package array

import (
	"fmt"
	"math"
)

func FindClosestElements(arr []int, k int, x int) []int {
	return findClosestElements(arr, k, x)
}

/**
思路1
求绝对值构建新数组
对新数组找第k小的数，快排。问题：比如x=5，那么7和3到5的距离一样，需要进行进一步区分
思路2
找到 x 在arr 中虚拟位置然后左右扩展
**/
// 658
func findClosestElements(arr []int, k int, x int) []int {
	start, end := doFindClosestElements(arr, k, x)
	return arr[start:end]
}

// 真到面试的时候，这个方法除非做过，否则很难想，要调整的细节比较多
func doFindClosestElements(arr []int, k int, x int) (int, int) {
	// 指向第一个大于等于x的位置
	cur := 0
	for arr[cur] < x && cur < len(arr)-1 {
		cur++
	}
	// 左右扩展
	c := 0
	li := cur - 1
	ri := cur
	ld := curd(arr, li, x)
	rd := curd(arr, ri, x)
	for c < k {
		// 如果rd 和 ld 一样，则优先左边
		for rd < ld && c < k && ri < len(arr) {
			c++
			ri++
			rd = curd(arr, ri, x)
		}
		for ld <= rd && c < k && li >= 0 {
			c++
			li--
			ld = curd(arr, li, x)
		}
	}
	return li + 1, ri
}
func curd(arr []int, cur, x int) int {
	if cur < 0 || cur >= len(arr) {
		return math.MaxInt
	}
	return abs(arr[cur] - x)
}

func findClosestElements2(arr []int, k int, x int) []int {

	doFindClosestElements2(arr, k, x, 0, len(arr)-1)
	fmt.Println(arr)
	return arr[:k]
}
func doFindClosestElements2(arr []int, k int, x int, start, end int) {
	if end-start < k {
		return
	}
	left := start
	right := end
	v := abs(arr[left] - x)
	for left < right {
		for abs(arr[right]-x) >= v && left < right {
			right--
		}
		for abs(arr[left]-x) <= v && left < right {
			left++
		}
		swap(arr, left, right)
	}

	swap(arr, start, left)
	// 至此，left 左边比left 小于等于，left 右边比left 大于等于
	if left-start >= k {
		doFindClosestElements2(arr, k, x, start, left)
	} else {
		doFindClosestElements2(arr, k-(left-start), x, left+1, end)
	}
}
