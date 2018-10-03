package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	tree := NewDigitalTree()

	tree.Add("Amy", " My Baby")
	tree.Add("Abby", "Sweetie")

	tree.ListKeys()

	file, err := os.Create("output.json")
	if err != nil {
		log.Fatal(err)
	}

	jenc := json.NewEncoder(file)

	err = jenc.Encode(tree)
	if err != nil {
		log.Fatal(err)
	}
}

// Node ...
type Node struct {
	Path    string           `json:"path,omitempty"`
	Name    string           `json:"name,omitempty"`
	Value   string           `json:"value,omitempty"`
	End     bool             `json:"end,omitempty"`
	Child   map[string]*Node `json:"child,omitempty"`
	Payload interface{}      `json:"payload,omitempty"`
}

// NewNode returns reference to new Node
func NewNode() *Node {
	return &Node{
		Child: make(map[string]*Node),
	}
}

// DigitalTree ...
type DigitalTree struct {
	Root *Node `json:"root,omitempty"`
}

// NewDigitalTree returns a ref to a new DigitalTree
func NewDigitalTree() *DigitalTree {
	return &DigitalTree{
		Root: &Node{
			Name:  "Root",
			Child: make(map[string]*Node),
		},
	}
}

// Add a word to the DigitalTree
func (dt *DigitalTree) Add(word string, payload interface{}) {
	found, _ := dt.Find(word)
	if !found {
		node := dt.Root
		var path string

		for _, letter := range word {
			char := string(letter)
			_, found := node.Child[char]
			if found {
				path += char
				node = node.Child[char]
			} else {
				newNode := NewNode()
				path += char
				newNode.Path = path
				node.Child[char] = newNode
				node = node.Child[char]
			}
		}
		node.Payload = payload
		node.End = true
	}
}

// Find by key
// @word string, key for lookup
// return: truthy if found, falsy if not found
// return payload if found
func (dt *DigitalTree) Find(word string) (bool, interface{}) {
	node := dt.Root

	for _, letter := range word {
		char := string(letter)
		_, found := node.Child[char]
		if found {
			node = node.Child[char]
		} else {
			return false, nil
		}
	}
	return true, node.Payload
}

// Walk ...
func Walk(word string, node *Node) {

	for char, childnode := range node.Child {
		if childnode.End {
			fullWord := word + char
			fmt.Println(fullWord, childnode.Payload)
		} else {
			Walk(word+char, childnode)
		}
	}
}

// ListKeys ...
func (dt *DigitalTree) ListKeys() {
	Walk("", dt.Root)
}
