package utils

import (
	"fmt"
	"strings"
)

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

func StringsToSlice(s, accordingTo string) []string {
	if len(s) == 0 {
		return []string{}
	}
	return strings.Split(s, accordingTo)
}

func FromSliceFindString(list []string, s string) bool {
	for _, str := range list {
		if str == s {
			return true
		}
	}
	return false
}
