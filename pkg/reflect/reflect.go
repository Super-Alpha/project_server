package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	Name        string `label:"Person Name: " uppercase:"true"`
	Age         int    `label:"Age is: "`
	Sex         string `label:"Sex is: "`
	Description string
}

func PrintUseTag(ptr interface{}) error {
	// 获取入参的类型
	t := reflect.TypeOf(ptr)

	// 入参类型校验
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("参数应该为结构体指针")
	}

	// 取指针指向的结构体变量
	v := reflect.ValueOf(ptr).Elem()

	// 解析字段
	for i := 0; i < v.NumField(); i++ {
		// 取tag
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag

		// 解析label tag
		label := tag.Get("label")
		if label == "" {
			label = fieldInfo.Name + ": "
		}

		// 解析uppercase tag
		value := fmt.Sprintf("%v", v.Field(i))

		if fieldInfo.Type.Kind() == reflect.String {
			uppercase := tag.Get("uppercase")
			if uppercase == "true" {
				value = strings.ToUpper(value)
			} else {
				value = strings.ToLower(value)
			}
		}

		fmt.Println(label + value)
	}

	return nil
}

type resume struct {
	Name string `json:"name" doc:"我的名字"`
}

func findDoc(str interface{}) map[string]string {
	t := reflect.TypeOf(str).Elem()
	doc := make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		doc[t.Field(i).Tag.Get("json")] = t.Field(i).Tag.Get("doc")
	}

	return doc
}

type Reflect struct {
	Name string
	Age  int
}

func (r Reflect) getType() {
	t := reflect.TypeOf(r.Name)

	fmt.Println(t)
	fmt.Println(t.Kind())
	fmt.Println(t.Name())
}

func (r Reflect) getValue() {
	v := reflect.ValueOf(r.Age)

	fmt.Println(v)
	fmt.Println(v.Kind())
}

func (r Reflect) setValue() {
	v := reflect.ValueOf(&r.Age)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	v.SetInt(20)

	fmt.Println(r.Age)
}

func (r Reflect) reflectStruct() {
	t := reflect.TypeOf(r)
	v := reflect.ValueOf(r)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("%s: %v\n", field.Name, value)
	}

	fmt.Println(t.Kind())
}

func main() {
	//person := test.Person{
	//	Name:        "Tom",
	//	Age:         29,
	//	Sex:         "Male",
	//	Description: "Cool",
	//}

	//_ = test.PrintUseTag(&person)

	r := Reflect{
		Name: "Jack",
		Age:  18,
	}

	//r.getType()
	//r.getValue()
	//r.setValue()
	r.reflectStruct()
}
