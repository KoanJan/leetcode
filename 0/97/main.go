package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	return isInterleaveBytes([]byte(s1), []byte(s2), []byte(s3))
}

func isInterleaveBytes(data1, data2, data3 []byte) bool {
	ln1, ln2, ln3 := len(data1), len(data2), len(data3)
	if ln1+ln2 != ln3 {
		return false
	}
	if ln3 == 0 {
		return true
	}
	if ln1 == 0 {
		return string(data2) == string(data3)
	} else if ln2 == 0 {
		return string(data1) == string(data3)
	}
	dp := make([][]bool, ln1+1)
	for i := range dp {
		dp[i] = make([]bool, ln2+1)
	}
	dp[0][0] = true
	for i := 1; i <= ln1; i++ {
		if data3[i-1] == data1[i-1] {
			dp[i][0] = true
		} else {
			break
		}
	}
	for j := 1; j <= ln2; j++ {
		if data3[j-1] == data2[j-1] {
			dp[0][j] = true
		} else {
			break
		}
	}
	for i := 1; i <= ln1; i++ {
		for j := 1; j <= ln2; j++ {
			if data1[i-1] == data3[i+j-1] {
				dp[i][j] = dp[i-1][j] || dp[i][j]
			}
			if data2[j-1] == data3[i+j-1] {
				dp[i][j] = dp[i][j-1] || dp[i][j]
			}
		}
	}
	return dp[ln1][ln2]
}

//func isInterleaveBytes(data1, data2, data3 []byte) bool {
//	fmt.Printf("s1: %s s2: %s s3: %s\n", string(data1), string(data2), string(data3))
//	ln1, ln2 := len(data1), len(data2)
//	i, j := 0, 0
//	for k := range data3 {
//		if (i < ln1 && data3[k] == data1[i]) && (j < ln2 && data3[k] == data2[j]) {
//			return isInterleaveBytes(data1[i+1:], data2[j:], data3[k+1:]) || isInterleaveBytes(data1[i:], data2[j+1:], data3[k+1:])
//		} else if i < ln1 && data3[k] == data1[i] {
//			i++
//		} else if j < ln2 && data3[k] == data2[j] {
//			j++
//		} else {
//			return false
//		}
//	}
//	return i == ln1 && j == ln2
//}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	//s3 := "aadbbbaccc"
	fmt.Println(isInterleave(s1, s2, s3))
}
