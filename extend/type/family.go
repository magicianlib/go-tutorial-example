package main

import "fmt"

type Parent struct {
	Name string
}

type Children struct {
	// Name string
	Age uint8
	Parent
}

func main() {

	var children = Children{
		Age: uint8(18),
		// Name: "Smith",
		Parent: Parent{
			Name: "Bob",
		},
	}

	fmt.Println("children.Age:", children.Age)
	fmt.Println("children.Name:", children.Name)
	fmt.Println("children.Parent.Name:", children.Parent.Name)
}
