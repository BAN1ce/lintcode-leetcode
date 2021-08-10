package string


func reverseLeftWords(s string, n int) string {

	if n==0 {
		return s
	}

	left := s[0:n]

	right := s[n:]

	return right+left
}