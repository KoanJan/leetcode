package main

import (
	"fmt"
)

func main() {
	// test
	obj := Constructor()
	obj.Posttweet(1, 5)
	fmt.Printf("User %d get newsfeed: %v\n", 1, obj.Getnewsfeed(1))
	fmt.Println("User 1 follows User 2")
	obj.Follow(1, 2)
	obj.Posttweet(2, 6)
	fmt.Printf("User %d get newsfeed: %v\n", 1, obj.Getnewsfeed(1))
	fmt.Println("User 1 unfollows User 2")
	obj.Unfollow(1, 2)
	fmt.Printf("User %d get newsfeed: %v\n", 1, obj.Getnewsfeed(1))
}
