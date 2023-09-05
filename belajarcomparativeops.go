package golang_dasar

import "fmt"

func BelajarComparativeOps() {
	var name1 = "John"
	var name2 = "John"

	var result = name1 == name2
	fmt.Println(result)

	name1 = "Budi"
	var result2 = name1 == name2
	fmt.Println(result2)

	a := 100
	b := 200
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	fmt.Println(a == b)
	fmt.Println(a != b)
}
