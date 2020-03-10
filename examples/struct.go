package main

import "fmt"

func main() {
	type Position struct {
		X int
		Y int
	}

	var position1 = Position{0, 0}
	var position2 = Position{3, 2}
	fmt.Println(position1, position2)
	
	// ponteiro para struct
	var k *Position = &position2
	k.X = 33
	fmt.Println(k.X)
}