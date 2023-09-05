package golang_dasar

import "fmt"

func BelajarMathematicOps() {
	var a = 10
	var b = 5

	var c = a + b
	var d = a - c
	var e = a / d
	var f = a * e

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	var i = 10
	i += 10

	fmt.Println(i)

	i++
	fmt.Println(i)

	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)
}
