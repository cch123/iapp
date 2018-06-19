package main

import (
	"strconv"
)

func main() {
	for i := 11; ; i++ {
		binary := strconv.FormatInt(int64(i), 2)
		oct := strconv.FormatInt(int64(i), 8)
		base := strconv.Itoa(i)
		if base == reverse(base) &&
			binary == reverse(binary) &&
			oct == reverse(oct) {
			println(i, binary, oct)
			break
		}
	}
}

func reverse(input string) string {
	var res = make([]byte, 0, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		res = append(res, input[i])
	}
	return string(res)
}
