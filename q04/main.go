package main

func main() {
	m, n := 100, 5
	counter := 0
	left := m
	for left > 2*n {
		left -= n
		counter++
	}

	for left > 1 {
		if left%2 == 0 {
			left = left / 2
			counter++
		} else {
			left = (left + 1) / 2
			counter++
		}
	}
	println(m, n, counter)
}
