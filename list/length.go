package list

/* Ricevente o target della chiamata del metodo */
func (first *Node) Length() int {
	c := 0
	for first != nil {
		c++
		first = first.Next
	}
	return c
}
