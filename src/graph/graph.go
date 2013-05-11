package graph

import (
	"fmt"
)

type Graph struct {
	//keep a list of the nodes
	nodes []*Node

	//the adjancey list of each node (the edges).
	//TODO I'll need to enhance this to support properties on edges
	graph map[uint64][]*Node

	//automatically starts at 0
	nextUID uint64

	//should we log debug statements?
	debug_on bool
}

func NewGraph() *Graph {
	return &Graph{
		graph: make(map[uint64][]*Node)}
}

func NewGraphWithDebug() *Graph {
	return &Graph{
		graph:    make(map[uint64][]*Node),
		debug_on: true}

}

func (g *Graph) AddNode(name string) *Node {
	var newNode *Node = NewNode()
	newNode.name = name
	newNode.id = g.nextUID
	g.nextUID += 1

	if g.debug_on {
		for i := range g.nodes {
			fmt.Printf("before append g.nodes[%v] ptr = %p \n", i, g.nodes[i])
		}
	}

	g.nodes = append(g.nodes, newNode)
	if g.debug_on {
		for i := range g.nodes {
			fmt.Printf("after append g.nodes[%v] ptr = %p \n", i, g.nodes[i])
		}
	}

	g.graph[newNode.id] = []*Node{}

	if g.debug_on {
		fmt.Printf("node.id = %v ptr = %p \n", newNode.id, newNode)
	}
	return newNode
}

/*
	TODO
	This could be a cluster (literally); need to think about how I want handle removes.

	bool - true if the node was found and removed, false indicates it wasn't found

*/
func (g *Graph) RemoveNode(node *Node) (*Node, bool) {
	result, found := g.FindNode(node)

	if !found {
		if g.debug_on {
			fmt.Println("couldn't find node to remove")
		}
		return nil, false
	}

	/*
		we need to do 3 things to remove the node.
		1. we need to remove the node from the node list
		2. we need to remove the node from the node map
		3. we need to remove the node from any edges, which implies the following:

		we need to remove the node from the corresponding edge this may leave
		us a dangling edge, and we'll cascade and remove those as well.
	*/

	//we need to do (1) - remove it from the node list
	g.removeNodeFromNodeList(node)

	//we need to do (2) - remove it from the node map/adjenancy list
	delete(g.graph, node.id)

	//we need to do (3) remove any edges
	g.removeEdgesConnectedToNode(node)

	return result, true 
}

func (g *Graph) removeNodeFromNodeList(candidate *Node) {
	index, found := g.findNodeIndex(candidate)
	if g.debug_on {
		fmt.Printf("removeNodeFromNodeList - index = %v, found = %v \n", index, found)
	}

	if found {
		//we need to pull that one node out of the slice
		g.nodes = append(g.nodes[:index], g.nodes[index+1:]...)
	}
}

/* TODO need to think about the return type */
func (g *Graph) Connect(source *Node, dest *Node) {
	childNodes := g.graph[source.id]

	if Contains(childNodes, dest) {
		//do we already have a connection? If so, ignore
		return
	} else {
		//else we don't have a connection, so we should add it.
		g.graph[source.id] = append(childNodes, dest)
	}
}

func (g *Graph) GetConnectedNodes(node *Node) []*Node {
	if g.debug_on {
		fmt.Printf("GetConnectedNodes = %v \n", g.graph)
	}

	return g.graph[node.id]
}

/*
	Returns the node index in the nodes list

	returns (int, bool)

	the int contains the index assuming the bool is true
	if the bool isn't true, it wasn't found and the index doesn't mean anything
*/
func (g *Graph) findNodeIndex(candidate *Node) (int, bool) {
	found := false
	index := 0
	for i := range g.nodes {
		if g.nodes[i] == candidate {
			found = true
			index = i
		}
	}
	return index, found
}

func (g *Graph) FindNode(candidate *Node) (*Node, bool) {
	index, found := g.findNodeIndex(candidate)
	var result *Node = nil
	if found {
		result = g.nodes[index]
	}
	return result, found
}

func (g *Graph) NodeCount() int {
	return len(g.nodes)
}

func (g *Graph) removeEdgesConnectedToNode(node *Node) {
	for key, _ := range g.graph {
		if g.debug_on {
			fmt.Printf("removingEdgesConnectedToNode key = %v\n", key)
		}
	
		g.removeEdgeFromAdjacencyList(key, node)
	}
}

func (g *Graph) removeEdgeFromAdjacencyList(id uint64, node *Node) {
	//I'm sure this is completely inefficient, but it'll do for now
	edges := g.graph[id]
	
	if g.debug_on {
		fmt.Printf("node to remove = %p edges = %v \n", node, edges)
	}
		
	//I don't slice out the node because I could have multiple edges connecting it
	
	newEdges := make([]*Node, 0)
	for i := range edges {
		if edges[i] != node {
			newEdges = append(newEdges, edges[i])
		}
	}
	
	g.graph[id] = newEdges
	
}

func Contains(nodes []*Node, newNode *Node) bool {
	result := false
	for i := range nodes {
		if nodes[i] == newNode {
			result = true
			break
		}
	}
	return result
}
