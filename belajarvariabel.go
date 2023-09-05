package golang_dasar

import "fmt"

func DeclareVariable() {
	var name string = "John Owen"
	fmt.Println("My name is " + name)

	var friendname = "Budi"
	fmt.Println("My new name is " + friendname)

	var age int8 = 30
	fmt.Println(age)

	newName := "John Owen"
	fmt.Println("My new name is " + newName)

	country := "Indonesia"
	fmt.Println(country)

	cc := 62
	fmt.Println(cc)

	var (
		firstName = "John"
		lastName  = "Owen"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(firstName + " " + lastName)
}
