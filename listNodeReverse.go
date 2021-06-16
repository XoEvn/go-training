package training

type ListNode struct {
	val int
	next *ListNode

}

func reverse(head *ListNode)(*ListNode){
	if head.next == nil{
		return head
	}
	last := reverse(head.next)
	head.next.next = head
	head.next = nil
	return last
}

// https://labuladong.gitee.io/algo/      算法
// https://lailin.xyz/post/singleton.html go设计模式
