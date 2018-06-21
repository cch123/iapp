package main

import "fmt"

func main() {
	N := 1000
	// 10, 50, 100, 500
	var valArr = []int{10, 50, 100, 500} // const
	var curCntArr = []int{0, 0, 0, 0}
	var maxCntArr = []int{15, 15, 10, 2}
	var counter int
	for {
		if calcTotal(valArr, curCntArr) == N && totalCoin(curCntArr) <= 15 { //&& totalCoin(curCntArr) < 15 {
			fmt.Println(curCntArr)
			counter++
		}

		if !push(N, curCntArr, valArr, maxCntArr) {
			break
		}
	}
	fmt.Println(counter)
}

func push(n int, curCntArr []int, valArr []int, maxCntArr []int) bool {
	for idx := 0; idx < len(curCntArr); idx++ {
		if curCntArr[idx] < maxCntArr[idx] {
			curCntArr[idx]++
			for i := 0; i < idx; i++ {
				curCntArr[i] = 0
			}
			return true
		}
	}
	return false
}

func calcTotal(valArr []int, curCntArr []int) int {
	cnt := 0
	for i := 0; i < len(valArr); i++ {
		cnt += curCntArr[i] * valArr[i]
	}
	return cnt
}

func totalCoin(curCntArr []int) int {
	cnt := 0
	for i := 0; i < len(curCntArr); i++ {
		cnt += curCntArr[i]
	}
	return cnt
}
