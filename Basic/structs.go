package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20}) // {Bob 20}
	fmt.Println(person{name:"Alice", age: 30}) // {Alice 30}

	fmt.Println(person{name:"Fred"}) // {Fred 0}
	fmt.Println(&person{name:"Ann", age: 40}) // {Ann 40}
	s := person{name:"Sean", age:50}
	fmt.Println(s.name) // Sean

	sp := &s
	fmt.Println(sp.age) // 50

	sp.age = 51
	fmt.Println(sp.age) // 51
}

/*

$ go run Basic/structs.go
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
Sean
50
51

*/