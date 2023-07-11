package main

import "fmt"

func main() {
	var a = []int{2, 8, 5, 6, 10}
	b := a[1:4]

	/*fmt.Printf("a has value : %#v and type: %T\n", a, a)
	fmt.Printf("b has value: %v and type: %T\n", b, b)*/
	fmt.Printf("slice = %v\n", b)
	fmt.Printf("length = %v\n", len(b))
	fmt.Printf("capacity = %v\n", cap(b))

}
