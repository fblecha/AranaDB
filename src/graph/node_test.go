package graph

import (
	"testing"
	//"fmt"
)


func TestSimpleNodeCreation(t *testing.T) {
	a := NewNode()
	b := NewNode()
	
	a.id = 1
	b.id = 2
	
	a.name = "Bob"
	b.name = "Bob"
	
	if a.name != b.name {
		t.Errorf("Expected node names to be equal")
	}
	
	if a.id != 1 || b.id != 2 {
		t.Errorf("Expected ids to be set to 1 or 2")
	}
	
}


func TestNodePropertyOperations(t *testing.T) {
	a := NewNode()
	a.properties["last_name"] = "Smith"
	a.properties["age"] = "21"
	
	if "Smith" != a.properties["last_name"] { 
		t.Errorf("Expected last_name to be Smith, was %v", a.properties["last_name"])
	}
	if "21" != a.properties["age"] {
		t.Errorf("Expected age to be 21")
	}
	


}
