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
unc Constructor(capacity int) LRUCache {
	return LRUCache{Queue: list.New(), Items: make(map[int]*Node), Capacity: capacity}
}