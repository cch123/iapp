package main

import "fmt"

func main() {
	N := 1000
	// 10, 50, 100, 500
	var valArr = []int{10, 50, 100, 500} // const
	var curCntArr = []int{0, 0, 0, 0}
	for !canStop(N, curCntArr, valArr) {
		push(N, curCntArr, valArr)
		if calcTotal(valArr, curCntArr) == N {
			fmt.Println(curCntArr)
		}
	}
}

func push(n int, curCntArr []int, valArr []int) {
	if curCntArr[0]*valArr[0] < n {
		curCntArr[0]++
		return
	}

	curCntArr[0] = 0
	for i := 1; i < len(curCntArr); i++ {
		if curCntArr[i]*valArr[i] < n {
			curCntArr[i]++
			return
		}
	}
}

func calcTotal(valArr []int, curCntArr []int) int {
	cnt := 0
	for i := 0; i < len(valArr); i++ {
		cnt += curCntArr[i] * valArr[i]
	}
	return cnt
}

func canStop(n int, curCntArr []int, valArr []int) bool {
	var canStop = true
	for i := 0; i < len(curCntArr); i++ {
		if valArr[i]*curCntArr[i] < n {
			canStop = false
			break
		}
	}
	return canStop
}
