package test

import (
	"testing"
	"log"
	"runtime"
	"io/ioutil"
	"path/filepath"
	_ "ria/routers"
//	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	httplib "github.com/astaxie/beego/testing"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}


// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	request := httplib.Get("/user/get")
	request.Param("id", "1")
	response, _ := request.Response()
//	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	log.Fatalln(fmt.Sprintf("%v", string(contents)))
//	var s ShortResult
//	json.Unmarshal(contents, &s)
//	if s.UrlLong == "" {
//		t.Fatal("urllong is empty")
//	}

}

