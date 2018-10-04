package digitaltree

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestLastLetterFunc(t *testing.T) {
	word := "Hallo"
	found, char := lastLetter(word)
	if found && char == "o" {
		fmt.Println("Hey sweet")
	} else {
		t.Error("lastLetter func failed")
	}

	nextWord := ""
	found1, char := lastLetter(nextWord)
	if found1 {
		t.Error("lastLetter Shouldn't have found anything")
	}
}

func TestAllButLastLetterFunc(t *testing.T) {
	word := "Hallo"
	found, char := allButLastLetter(word)
	if found && char == "Hall" {
		fmt.Println("Hey sweet")
	} else {
		t.Error("allButLastLetter failed")
	}

	nextWord := ""
	found1, _ := allButLastLetter(nextWord)
	if found1 {
		t.Error("lastLetter Shouldn't have found anything")
	}
}

func TestListKeys(t *testing.T) {
	tree := NewDigitalTree()
	tree.Add("Absolute", "Perfection")
	tree.Add("Amy", "In the morning")
	tree.Add("AmyDoodle", "In the evening")

	tree.ListKeys()
}
func BenchmarkListKeys(t *testing.B) {
	tree := NewDigitalTree()

	file, err := os.Open("names.txt")
	if err != nil {
		log.Fatalln("Couldn't open names.txt")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tree.Add(scanner.Text(), "Temp Payload")
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	tree.ListKeys()
}

func TestMoreDelete(t *testing.T) {
	tree := NewDigitalTree()
	tree.Add("AB", "Number One")
	tree.Add("ABC", "Number Two")
	tree.Add("ABCDE", "Number Three")
	tree.Delete("ABCDE")
	found, _ := tree.Find("ABCDE")
	if found {
		t.Error("Shouldn't have been able to find Hit")
	}

}

func TestDeleteKey(t *testing.T) {
	tree := NewDigitalTree()
	tree.Add("Hi", "Number One")
	tree.Add("Hit", "Number Two")
	tree.Add("Hits", "Number Three")
	tree.Delete("Hit")
	found, _ := tree.Find("Hit")
	if found {
		t.Error("Shouldn't have been able to find Hit")
	}

	tree.Delete("Hi")
	found1, _ := tree.Find("Hi")
	if found1 {

		t.Error("Shouldn't have been able to find Hi")
	}

	foundHits, _ := tree.Find("Hits")
	if !foundHits {
		t.Fail()
	}

}

func TestAdd(t *testing.T) {
	tree := NewDigitalTree()
	tree.Add("Absolute", "Perfection")
	tree.Add("Amy", "In the morning")

	found, _ := tree.Find("Amy")
	if !found {
		t.Error("Didn't find Amy")
	}

}
func BenchmarkAdd(t *testing.B) {
	tree := NewDigitalTree()
	tree.Add("Amy", "Baby")

	if !tree.Root.Child["A"].Child["m"].Child["y"].End {
		t.Error("Didn't find Amy")
	}
}

func TestFind(t *testing.T) {
	tree := NewDigitalTree()
	tree.Add("Amy", "Baby")

	found, _ := tree.Find("Amy")
	if !found {
		t.Error("Couldn't find Amy")
	}

	found2, _ := tree.Find("Luda")
	if found2 {
		t.Error("Shouldn't be able to find Luda")
	}

}
func BenchmarkFind(t *testing.B) {
	tree := NewDigitalTree()
	tree.Add("Amy", "Baby")

	found, payload := tree.Find("Amy")
	if found {
		fmt.Println("Found Amy: ", payload)
	} else {
		t.Error("Couldn't find Amy")
	}

	found2, _ := tree.Find("Luda")
	if found2 {
		t.Error("Shouldn't be able to find Luda")
	}
}
