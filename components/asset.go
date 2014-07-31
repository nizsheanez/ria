package components

import (
	"bytes"
	"errors"
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
		result += "<script src=\""+this.buildUrl(js)+"\"></script>\n"
	}

	return result, nil
}

func (this *Asset) buildUrl(relativeUrl string) (url string) {
	if strings.HasPrefix(relativeUrl, "http:") ||
			strings.HasPrefix(relativeUrl, "https:") ||
			strings.HasPrefix(relativeUrl, "//") {
		url = relativeUrl
	} else {
		url = this.baseUrl+"/"+relativeUrl
	}

	return
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
		result += "<link href=\""+this.buildUrl(css)+"\" rel=\"stylesheet\">\n"
	}

	return result, nil
}

func convert(file string, baseUrl string) (result string, err error) {
	if err != nil {
		return
	}

	return baseUrl+"/"+file, nil
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

