package matrix

import (
	"fmt"
	"math/rand"
	"time"
)

func max3(l, r, rr int) int {
	return max(l, max(r, rr))
}
func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}
func min(l, r int) int {
	if l < r {
		return l
	}
	return r
}
func abs(l int) int {
	if l < 0 {
		return -l
	}
	return l
}

func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
func equal(s1, s2 []int) bool {
	if s1 == nil || s2 == nil {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
func contains(s []int, target int) bool {
	if len(s) == 0 {
		return false
	}
	for _, l := range s {
		if l == target {
			return true
		}
	}
	return false
}
func swap(values []int, left, right int) {
	tmp := values[left]
	values[left] = values[right]
	values[right] = tmp
}
func containSlice(ss [][]int, s []int) bool {
	for _, l := range ss {
		if equal(l, s) {
			return true
		}
	}
	return false
}
func join(nums []int, sep string) string {
	if len(nums) == 0 {
		return ""
	}
	if len(nums) == 1 {
		return fmt.Sprintf("%d", nums[0])
	}
	result := ""
	for i := 0; i < len(nums)-1; i++ {
		result += fmt.Sprintf("%d%s", nums[i], sep)
	}
	result += fmt.Sprintf("%d", nums[len(nums)-1])
	return result
}
func TraverseChildSequence(nums []int, handler func(nums []int, start, end int)) {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			handler(nums, i, j)
		}
	}
}

const non_choose = 1
const choose = 2

func TraverseChildArrayWithCondition(nums []int, predict func(nums, bits []int) bool, handler func(nums []int)) {
	bits := make([]int, len(nums))
	traverseChildArrayWithCondition(nums, bits, 0, len(nums)-1,
		func(nums, bits []int, end int) bool {
			array := make([]int, 0)
			array_bits := make([]int, 0)
			for i := 0; i <= end; i++ {
				if bits[i] == choose {
					array = append(array, nums[i])
				}
				array_bits = append(array_bits, bits[i])
			}
			return predict(array, array_bits)
		},
		func(nums []int, bits []int) {
			array := make([]int, 0)
			for i, v := range bits {
				if v == choose {
					array = append(array, nums[i])
				}
			}
			handler(array)
		})
}
func traverseChildArrayWithCondition(nums, bits []int, start, end int,
	predict func(nums, bits []int, end int) bool,
	handler func(nums, bits []int)) {

	if start == end {
		bits[start] = non_choose
		if predict(nums, bits, start) {
			handler(nums, bits)
		}
		bits[start] = choose
		if predict(nums, bits, start) {
			handler(nums, bits)
		}
		return
	}
	bits[start] = non_choose
	if predict(nums, bits, start) {
		traverseChildArrayWithCondition(nums, bits, start+1, end, predict, handler)
	}
	bits[start] = choose
	if predict(nums, bits, start) {
		traverseChildArrayWithCondition(nums, bits, start+1, end, predict, handler)
	}
}

func TraverseChildArray(nums []int, handler func(nums []int)) {
	bits := make([]int, len(nums))
	traverseChildArray(nums, bits, 0, len(nums)-1, func(nums, bits []int) {
		array := make([]int, 0)
		for i, v := range bits {
			if v == choose {
				array = append(array, nums[i])
			}
		}
		handler(array)
	})
}
func traverseChildArray(nums, bits []int, start, end int, handler func(nums, bits []int)) {
	if start == end {
		bits[start] = non_choose
		handler(nums, bits)
		bits[start] = choose
		handler(nums, bits)
		return
	}
	bits[start] = non_choose
	traverseChildArray(nums, bits, start+1, end, handler)
	bits[start] = choose
	traverseChildArray(nums, bits, start+1, end, handler)
}
func RandomInt(start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(end - start)
	random = start + random
	return random
}
func RandomFloat64(start float64, end float64) float64 {
	rand.Seed(time.Now().UnixNano())
	random := rand.Float64()
	random = start + random*(end-start)
	return random
}
