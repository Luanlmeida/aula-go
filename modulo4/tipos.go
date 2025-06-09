package main

import (
	"fmt"
)

func teste() {
	fmt.Printf("Type: %T - Value %v\n", true, true)       //bool (true/false)
	fmt.Printf("Type: %T - Value %v\n", "steph", "steph") //string (sequencia de bytes)
	fmt.Printf("Type: %T - Value %v\n", 1, 1)             //int
	fmt.Printf("Type: %T - Value %v\n", "1", "1")         //string
	fmt.Printf("Type: %T - Value %v\n", 1.233, 1.233)     //float
}
