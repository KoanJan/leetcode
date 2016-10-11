package main

import (
	"fmt"
)

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	res := make([]float64, len(queries))

	// not existed
	if len(equations) == 0 {
		for i := 0; i < len(res); i++ {
			res[i] = -1.0
		}
		return res
	}

	// init
	vals := map[string]map[string]float64{}

	// store
	for i := 0; i < len(equations); i++ {
		var (
			a string = equations[i][0]
			b string = equations[i][1]
		)
		// set a
		if _, existA := vals[a]; existA {
			vals[a][b] = values[i]
		} else {
			vals[a] = map[string]float64{b: values[i]}
		}
		// set b
		if _, existB := vals[b]; existB {
			vals[b][a] = 1.0 / values[i]
		} else {
			vals[b] = map[string]float64{a: 1.0 / values[i]}
		}
	}

	// calc
	for i := 0; i < len(queries); i++ {
		res[i], _ = dst(queries[i][0], queries[i][1], vals, map[string]bool{})
	}

	return res
}

func dst(a, b string, src map[string]map[string]float64, hasPassedBy map[string]bool) (float64, bool) {
	_, existA := src[a]
	_, existB := src[b]
	if existA && existB {
		// recursive
		if d, existDirected := src[a][b]; existDirected {
			return d, true
		}
		for k, v := range src[a] {
			if !hasPassedBy[k] {
				hasPassedBy[k] = true
				if d, yes := dst(k, b, src, hasPassedBy); yes {
					return d * v, true
				}
			}
		}
	}
	return -1.0, false
}

func formatFloat64Slice(input []float64) string {
	out := "%.1f"
	tmp := make([]interface{}, len(input))
	for i := 0; i < len(input)-1; i++ {
		out += " %.1f"
		tmp[i] = input[i]
	}
	tmp[len(input)-1] = input[len(input)-1]
	return fmt.Sprintf(out, tmp...)
}

func main() {

	// test
	var (
		equations = [][]string{
			[]string{"x1", "x2"},
			[]string{"x2", "x3"},
			[]string{"x3", "x4"},
			[]string{"x4", "x5"},
		}
		values  = []float64{3.0, 4.0, 5.0, 6.0}
		queries = [][]string{
			[]string{"x1", "x5"},
			[]string{"x5", "x2"},
			[]string{"x2", "x4"},
			[]string{"x2", "x2"},
			[]string{"x2", "x9"},
			[]string{"x9", "x9"},
		}
	)
	fmt.Printf("equations: %v\nvalues: %v\nqueries: %v\ncalcEquation got %v\n", equations, formatFloat64Slice(values), queries, formatFloat64Slice(calcEquation(equations, values, queries)))
}
