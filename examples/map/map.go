package map

import "fmt"

func main() {
	var notas map[string]int = make(map[string]int)

	notas["nico"] = 10
	notas["leo"] = 9

	fmt.Println(notas)
	
	
	// map literal ---------
	
	notas2 := map[string]int{
		"nico": 7,
		"leo": 6,
	}


	fmt.Println(notas2)
}