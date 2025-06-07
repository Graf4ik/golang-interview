package leetcode

func main() {
	isValid("()")
}

func isValid(s string) bool {
	stack := []rune{}
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if isClosedBracket(char) {
			if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func isClosedBracket(symbol rune) bool {
	closingSymbols := []rune{')', '}', ']'}
	for _, s := range closingSymbols {
		if s == symbol {
			return true
		}
	}
	return false
}
