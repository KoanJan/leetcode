package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	var (
		res        []string = []string{}
		start, end int
		cr         string
	)
	if len(nums) == 0 {
		return res
	}

	start = nums[0]
	end = start
	cr = int2Str(start)

	for i := 1; i < len(nums); i++ {
		if nums[i]-end == 1 {
			end = nums[i]
		} else if end == start {
			res = append(res, cr)
			start = nums[i]
			cr = int2Str(start)
			end = start
		} else {
			cr = cr + "->" + int2Str(end)
			res = append(res, cr)
			start = nums[i]
			cr = int2Str(start)
			end = start
		}
	}
	if end > start {
		cr = cr + "->" + int2Str(end)
	}
	res = append(res, cr)
	return res
}

func int2Str(a int) string {
	if a == 0 {
		return "0"
	}
	var (
		tmp      []rune = []rune{}
		aIsMinus bool   = a < 0
		d        []rune = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
		rev      []rune = []rune{}
	)
	if a < 0 {
		a = -a
	}
	for a > 0 {
		tmp = append(tmp, d[a%10])
		a /= 10
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		rev = append(rev, tmp[i])
	}
	if aIsMinus {
		return "-" + string(rev)
	}
	return string(rev)
}

func main() {
	// test
	nums := []int{0, 1}
	fmt.Printf("summary ranges of %v is %v\n", nums, summaryRanges(nums))
}

/**
 * Because of Golang(unicode), here's ACed code with JavaScript
 *
 * /**
 * @param {number[]} nums
 * @return {string[]}
 *
	var summaryRanges = function(nums) {
		var res = [];
		var start, end = 0;
		var cr = "";

		if (nums.length === 0){
			return res;
		}

		start = nums[0];
		end = start;

		for (var i = 1; i < nums.length; i++) {
			if (nums[i] - end === 1){
				end = nums[i];
			}else if (end === start){
				res.push(start+"");
				start = nums[i];
				end = start;
			}else{
				res.push(start + "->" + end + "");
				start = nums[i];
				end = start;
			}
		}

		if (end > start){
			res.push(start + "->" + end + "")
		}else{
			res.push(start + "");
		}
		return res;
	};
 *
 *
*/
