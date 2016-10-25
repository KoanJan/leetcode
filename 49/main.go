package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	var (
		res  [][]string     = [][]string{}
		keys []map[byte]int = []map[byte]int{}
	)
	for i := 0; i < len(strs); i++ {
		// convert
		tmp := []byte(strs[i])
		tmpKeys := map[byte]int{}
		for j := 0; j < len(tmp); j++ {
			tmpKeys[tmp[j]] += 1
		}
		// compare
		found := false
		for j := 0; j < len(keys); j++ {
			found = true
			if len(keys[j]) != len(tmpKeys) {
				found = false
				continue
			}
			for k, v := range keys[j] {
				if tmpKeys[k] != v {
					found = false
					break
				}
			}
			if found {
				res[j] = append(res[j], strs[i])
				break
			}
		}
		// add
		if !found {
			res = append(res, []string{strs[i]})
			keys = append(keys, tmpKeys)
		}
	}
	return res
}

func main() {
	// test
	strs := []string{"tea", "and", "ace", "ad", "eat", "dans"}
	fmt.Printf("strs: %v\ngroupAnagrams produced: %v\n", strs, groupAnagrams(strs))
}
