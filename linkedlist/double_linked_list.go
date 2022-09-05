package linkedlist

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}
type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func NewDoubleLinkedList() *DoubleLinkedList {
	// 使用一个伪头部（dummy head）和伪尾部（dummy tail）标记界限，这样在添加节点和删除节点的时候就不需要检查相邻的节点是否存在
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return &DoubleLinkedList{
		head: head, // 不指向真实数据
		tail: tail,
	}
}
func (d *DoubleLinkedList) AppendToFirst(node *Node) {
	if node == nil {
		return
	}
	// node 放在第一个
	d.head.Next.Prev = node
	node.Next = d.head.Next
	node.Prev = d.head
	d.head.Next = node
}
func (d *DoubleLinkedList) MoveToFirst(node *Node) {
	if node == nil || node == d.head || d.head.Next == node {
		return
	}

	// 连接node 前后两点
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	// node 放在第一个
	d.head.Next.Prev = node
	node.Next = d.head.Next
	node.Prev = d.head
	d.head.Next = node
}

func (d *DoubleLinkedList) Head() *Node {
	return d.head.Next
}
func (d *DoubleLinkedList) Tail() *Node {
	return d.tail.Prev
}

func (d *DoubleLinkedList) Delete(node *Node) {
	if node == nil || node == d.head {
		return
	}
	// 连接node 前后两点
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
func (d *DoubleLinkedList) IsEmpty() bool {
	return d.head.Next == d.tail
}
