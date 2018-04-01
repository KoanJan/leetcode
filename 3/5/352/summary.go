package main

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

type SummaryRanges struct {
	intervals []Interval
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{intervals: []Interval{}}
}

func (this *SummaryRanges) Addnum(val int) {
	var (
		l, r int = 0, len(this.intervals) - 1
		mid  int = (l + r) / 2
		loc  int = -1
	)
	if r == -1 {
		this.intervals = append(this.intervals, Interval{Start: val, End: val})
		return
	}
	for r-l > 1 {
		interval := this.intervals[mid]
		if interval.Start <= val && interval.End >= val {
			return
		}
		if interval.Start > val {
			r = mid
		} else {
			l = mid
		}
		mid = (l + r) / 2
	}
	if this.intervals[l].Start > val {
		loc = l
	} else if this.intervals[l].End >= val {
		return
	} else if this.intervals[r].Start > val {
		loc = r
	} else if this.intervals[r].End >= val {
		return
	} else {
		loc = r + 1
	}
	tmp := []Interval{}
	if loc == 0 {
		if val == this.intervals[0].Start-1 {
			this.intervals[0].Start = val
		} else {
			tmp = append(tmp, Interval{Start: val, End: val})
			for i := 0; i < len(this.intervals); i++ {
				tmp = append(tmp, this.intervals[i])
			}
			this.intervals = tmp
		}
	} else if loc == len(this.intervals) {
		if val == this.intervals[loc-1].End+1 {
			this.intervals[loc-1].End = val
		} else {
			this.intervals = append(this.intervals, Interval{Start: val, End: val})
		}
	} else {
		if val == this.intervals[loc].Start-1 {
			this.intervals[loc].Start = val
		}
		if val == this.intervals[loc-1].End+1 {
			this.intervals[loc-1].End = val
		}
		for i := 0; i < loc; i++ {
			tmp = append(tmp, this.intervals[i])
		}
		if tmp[loc-1].End != val {
			tmp = append(tmp, Interval{Start: val, End: val})
		}
		if tmp[len(tmp)-1].End == this.intervals[loc].Start {
			tmp[len(tmp)-1].End = this.intervals[loc].End
		} else {
			tmp = append(tmp, this.intervals[loc])
		}
		for i := loc + 1; i < len(this.intervals); i++ {
			tmp = append(tmp, this.intervals[i])
		}
		this.intervals = tmp
	}
}

func (this *SummaryRanges) Getintervals() []Interval {
	return this.intervals
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Addnum(val);
 * param_2 := obj.Getintervals();
 */
