package utils

import "fmt"

func JoinString(a ...any) string {
	return fmt.Sprintln(a)
}

func NotEmpty(s string) bool {
	return len(s) != 0
}

func StringDefault(s, d string) string {
	if len(s) == 0 {
		return d
	}
	return s
}
