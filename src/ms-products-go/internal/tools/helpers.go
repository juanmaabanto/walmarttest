package tools

func EsPalindrome(valor string) bool {
	var igual, aux int

	for char := len(valor) - 1; char >= 0; char-- {
		if valor[char] == valor[aux] {
			igual++
		}
		aux++
	}

	return len(valor) == igual
}
