package queue

import (
	"testing"
	U "github.com/jingminglang/gotools/utils"
)

func TestQueue_Push(t *testing.T) {
	node := &Node{Message{Msg:"test",Type:"test",To:"jimila"}}
	QUEUE.Push(node)
	n := QUEUE.Pop()
	U.LogPrintln(n)
	n = QUEUE.Pop()
	U.LogPrintln(n)
}
