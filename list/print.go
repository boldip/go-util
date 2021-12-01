package list

import (
	"fmt"

	"github.com/holizz/terrapin"
)

func (first *Node) Print() {
	for first != nil {
		fmt.Print(first.Value)
		if first.Next != nil {
			fmt.Print(" -> ")
		}
		first = first.Next
	}
	fmt.Println()
	t := terrapin.NewTerrapin(i, terrapin.Position{150.0, 150.0})
}
