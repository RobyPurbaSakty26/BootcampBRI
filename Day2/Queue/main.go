package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	items  *list.List
} 


func (q Queue) Enqueue(items any) any{
 q.items.PushBack(items)
 data := q.items
 return data

}

func (q Queue) Denqueue() any{
	if q.items.Front() == nil {
		return -1
	}

	front := q.items.Front()

	val := q.items.Remove(front)
	return val
}

func main() {


	q := Queue{list.New()}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println(q.Enqueue(59))


	
	fmt.Println(q.Denqueue(), q.Denqueue(), q.Denqueue())
}