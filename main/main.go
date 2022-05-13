package main

import (
	"github.com/phantom820/collections/lists"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/lists/list"
	"github.com/phantom820/collections/queues"
	"github.com/phantom820/collections/queues/listqueue"
	"github.com/phantom820/collections/queues/slicequeue"
	"github.com/phantom820/collections/stacks"
	"github.com/phantom820/collections/stacks/liststack"
	"github.com/phantom820/collections/stacks/slicestack"
	"github.com/phantom820/collections/types"
)

func main() {

	var l1 lists.List[types.Integer] = forwardlist.New[types.Integer]()
	var l2 lists.List[types.Integer] = list.New[types.Integer]()
	l1.Back()
	l2.Back()

	var q1 queues.Queue[types.Integer] = listqueue.New[types.Integer]()
	q1.Front()
	var q2 queues.Queue[types.Integer] = slicequeue.New[types.Integer]()
	q2.Front()

	var s1 stacks.Stack[types.Integer] = slicestack.New[types.Integer]()
	var s2 stacks.Stack[types.Integer] = liststack.New[types.Integer]()

}
