package iteration

const repeatCount = 5

func Repeat(char string) (result string) {
	result = ""
	for i := 0; i < repeatCount; i++ {
		result += char
	}
	return
}
