package golang_dasar

import "fmt"

func BelajarTypeDeclaration() {
	type NoKTP string
	type Married bool

	var noKtpJohn NoKTP = "11232323232323232"
	var marriageStatus Married = false

	fmt.Println(noKtpJohn)
	fmt.Println(marriageStatus)
}
