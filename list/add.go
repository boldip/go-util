package list

func (first *Node) AddFirst(x string) *Node {
	var nuovo *Node

	nuovo = new(Node)
	nuovo.Value = x
	nuovo.Next = first

	return nuovo
}

func (first *Node) AddInOrder(x string) *Node {
	var nuovo *Node

	nuovo = new(Node)
	nuovo.Value = x

	// Cerca dove inserirlo e fai sartoria con i puntatori
	var prev, curr *Node
	prev = nil
	curr = first
	// Esco dal ciclo quando curr==nil oppure la prima volta che curr.Value>=x
	for curr != nil && curr.Value < x {
		prev = curr
		curr = curr.Next
	}
	if prev == nil {
		return first.AddFirst(x)
	}
	prev.Next = nuovo
	nuovo.Next = curr
	return first
}
