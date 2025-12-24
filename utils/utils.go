package utils

func Check(cond bool, msg string) {
	if (!cond) {
		panic(msg)
	}
}
