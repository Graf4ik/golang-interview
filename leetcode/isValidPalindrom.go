package leetcode

import "unicode"

func main() {
	isPalindrome("A man, a plan, a canal: Panama")
}

func isPalindrome(s string) bool {
	str := []rune(s)
	left := 0
	right := len(str) - 1

	for left < right {
		l := unicode.ToLower(str[left])
		r := unicode.ToLower(str[right])

		if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
			left++
			continue
		}

		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			right--
			continue
		}

		if l != r {
			return false
		}

		left++
		right--
	}
	return true
}
