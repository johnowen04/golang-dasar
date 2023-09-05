package golang_dasar

import "fmt"

func DeclareConstant() {
	const firstName string = "John"
	const lastName = "Owen"

	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	//firstName = "john" // error karena const tidak bisa diubah

	const (
		country  = "Indonesia"
		cc       = 62
		province = "East Java"
	)

	fmt.Println(country)
	fmt.Println(cc)
	fmt.Println(province)
}
