package main

import "iter"

import "fmt"

type Seq[V any] func(yield func(V) bool)

type Seq2[K, V any] func(yield func(K, V) bool)

func Count(n int) iter.Seq[int] {
	return func(test func(int) bool) {
		for i := 1; i <= n; i++ {
			if !test(i) {
				return
			}
		}
	}
}

func Fibonacci(max int) iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for a <= max {
			if !yield(a) {
				return // Stop if consumer breaks
			}
			a, b = b, a+b
		}
	}
}

func main() {
	for num := range Count(5) {
		fmt.Println(num)
	}

	for num := range Fibonacci(100) {
		fmt.Println(num)
		if num > 50 {
			break
		}
	}
}
