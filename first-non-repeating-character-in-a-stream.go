package main

import "fmt"

func firstNonRepeatingCharacter(stream string) string {

	var (
		n      = len(stream)
		ans    = make([]byte, n)
		buffer = make([]byte, 26)
	)

	for i := 0; i < n; i++ {
		ind := stream[i] - 'a'
		buffer[ind]++

		out := byte('#')
		for j := 0; j < 26; j++ {
			if buffer[j] == 1 {
				out = byte(j) + 'a'
				break
			}
		}
		ans[i] = out
	}
	return string(ans)
}
func main() {
	stream := "aabcdbbcc"
	fmt.Println(firstNonRepeatingCharacter(stream))
}
