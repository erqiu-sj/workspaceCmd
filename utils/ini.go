package utils

import (
	"fmt"
	"path/filepath"
)

func AddNewWorkingGroup(groupName string) {
	CreateFile(filepath.Join("./ini/", fmt.Sprint(groupName, ".ini")))
}
