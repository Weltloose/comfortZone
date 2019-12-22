package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// 传入结构体指针与请求将请求信息进行解析
// 支持转换类型为int, bool, string
// 用户需在结构体的tag中设置from字段标明从哪里获取,Header, Body, Rest, Form, PostForm
// 可以增加name标签选择参数在请求中的键，默认是结构体名称的小写
func Unpack(req *http.Request, ptr interface{}) error {

	// 构建由有效名称键控的字段的映射。
	var err error
	v := reflect.ValueOf(ptr).Elem() // 结构变量
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // 反射.StructField
		tag := fieldInfo.Tag           // 一个 reflect.StructTag
		name := tag.Get("name")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		err = parseAccordingTag(req, tag, v.Field(i), name)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseAccordingTag(r *http.Request, tag reflect.StructTag, v reflect.Value, name string) error {
	var value string
	var err error
	switch tag.Get("from") {
	case "Header":
		value = r.Header.Get(name)
	case "Body":
		var tmp []byte
		tmp, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		value = string(tmp)
	case "Rest":
		value = mux.Vars(r)[name]
	case "Form":
		value = r.FormValue(name)
	case "PostForm":
		value = r.PostFormValue(name)
	default:
		return fmt.Errorf("Unknown type tag %s", tag.Get("from"))
	}
	return populate(v, value)
}

// 将字符串转为对应类型，如果转换失败，返回对应默认值
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			v.SetInt(0)
		} else {
			v.SetInt(i)
		}
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			v.SetBool(false)
		} else {
			v.SetBool(b)
		}
	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
