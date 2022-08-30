package array

func Rotate(nums []int, k int) {
	rotate(nums, k)
}

// 可以视为 以k个元素为单位，一块的长度为k，整体向右移动一位
// 因为辅助变量只能记录一个值，因此一次只能移动一位，用程序去模拟序列。 A 移到B 之后就必须为B 找下家
// 一个一个移，但是跳跃着

// 189
func rotate(nums []int, k int) {

	if k == 0 {
		return
	}
	n := len(nums)
	if n == 1 {
		return
	}
	if k >= n {
		if k%n == 0 {
			return
		}
		k = k % n
	}
	// 将i 移动到 i+k(nums[i+k]=nums[k]), i+k 移动到 i+k+k, 如果大于n 则对n求余，直到回到i，一轮就结束了(再移动也是重复的)
	// i+1 重新开始新的一轮
	count := 0 // 已经移动了多少快，控制循环结束
	i := 1     // 此次处理的是第几块
	j := 0     // 表示 本轮移动块内第几个元素，一定会小于k
	tmp := nums[j]
	for count < n {
		index := i*k + j
		if index >= n {
			index = index % n
		}
		v := nums[index]
		nums[index] = tmp
		tmp = v
		count++
		if index == j { // 已经处理了所有块，回到了第一块的位置，之前j 已经移动过了， 就不用再移动了
			i = 1
			j++
			tmp = nums[j]
			continue
		}
		i++
	}

}

// n 不是 k 的整数倍
func rotateA(nums []int, k int) {
	n := len(nums)
	count := 0
	i := 1
	j := 0 // 将tmp 转移到i, j 指向i 下一个位置
	tmp := nums[j]
	for count < n {
		index := i*k + j
		if index >= n {
			index = index % n
		}
		v := nums[index]
		nums[index] = tmp
		tmp = v
		count++
		if index == j {
			i = 1
			j++
			tmp = nums[j]
			continue
		}
		i++
	}
}

// n 是 k 的整数倍，将向右移的动作操作了k 次，一次移动 n/k 个
func rotateB(nums []int, k int) {
	n := len(nums)
	blocks := n / k
	// 把tmp 移到index
	for j := 0; j < k; j++ {
		tmp := nums[j]
		for i := 1; i <= blocks; i++ {
			index := i*k + j
			if index >= n {
				index = index % n
			}
			v := nums[index]
			nums[index] = tmp
			tmp = v
		}
	}
}

func rotateC(nums []int, k int) {
	n := len(nums)
	blocks := n / k
	if n%k > 0 {
		blocks += 1
	}
	// 把tmp 移到index
	for j := 0; j < k; j++ {
		tmp := nums[j]
		for i := 1; i <= n; i++ {
			index := i*k + j
			if index > n {
				index = index % n
			}
			v := nums[index]
			nums[index] = tmp
			tmp = v
			if index == j {
				break
			}
		}
	}
}

func rotate2(nums []int, k int) {
	n := len(nums)
	cycle := n / k
	if n%k > 0 {
		cycle += 1
	}
	for j := 0; j < k; j++ {
		tmp := nums[j]
		for i := 1; i <= cycle; i++ {
			index := i*k + j
			if index >= n {
				index = index % n
			}
			v := nums[index]
			nums[index] = tmp
			tmp = v
		}
	}
}

func rotate1(nums []int, k int) {
	for i := 0; i < k; i++ {
		move(nums)
	}
}

func move(nums []int) {
	last := nums[len(nums)-1]
	tmp := nums[0]
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		nums[i] = tmp
		tmp = v
	}
	nums[0] = last
}
