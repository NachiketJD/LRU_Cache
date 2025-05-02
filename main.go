import (
	"container/list"
	"fmt"
)

type Node struct {
	Data   int
	KeyPtr *list.Element
}

type LRUCache struct {
	Queue    *list.List
	Items    map[int]*Node
	Capacity int
}