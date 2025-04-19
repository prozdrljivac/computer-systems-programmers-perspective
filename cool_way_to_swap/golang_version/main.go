package main

import "fmt"

func main() {
	test := []int{1, 2, 3}
	fmt.Println(test)
	reverse_array(test, len(test))
	fmt.Println(test)
}

func inplaceSwap(x *int, y *int) {
	*y = *x ^ *y
	*x = *x ^ *y
	*y = *x ^ *y
}

func reverse_array(a []int, cnt int) {
	first := 0
	last := cnt - 1

	for first < last {
		inplaceSwap(&a[first], &a[last])

		first++
		last--
	}
}
