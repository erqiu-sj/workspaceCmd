package utils

import "fmt"

func GreenTips(text string) {
	fmt.Printf("\x1b[%dm"+fmt.Sprint(text)+"  \x1b[0m\n", 32)
}

func FillGrayTips(text interface{}) {
	fmt.Printf("\x1b[%d;%dmtime consuming: \x1b[0m "+fmt.Sprint(text)+" \n", 47, 30)
}

func RedTips(text interface{}) {
	fmt.Printf("\x1b[%dm"+fmt.Sprint(text)+" \x1b[0m\n", 31)
}
