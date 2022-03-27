package utils

import (
	"os"
	"os/exec"
)

func Shell(shell, args string, sucCb func(), failCb func(msg string)) {
	cmd := exec.Command(shell, args)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		BeforeStoppingProcess(func() {
			failCb(err.Error())
		})
	} else {
		sucCb()
	}
}
