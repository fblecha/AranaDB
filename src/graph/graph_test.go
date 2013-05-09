package graph

import (
	"testing"
//	"fmt"
)

/*
	For these tests, I'm going to use something that's easy to talk about as a
	graph and at least seems somewhat fun:  Biological Classification

	I vastly prefer to have a collection of small individual tests that are focused
	on specific functionality than a few monster tests.


	So we start off with:

	Life > Domain > Kingdom > Phylum > Class > Order -> Family > Genus > Species

	Since we're testing AranaDB, we'll start off with Class Arachnida, using
	the following:

	There's only one type of (extinct)extinct spider in the Haptopoda Order
	(Class) Arachnida > (Order) Haptopoda > (Family) Plesiosironidae
	> (Genus) Plesiosiro > (species)madeleyi

	This particular species (Plesiosiro madeley) was discovered/described
	by Reginald Pocock (4 March 1863 – 9 August 1947) in 1911.  It's
	location can be listed as (city)Staffordshire, (country)United Kingdom.
*/

func TestSimpleNodeCreationViaGraph(t *testing.T) {
	graph := NewGraph()

	// make a root node that corresponds to the arachnid class
	var arachnida *Node = graph.AddNode("Arachnida")
	var haptopoda *Node = graph.AddNode("Haptopoda")

	if nil == arachnida || nil == haptopoda {
		t.Errorf("Expected nonnil nodes")
	}

	if 2 != graph.NodeCount() {
		t.Errorf("Expected node count to be 2, it was %v ", graph.NodeCount())
	}

}

func TestSimpleConnections(t *testing.T) {
	graph := NewGraph()

	// make a root node that corresponds to the arachnid class
	var arachnida *Node = graph.AddNode("Arachnida")
	var haptopoda *Node = graph.AddNode("Haptopoda")

	if nil == arachnida || nil == haptopoda {
		t.Errorf("Expected nonnil nodes")
	}

	graph.Connect(arachnida, haptopoda)

	connectedNodes := graph.GetConnectedNodes(arachnida)

	if !Contains(connectedNodes, haptopoda) {
		t.Errorf("Expected haptopoda to be connected to arachnida")
	}
}

func TestRemoveUnconnectedNodeFromGraph(t *testing.T) {
	graph := NewGraph()

	// make a root node that corresponds to the arachnid class
	var arachnida *Node = graph.AddNode("Arachnida")
	var haptopoda *Node = graph.AddNode("Haptopoda")
	
	//just have this line so there's a reason for haptopda to exist 
	//and go compiler doesn't complain
	haptopoda.properties["foo"] = "blah"

	removedNode, didRemove := graph.RemoveNode( arachnida )
	
	if didRemove == false || removedNode != arachnida {
		t.Errorf("Expected arachnida to be the removed node")
	}

	if 1 != graph.NodeCount() {
		t.Errorf("Expected 1 node to still be in the graph but got %v \n", graph.NodeCount())
	}
}

func TestRemoveConnectedNode(t *testing.T) {
	graph := NewGraph()

	// make a root node that corresponds to the arachnid class
	var arachnida *Node = graph.AddNode("Arachnida")
	var haptopoda *Node = graph.AddNode("Haptopoda")

	graph.Connect(arachnida, haptopoda)

	removedNode, didRemove := graph.RemoveNode( haptopoda )
	
	if didRemove == false || removedNode != haptopoda {
		t.Errorf("Expected arachnida to be the removed noded")
	}

	if 1 != graph.NodeCount() {
		t.Errorf("Expected 1 node to still be in the graph but got %v \n", graph.NodeCount())
	}
	
	numConnectedNodes := len(graph.GetConnectedNodes(arachnida))
	
	//fmt.Printf("nodes = %v \n", graph.GetConnectedNodes(arachnida))
	
	
	if numConnectedNodes != 0 {
		t.Errorf("Expected 0 connected nodes for arachnida but got %v \n", numConnectedNodes)
	} 
}

/*
func TestRemoveEdge(t *testing.T) {

}
*/

/*
//TODO future functionality
func testEdgePropertyOperations(t *testing.T) {

}
*/

func TestGraphCreation(t *testing.T) {

	graph := NewGraph()
	/* make a root node that corresponds to the arachnid class */
	arachnida := graph.AddNode("Arachnida")
	haptopoda := graph.AddNode("Haptopoda")
	graph.Connect(arachnida, haptopoda)

	plesioironidae := graph.AddNode("Plesiosironidae")
	graph.Connect(haptopoda, plesioironidae)

	plesiosiro := graph.AddNode("Plesiosiro")
	graph.Connect(plesioironidae, plesiosiro)

	madeleyi := graph.AddNode("Plesiosiro madeleyi")
	graph.Connect(plesiosiro, madeleyi)

	discoverer := graph.AddNode("Reginald Pocock")
	discoverer.AddProperty("born", "3/4/1863")
	discoverer.AddProperty("died", "8/9/1947")
	discoverer.AddProperty("nationality", "British")
	discoverer.AddProperty("profession", "zoologist")
	//TODO add edge properties
	//graph.Connect(madeleyi, discoverer, "discovered by")

	madeleyi.AddProperty("city", "Staffordshire")
	madeleyi.AddProperty("country", "United Kingdom")
	madeleyi.AddProperty("date_discovered", "1911")
	madeleyi.AddProperty("extinct", "true")

	//we should have added 6 nodes
	expectedSize := 6
	actualSize := graph.NodeCount()

	if expectedSize != actualSize {
		t.Errorf("Expected graph of size %v, but got %v", expectedSize, actualSize)
	}
}
