package leetcode

func main() {
	isSubsequence("abc", "ahbgdc")
}

func isSubsequence(s string, t string) bool {
	i := 0
	j := 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] { // Если текущие символы совпадают (s[i] == t[j]) — сдвигаем оба указателя (нашли нужную букву).
			i++
			j++
		} else {
			j++ // Иначе — двигаем только j (ищем следующую возможную позицию для s[i] в t).
		}
	}

	return i == len(s) // Если прошли всю строку s (i == len(s)), значит, все символы s нашли в t по порядку → ✅ true.
}
