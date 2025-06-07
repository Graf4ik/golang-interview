package leetcode

func main() {
	isAnagram("anagram", "nagaram")
}

func isAnagram(s string, t string) bool {
	chars := make(map[rune]int, len(s))

	for _, v := range s {
		chars[v]++
	}

	for _, v := range t {
		chars[v]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}
