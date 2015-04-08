/*
Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

For example,
Given input array A = [1,1,2],

Your function should return length = 2, and A is now [1,2].
*/

package main

import (
	"fmt"
)

func removeDuplicates(A []int) int {
	if len(A) == 0 {
		return 0
	}

	idx := 0
	for i := 1; i < len(A); i++ {
		if A[idx] != A[i] {
			idx = idx + 1
			A[idx] = A[i]
		}
	}

	return idx + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2, 3, 4, 4}))
}
