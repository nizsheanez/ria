package components

import (
	"errors"
	"os"
	"fmt"
	"path/filepath"
	"strings"
)

type Less struct {}

func (this *Less) CanConvert(file string) bool {
	return strings.HasSuffix(file, ".less")
}

func (this *Less) Convert(sourcePath string, targetPath string, stdout chan int, stderr chan error) {

	fi1, err := os.Stat(sourcePath)
	if err != nil {
		stderr <- err
		return
	}
	mtime1 := fi1.ModTime().Unix()
	fi2, err := os.Stat(targetPath)
	var mtime2 int64
	if err != nil {
		if !os.IsNotExist(err) {
			stderr <- err
			return
		}
	} else {
		mtime2 = fi2.ModTime().Unix()
	}

	if mtime2 < mtime1 {
		cmd := "lessc %s %s --no-color"
		cmd = fmt.Sprintf(cmd, sourcePath, targetPath)
		_, err := execCmd(cmd, filepath.Dir(sourcePath))
		if err != nil {
			stderr <- errors.New(fmt.Sprintf("Command failed: %s\n%s", cmd, err.Error()))
			return
		}
	}

	stdout <- 1
}
