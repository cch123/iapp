package main

import (
	"strconv"

	"github.com/Knetic/govaluate"
)

func main() {
	var ops = []string{"+", "-", "*", "/", ""}
	for i := 1000; i < 10000; i++ {
		str := strconv.Itoa(i)
		for j := 0; j < len(ops); j++ {
			for k := 0; k < len(ops); k++ {
				for l := 0; l < len(ops); l++ {
					exprStr := byteToStr(str[0]) + ops[j] + byteToStr(str[1]) +
						ops[k] + byteToStr(str[2]) + ops[l] + byteToStr(str[3])

					expr, _ := govaluate.NewEvaluableExpression(string(exprStr))
					res, _ := expr.Evaluate(map[string]interface{}{})
					resF := res.(float64)
					if strconv.Itoa(int(resF)) == reverse(str) && (j != 4 || k != 4 || l != 4) {
						println(exprStr)
					}
				}
			}

		}

	}
}

func byteToStr(b byte) string {
	return string([]byte{b})
}

func reverse(input string) string {
	var res = make([]byte, 0, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		res = append(res, input[i])
	}
	return string(res)
}
