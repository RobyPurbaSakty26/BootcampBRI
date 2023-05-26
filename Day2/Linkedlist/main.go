package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	Head *Node
}

// Append adds a new with the given data to the end of the linked list
func (ll *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	// liping selama current != nill maka current di ubah manajadi current . next
	current := ll.Head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

// delete dataa
func (ll *LinkedList) Delete(data int){
	if ll.Head == nil {
		return
	}

	if ll.Head.data == data {
		ll.Head = ll.Head.next
		return
	}

	// mencari lokasi data
	current := ll.Head
	previous := current

	// dugunakan unutk menyambuing data yang kosong
	for current != nil && current.data != data {
		previous = current
		current = current.next
	}

	if current == nil {
		return
	}

	previous.next = current.next
}

func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main(){
	ll := &LinkedList{}

	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	ll.Append(40)
	ll.Append(50)

	// ll.Delete(30)

	// ll.Display()
}
