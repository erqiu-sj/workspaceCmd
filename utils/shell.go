package utils

import (
	"os"
	"os/exec"
)

func Shell(shell string, sucCb func(), failCb func(msg string)) {
	cmd := exec.Command("/bin/sh", "-c", shell)
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
