package main

import (
	"fmt"
	
	"lessons2/mylib"
	"lessons2/mylib/under"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()

	under.Hello()
}