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

func insert(intervals []Interval, newInterval Interval) []Interval {

	ln := len(intervals)
	newIntervals := make([]Interval, 0)

	// special
	if ln == 0 {
		return append(newIntervals, newInterval)
	}
	// 1. find start index to insert
	left1, right1 := 0, ln-1
	m1 := (left1 + right1) / 2
	for left1 < right1-1 {
		if newInterval.Start <= intervals[m1].Start {
			right1 = m1
		} else {
			left1 = m1
		}
		m1 = (left1 + right1) / 2
	}
	// fix m1
	if newInterval.Start <= intervals[m1].Start {
		m1 -= 1
	} else if m1+1 < ln && newInterval.Start > intervals[m1+1].Start {
		m1 += 1
	}

	// 2. find end index to insert
	left2, right2 := 0, ln-1
	m2 := (left2 + right2) / 2
	for left2 < right2-1 {
		if newInterval.End <= intervals[m2].End {
			right2 = m2
		} else {
			left2 = m2
		}
		m2 = (left2 + right2) / 2
	}
	// fix m2
	if newInterval.End <= intervals[m2].End {
		m2 -= 1
	} else if m2+1 < ln && newInterval.End > intervals[m2+1].End {
		m2 += 1
	}

	// 3. merge if necessary
	// a is the index of Interval with largest Start less than the new Interval
	// b is the index of Interval with smallest End more than the new Interval
	a, b := m1, m2+1
	// fix
	if a >= 0 && newInterval.Start <= intervals[a].End {
		newInterval.Start = intervals[a].Start
		a -= 1
	}
	if b < ln && newInterval.End >= intervals[b].Start {
		newInterval.End = intervals[b].End
		b += 1
	}
	for i := 0; i <= a; i++ {
		newIntervals = append(newIntervals, intervals[i])
	}
	newIntervals = append(newIntervals, newInterval)
	for i := b; i < ln; i++ {
		newIntervals = append(newIntervals, intervals[i])
	}
	return newIntervals
}

func main() {
	intervals := []Interval{
		{2, 6},
		{7, 9},
		//{10, 11},
		//{12, 13},
	}
	fmt.Println(intervals)
	newInterval := Interval{6, 10}
	fmt.Println(insert(intervals, newInterval))
}
