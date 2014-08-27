package controllers

import (
	"github.com/astaxie/beego"
	"ria/modules/v1/article/models"
	"errors"
	"net/url"
	"fmt"
	"reflect"
	"strings"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	var id int
	this.Ctx.Input.Bind(&id, "id")

	if id <= 0 {
		this.Data["json"] = errors.New("User id is required")
		this.ServeJson()
		return
	}

	article, err := models.NewArticle().FindById(id)
	if err != nil {
		beego.Error(err)
	}

	this.Data["json"] = article

	this.ServeJson()
}

func (this *ArticleController) Post() {
	article := models.NewArticle()
	err := this.ParseForm(article)
	if err != nil {
		this.Data["json"] = err
	} else {
		ok, err := article.Validate("create");
		if ok {
			this.Data["json"] = article
		} else {
			this.Data["json"] = err
		}
	}
	this.ServeJson()
}

func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func ParseForm(form url.Values, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	if !isStructPtr(objT) {
		return fmt.Errorf("%v must be  a struct pointer", obj)
	}
	objT = objT.Elem()
	objV = objV.Elem()

	for i := 0; i < objT.NumField(); i++ {
		fieldV := objV.Field(i)
		if !fieldV.CanSet() {
			continue
		}

		fieldT := objT.Field(i)
		tags := strings.Split(fieldT.Tag.Get("form"), ",")

		var tag string
		if len(tags) == 0 || len(tags[0]) == 0 {
			tag = fieldT.Name
		} else if tags[0] == "-" {
			continue
		} else {
			tag = tags[0]
		}

		value := form.Get(tag)
		if len(value) == 0 {
			continue
		}

		switch fieldT.Type.Kind() {
		case reflect.Bool:
			b, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			fieldV.SetBool(b)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			fieldV.SetInt(x)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			x, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}
			fieldV.SetUint(x)
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			fieldV.SetFloat(x)
		case reflect.Interface:
			fieldV.Set(reflect.ValueOf(value))
			panic(fieldV)
		case reflect.String:
			fieldV.SetString(value)
		}
	}
	return nil
}

func (this *ArticleController) List() {
	this.ServeJson()
}
