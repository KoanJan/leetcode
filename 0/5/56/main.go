package main

import "fmt"

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	ln := len(intervals)
	qsort(intervals, 0, ln-1)
	merged := make([]bool, ln)
	result := make([]Interval, 0)
	for i := 0; i < ln; i++ {
		if !merged[i] {
			cur := intervals[i]
			for j := i + 1; j < ln; j++ {
				if !merged[j] && intervals[j].Start <= cur.End {
					if cur.Start > intervals[j].Start {
						cur.Start = intervals[j].Start
					}
					if cur.End < intervals[j].End {
						cur.End = intervals[j].End
					}
					merged[j] = true
				}
			}
			result = append(result, cur)
			merged[i] = true
		}
	}

	return result
}

func qsort(intervals []Interval, left, right int) {
	if left < right {
		initialPivot(intervals, left, right)
		p := partition(intervals, left, right, left)
		qsort(intervals, left, p-1)
		qsort(intervals, p, right)
	}
}

func partition(intervals []Interval, left, right, pivotIndex int) int {
	i, j := left, right
	p := intervals[pivotIndex].Start
	for i < j {
		for i < right && intervals[i].Start <= p {
			i++
		}
		for j > left && intervals[j].Start > p {
			j--
		}
		intervals[i], intervals[j] = intervals[j], intervals[i]
	}
	intervals[i], intervals[j] = intervals[j], intervals[i]
	return i
}

func initialPivot(intervals []Interval, left, right int) {
	a := intervals[left].Start
	i := left + 1
	for ; i <= right; i++ {
		if intervals[i].Start < a {
			intervals[left], intervals[i] = intervals[i], intervals[left]
			break
		}
	}
}

func main() {
	intervals := []Interval{
		Interval{1, 2},
		Interval{2, 4},
		Interval{3, 7},
		Interval{14, 16},
	}
	fmt.Println(merge(intervals))
}
