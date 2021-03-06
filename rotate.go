// rotate.go project main.go
package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func rotateLeft(xs ...*int) {
	if len(xs) == 0 {
		return
	}
	first := *xs[0]
	for i := 1; i < len(xs); i++ {
		*xs[i-1] = *xs[i]
	}
	*xs[len(xs)-1] = first
}

func rotateRight(xs ...*int) {
	if len(xs) == 0 {
		return
	}
	last := *xs[len(xs)-1]
	for i := len(xs) - 1; i > 0; i-- {
		*xs[i] = *xs[i-1]
	}
	*xs[0] = last
}

func main() {
	x, y := 5, 10
	fmt.Print("Before Swapping: ")
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Print("After  Swapping: ")
	fmt.Println(x, y)

	a, b, c, d, e := 1, 2, 3, 4, 5
	fmt.Print("original order: ")
	fmt.Println(a, b, c, d, e)
	rotateLeft(&a, &b, &c, &d, &e)
	fmt.Print("After left rotate: ")
	fmt.Println(a, b, c, d, e)

	a, b, c, d, e = 1, 2, 3, 4, 5
	fmt.Print("original order: ")
	fmt.Println(a, b, c, d, e)
	rotateRight(&a, &b, &c, &d, &e)
	fmt.Print("After left rotate: ")
	fmt.Println(a, b, c, d, e)
}
