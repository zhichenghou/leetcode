/*

Follow up for "Remove Duplicates":
What if duplicates are allowed at most twice?

For example,
Given sorted array A = [1,1,1,2,2,3],

Your function should return length = 5, and A is now [1,1,2,2,3].

*/

package main

import (
	"fmt"
)

func removeDuplicates2(A []int) int {
	index := 0
	for i := 0; i < len(A); i++ {
		if i > 0 && i < len(A)-1 && A[i] == A[i-1] && A[i] == A[i+1] {
			continue
		}
		A[index] = A[i]
		index = index + 1
	}

	return index
}

func main() {
	fmt.Println(removeDuplicates2([]int{1, 1, 1, 2, 2, 3}))
}
