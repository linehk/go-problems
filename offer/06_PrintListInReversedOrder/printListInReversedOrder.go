package printListInReversedOrder

import (
	"container/list"
	"fmt"
	"github.com/linehk/GoProblems/offer/utils"
	"strconv"
)

func printListInReversedOrder(head *list.Element) {
	stack := utils.NewStack()
	cur := head
	for cur != nil {
		stack.Push(cur)
		cur = cur.Next()
	}

	s := ""
	for !stack.Empty() {
		i := stack.Pop()
		n := i.(*list.Element)
		v := n.Value.(int32)
		s += fmt.Sprintf("%s", strconv.Itoa(int(v)))
	}
	fmt.Println(s)
}

func printListInReversedOrderRecursively(head *list.Element) {
	if head != nil {
		if head.Next() != nil {
			printListInReversedOrderRecursively(head.Next())
		}
		fmt.Printf("%v", head.Value)
	}
}
