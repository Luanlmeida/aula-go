package main

import (
	"fmt"
)

func tipos() {
	fmt.Printf("Type: %T - Value %v\n", true, true)       //bool (true/false)
	fmt.Printf("Type: %T - Value %v\n", "steph", "steph") //string (sequencia de bytes)
	fmt.Printf("Type: %T - Value %v\n", "1", "1")         //string (sequencia de bytes)
	fmt.Printf("Type: %T - Value %v\n", 1, 1)             //int
	fmt.Printf("Type: %T - Value %v\n", 1.233, 1.233)     //float (float64/float32) - decimal
}
