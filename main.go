package main

import (
	"fmt"
)

func main() {

	tree := NewDigitalTree()

	// tree.Add("Amy", " My Baby")
	// tree.Add("Abby", "Sweetie")

	// tree.Delete("Abby")

	tree.Add("Hi", "blank")
	tree.Add("Hit", "blank")
	tree.Add("Hitter", "something")
	tree.ListKeys()

	fmt.Println("||================||")
	tree.Delete("Hitter")
	tree.ListKeys()

	// file, err := os.Create("output.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// jenc := json.NewEncoder(file)

	// err = jenc.Encode(tree)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

// Node ...
type Node struct {
	Path    string      `json:"path,omitempty"`
	Name    string      `json:"name,omitempty"`
	Value   string      `json:"value,omitempty"`
	End     bool        `json:"end,omitempty"`
	Delete  bool        `json:"delete,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
	Parent  *Node
	Child   map[string]*Node `json:"child,omitempty"`
}

func (n *Node) hasChildren() bool {
	if len(n.Child) == 0 {
		return false
	}
	return true

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
			newNode.Parent = node
			node.Child[char] = newNode
			node = node.Child[char]
		}
	}
	node.Payload = payload
	node.End = true
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

// Return the last node of the word
func (dt *DigitalTree) lastNodeOf(word string) *Node {
	node := dt.Root
	for _, letter := range word {
		node = node.Child[string(letter)]
	}
	return node
}

// Delete a word and payload
func (dt *DigitalTree) Delete(word string) {
	fmt.Println("Let's try to delete", word)
	// dt.stepUp(word, dt.lastNodeOf(word), true)

	lastNode := dt.lastNodeOf(word)
	fmt.Println("Last node", lastNode)

}

// ListKeys ...
func (dt *DigitalTree) ListKeys() {

	resultSet := &ResultSet{Name: "Results Set"}
	Walk("", dt.Root, resultSet)

	fmt.Println("Ok, let's see:")
	var count int
	fmt.Println(resultSet.Name)
	for index, word := range resultSet.results {
		fmt.Println(index, word)
		count++
	}
	fmt.Printf("Found %v words\n", count)
}

// ResultSet ...
type ResultSet struct {
	Name    string
	results []string
}

// NewResultSet ...
func NewResultSet(name string) *ResultSet {
	return &ResultSet{
		Name:    name,
		results: []string{},
	}
}

// Walk ...
func Walk(word string, node *Node, results *ResultSet) {
	for char, child := range node.Child {
		if child.End {
			fullWord := word + char
			results.results = append(results.results, fullWord)
			if child.hasChildren() {
				Walk(word+char, child, results)
			}
		} else {
			Walk(word+char, child, results)
		}
	}
}
