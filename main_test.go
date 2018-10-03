package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestListKeys(t *testing.T) {
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

func TestAdd(t *testing.T) {
	tree := NewDigitalTree()
	tree.Add("Amy", "Baby")
	tree.Add("Abby", "Baby")

	if !tree.Root.Child["A"].Child["m"].Child["y"].End {
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
