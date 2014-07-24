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
	Js() string
	Css() string
}

type Asset struct {
	basePath string
	baseUrl  string
	depends  []*Asset
	js       []string
	css      []string
}

func (asset *Asset) Js() (string, error) {
	result := ""
	for _, dep := range asset.depends {
		content, err := dep.Js()
		if err != nil {
			return "", err
		}
		result += content
	}

	for _, js := range asset.js {
		result += "<script src=\""+asset.baseUrl+"/"+js+"\"></script>\n"
	}

	return result, nil
}

func (asset *Asset) Css() (string, error) {
	result := ""

	for _, dep := range asset.depends {
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

	for _, css := range asset.css {
		relativeSourcePath := asset.baseUrl + "/" + css

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

func Assets() *Asset {
	Jquery := Asset{
		baseUrl: "static/vendor",
		js: []string{
			"jquery/dist/jquery.min.js",
		},
	}

	Angular := Asset{
		baseUrl: "static/vendor",
		js: []string{
			"angular/angular.js",
			"angular-resource/angular-resource.js",
			"angular-route/angular-route.min.js",
			// "angular-ui-router/release/angular-ui-router.min.js",
			"angular-translate/angular-translate.js",
		},
	}

	AngularUi := Asset{
		baseUrl: "static/vendor",
		css: []string{
			"angular-ui/build/angular-ui.min.css",
		},
		js: []string{
			"angular-ui/build/angular-ui.min.js",
			"angular-ui-sortable/src/sortable.js",
		},
	}

	UiBootstrap := Asset{
		baseUrl: "static/vendor/angular-bootstrap",
		js: []string{
			"ui-bootstrap.min.js",
			"ui-bootstrap-tpls.min.js",
		},
	}

	AngularElastic := Asset{
		baseUrl: "static/vendor/angular-elastic",
		js: []string{
			"elastic.js",
		},
	}

	AngularUiUtils := Asset{
		baseUrl: "static/vendor/angular-ui-utils",
		js: []string{
			"ui-utils.min.js",
		},
	}

	App := Asset{
		baseUrl: "static",
		css: []string{
			"less/site.less",
		},
		js: []string{
			"common/fixes.js",
			"common/debug.js",
			"common/components.js",
			"app/app.js",
			"app/goal/services/tpl.js",
			"app/goal/services/modal.js",
			"app/goal/services/user.js",
			"app/goal/services/category.js",
			"app/goal/services/report.js",
			"app/goal/services/goal.js",
			"app/goal/controllers/goal.js",
			"app/goal/controllers/nav.js",
			"app/goal/directives/editor.js",
			"app/goal/services/alert.js",
		},
		depends: []*Asset{
			&Jquery,
			&Angular,
			&AngularUi,
			&UiBootstrap,
			&AngularElastic,
			&AngularUiUtils,
			textArea(),
		},
	}

	return &App
}

func textArea() *Asset {
	AngularSanitize := Asset{
		baseUrl: "static/vendor/angular-sanitize",
		js: []string{
			"angular-sanitize.js",
		},
	}

	FontAwesome := Asset{
		baseUrl: "static/vendor/components-font-awesome",
		css: []string{
			"css/font-awesome.min.css",
		},
	}

	TextAngular := Asset{
		baseUrl: "static/vendor/textAngular",
		js: []string{
			"textAngular.js",
		},
		depends: []*Asset{
			&AngularSanitize,
			&FontAwesome,
		},
	}
	return &TextAngular
}
