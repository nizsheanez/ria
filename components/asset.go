package components

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

type Asseter interface {
	Js() (string, error)
	Css() (string, error)
	Depends() []Asseter
}

type AssetConverter interface {
	Convert(sourcePath string, targetPath string, stdout chan int, stderr chan error)
	CanConvert(file string) bool
}

var converters []AssetConverter

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

func AddConverter(converter AssetConverter) {
	converters = append(converters, converter)
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

	for _, css := range this.css {

		outputFile, err := convert(css, this.baseUrl)
		if err != nil {
			return "", err
		}

		result += "<link href=\""+outputFile+"\" rel=\"stylesheet\">\n"
	}

	return result, nil
}

func convert(file string, baseUrl string) (result string, err error) {
	appRoot, err := os.Getwd()
	if err != nil {
		return
	}

	relativeSourcePath := baseUrl + "/" + file

	for _, converter := range converters {
		if converter.CanConvert(file) {
			sourcePath := appRoot + "/" + relativeSourcePath
			relativeTargetPath := "static/assets/compile/" + strings.Replace(file, ".less", ".css", 1)
			targetPath := appRoot + "/" + relativeTargetPath

			stdout := make(chan int)
			stderr := make(chan error)
			go converter.Convert(sourcePath, targetPath, stdout, stderr)
			select {
			case <-stdout:
				//just finished nothing to do
			case err := <-stderr:
				return "", err
			}

			return relativeTargetPath, nil
		}
	}

	return relativeSourcePath, nil
}

func (this *Asset) Depends() []Asseter {
	return this.depends
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

