package main

import "fmt"

// Generic function
func FirstElement[T interface{}](items []T) T {
	return items[0]
}

func main() {
	names := []string{"Deepanshu Mehra", "Yogita Mehra", "Divya Mehra"}
	ages := []int{21, 27, 31}

	fmt.Println(FirstElement(names))
	fmt.Println(FirstElement(ages))

	// Using Custom Stack
	stack := Stack[int]{}
	stack.Push(4)
	stack.Push(8)
	stack.Push(11)
	stack.Push(3)

	stringStack := Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")
	stringStack.Push("!!")

	stack.Pop()
	stringStack.Pop()
	fmt.Println("Item :", stack.items)
	fmt.Println("String Stack :", stringStack.items)
}
