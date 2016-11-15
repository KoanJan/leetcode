package main

import (
	"fmt"
)

func main() {
	// test
	var (
		val       int
		intervals []Interval
	)
	obj := Constructor()

	val = 1
	obj.Addnum(val)
	intervals = obj.Getintervals()
	fmt.Println(intervals)

	val = 1
	obj.Addnum(val)
	intervals = obj.Getintervals()
	fmt.Println(intervals)

	// val = 7
	// obj.Addnum(val)
	// intervals = obj.Getintervals()
	// fmt.Println(intervals)

	// val = 2
	// obj.Addnum(val)
	// intervals = obj.Getintervals()
	// fmt.Println(intervals)

	// val = 6
	// obj.Addnum(val)
	// intervals = obj.Getintervals()
	// fmt.Println(intervals)
}
