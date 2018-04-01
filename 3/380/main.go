package main

import (
	"fmt"
)

func main() {
	// test
	var val int
	obj := Constructor()

	val = 0
	fmt.Printf("insert %d\n", val)
	if !obj.Insert(val) {
		fmt.Printf("%d has already existed\n", val)
	}

	val = 0
	fmt.Printf("remove %d\n", val)
	if !obj.Remove(val) {
		fmt.Printf("%d didn't exist\n", val)
	}

	val = -1
	fmt.Printf("insert %d\n", val)
	if !obj.Insert(val) {
		fmt.Printf("%d has already existed\n", val)
	}

	val = 0
	fmt.Printf("remove %d\n", val)
	if !obj.Remove(val) {
		fmt.Printf("%d didn't exist\n", val)
	}

	fmt.Printf("get random number: %d\n", obj.GetRandom())
}
