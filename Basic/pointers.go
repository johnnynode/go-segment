package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

/*
$ go run Basic/pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0xc0420461d0

*/
