package tree

// 大顶堆
type Heap struct {
	array    []int
	len      int
	capacity int
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		array:    make([]int, capacity),
		capacity: capacity,
	}
}

// 将新插入的节点放在完全二叉树最后的位置，再和父节点比较。如果new节点比父节点小，那么交换两者。交换之后，继续和新的父节点比较…… 直到new节点不比父节点小，或者new节点成为根节点。
func (h *Heap) Push(v int) {
	if h.len > h.capacity {
		// 查找堆中的最小元素，如果比smallest更大，则替换，否则直接放弃
		// 最小元素只会出现在叶子上
		smallIndex, small := smallest(h.array)
		if small >= v {
			return
		}
		h.array[smallIndex] = v
		fixDown(h.array, smallIndex)
		h.len++
		return
	}
	// 放在末尾
	h.array[h.len] = v
	fixDown(h.array, h.len)
	h.len++
}
func (h *Heap) Peek() int {
	if h.len == 0 {
		panic("heap is empty")
	}
	return h.array[0]
}

// 删除操作只能删除根节点。根节点删除后，我们会有两个子树，我们需要基于它们重构堆。 让最后一个节点last成为新的节点，从而构成一个新的二叉树。再将last节点不断的和子节点比较。如果last节点比两个子节点中小的那一个大，则和该子节点交换。直到last节点不大于任一子节点都小，或者last节点成为叶节点。
func (h *Heap) Pop() int {
	// 移除根元素
	value := h.array[0]
	fixUp(h.array, 0)
	h.len--
	return value
}
func smallest(array []int) (int, int) {

	return 0, 0
}

// i 移走，从子节点找节点替代
func fixUp(array []int, i int) {
	left := array[2*i+1]
	right := array[2*(i+1)]
	if left >= right {
		array[i] = left
		fixUp(array, 2*(i+1))
	} else {
		array[i] = right
		fixUp(array, 2*(i+1)+1)
	}
}

// i新加入，上浮到合适的位置
func fixDown(array []int, i int) {
	var parent int
	if i%2 == 1 {
		// 左叶子
		parent = (i - 1) / 2
	} else {
		parent = i / 2
	}
	if array[parent] > array[i] {
		return
	}
	swap(array, parent, i)
	fixDown(array, parent)

}
