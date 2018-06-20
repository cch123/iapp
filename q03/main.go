package main

func main() {
	var cards = [101]bool{}
	for i := 2; i <= 100; i++ {
		for j := i; j <= 100; j += i {
			cards[j] = !cards[j]
		}
	}

	for i := 1; i < 101; i++ {
		if !cards[i] {
			println(i)
		}
	}
}
