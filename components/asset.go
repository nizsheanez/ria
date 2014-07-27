package components

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
	"fmt"
	"path/filepath"
)

type Asseter interface {
	Js() (string, error)
	Css() (string, error)
	Depends() []Asseter
}

type Asset struct {
	basePath string
	baseUrl  string
	depends  []Asseter
	js       []string
	css      []string
}

func NewAsset() *Asset {
	return &Asset{}
}

func (this *Asset) SetCss(css ...string) *Asset {
	this.css = css
	return this
}

func (this *Asset) SetJs(js ...string) *Asset {
	this.js = js
	return this
}

func (this *Asset) SetBaseUrl(baseUrl string) *Asset {
	this.baseUrl = baseUrl
	return this
}

func (this *Asset) SetDependencies(dep ...Asseter) *Asset {
	this.depends = append(this.depends, dep...)
	return this
}

func (this *Asset) Js() (string, error) {
	result := ""
	for _, dep := range this.Depends() {
		content, err := dep.Js()
		if err != nil {
			return "", err
		}
		result += content
	}

	for _, js := range this.js {
		result += "<script src=\""+this.baseUrl+"/"+js+"\"></script>\n"
	}

	return result, nil
}

func (this *Asset) Css() (string, error) {
	result := ""

	for _, dep := range this.Depends() {
		content, err := dep.Css()
		if err != nil {
			return "", err
		}
		result += content
	}

	appRoot, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for _, css := range this.css {
		relativeSourcePath := this.baseUrl + "/" + css

		if strings.HasSuffix(css, ".less") {
			sourcePath := appRoot + "/" + relativeSourcePath
			relativeTargetPath := "static/assets/compile/" + strings.Replace(css, ".less", ".css", 1)
			targetPath := appRoot + "/" + relativeTargetPath

			stdout := make(chan int)
			stderr := make(chan error)
			go less(sourcePath, targetPath, stdout, stderr)
			select {
			case <-stdout:
				//just finished nothing to do
			case err := <-stderr:
				return "", err
			}

			css = relativeTargetPath
		} else {
			css = relativeSourcePath
		}

		result += "<link href=\""+css+"\" rel=\"stylesheet\">\n"
	}

	return result, nil
}

func (this *Asset) Depends() []Asseter {
	return this.depends
}

func less(sourcePath string, targetPath string, stdout chan int, stderr chan error) {

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

func execCmd(cmdStr string, chdir string) ([]byte, error) {

	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmdStr)
	head := parts[0]
	parts = parts[1:len(parts)]

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(head, parts...)
	cmd.Dir = chdir
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		return []byte{}, err
	}

	err = cmd.Wait()
	if err != nil {
		return []byte{}, errors.New(string(stderr.Bytes()))
	}

	//wg.Done() // Need to signal to waitgroup that this goroutine is done
	return stdout.Bytes(), nil
}

