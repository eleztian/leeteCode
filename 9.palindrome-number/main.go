//  回文数

package __palindrome_number

import (
	"strconv"
)

func isPalindromeStr(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	n := strconv.Itoa(x)
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		if n[i] != n[j] {
			return false
		}
	}
	return true
}

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	n := x
	nx := 0
	for n > 0 {
		nx = nx*10 + n%10
		n = n / 10
	}
	return x == nx
}
