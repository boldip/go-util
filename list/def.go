package list

/* Definisco il tipo di base dei nodi di una lista (concatenata semplice) di stringhe */

type Node struct {
	Value string
	Next  *Node
}
