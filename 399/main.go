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
	vals := map[string]float64{equations[0][1]: 1.0}

	// scan 2 times and got all valus
	for i := 0; i < len(equations); i++ {
		a, existA := vals[equations[i][0]]
		b, existB := vals[equations[i][1]]
		if (existA && existB) || (!existA && !existB) {
			continue
		}
		if existA {
			vals[equations[i][1]] = a / values[i]
		} else {
			vals[equations[i][0]] = b * values[i]
		}
	}
	for i := 0; i < len(equations); i++ {
		a, existA := vals[equations[i][0]]
		b, existB := vals[equations[i][1]]
		if existA && existB {
			continue
		}
		if existA {
			vals[equations[i][1]] = a / values[i]
		} else if existB {
			vals[equations[i][0]] = b * values[i]
		} else {
			// new numbers have no relations with numbers found in time1
			vals[equations[i][1]] = 1.0
			vals[equations[i][0]] = values[i] * 1.0
		}
	}

	// calc
	for i := 0; i < len(queries); i++ {
		a, existA := vals[queries[i][0]]
		b, existB := vals[queries[i][1]]
		if existA && existB {
			res[i] = a / b
		} else {
			// not existed
			res[i] = -1.0
		}
	}

	return res
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
		equations = [][]string{[]string{"a", "b"}, []string{"b", "c"}}
		values    = []float64{2.0, 3.0}
		queries   = [][]string{[]string{"a", "c"}, []string{"b", "a"}, []string{"a", "e"}, []string{"a", "a"}, []string{"x", "x"}}
	)
	fmt.Printf("equations: %v\nvalues: %v\nqueries: %v\ncalcEquation got %v\n", equations, formatFloat64Slice(values), queries, formatFloat64Slice(calcEquation(equations, values, queries)))
}
