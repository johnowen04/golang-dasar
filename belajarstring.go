package golang_dasar

func SayName() string {
	return "John"
}

func SayFullName() string {
	return "John Owen"
}

func GetLength(name string) int {
	return len(name)
}

func GetFirstChar(name string) byte {
	return name[0]
}

func GetLastChar(name string) byte {
	return name[GetLength(name)-1]
}
