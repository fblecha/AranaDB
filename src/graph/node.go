package graph

import (

)

/*
	A basic node, we'll get fancy later.
*/
type Node struct {

	//the id for a node should be unique to a given graph
	id uint64
	
	//the label for the node
	name string
	
	/* 
		TODO
		I wonder in the long run about the feasiblity of only supporting string values.
	*/
	properties map[string] string
}

func NewNode() *Node {
	return &Node{properties: make(map[string]string)}
}


/* TODO need to think about the return type */
func (n Node) AddProperty(name string, value string) {
	n.properties[name] = value
}

func (n Node) GetProperty(name string) string {
	return n.properties[name]
}

/* TODO need to think about the return type */
func (n Node) RemoveProperty(name string) {
	delete(n.properties, name)
}


