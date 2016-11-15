package main

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	set  map[int]int
	data []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{set: map[int]int{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, existed := this.set[val]; existed {
		return false
	}
	this.set[val] = len(this.data)
	this.data = append(this.data, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if ser, existed := this.set[val]; existed {
		tmp := this.data[len(this.data)-1]
		this.data = this.data[0 : len(this.data)-1]
		if tmp != val {
			this.data[ser] = tmp
			this.set[tmp] = this.set[val]
		}
		delete(this.set, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return this.data[random.Intn(len(this.data))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
