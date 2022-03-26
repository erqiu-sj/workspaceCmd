package utils

import "fmt"

func JoinString(a ...any) string {
	return fmt.Sprintln(a)
}

func NotEmpty(s string) bool {
	return len(s) != 0
}
