package list

import "fmt"

func (first *Node) Print() {
	for first != nil {
		fmt.Print(first.Value)
		if first.Next != nil {
			fmt.Print(" -> ")
		}
		first = first.Next
	}
	fmt.Println()
}
