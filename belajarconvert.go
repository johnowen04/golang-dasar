package golang_dasar

import "fmt"

func VariableConvertion() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)
	nilai8 := int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	var name = "John Owen"
	var j = name[0]

	var jString = string(j)

	fmt.Println(name)
	fmt.Println(j)
	fmt.Println(jString)
}