package baoFmt

import (
	"fmt"
	"reflect"
)

func PrintLn(v any) {
	Type := reflect.TypeOf(v)
	Value := reflect.ValueOf(v)
	var res []string
	res = append(res, "{")
	if Type.Kind() == reflect.Struct {
		for i := 0; i < Type.NumField(); i++ {
			field := Type.Field(i)
			if i != Type.NumField()-1 {
				res = append(res, printfJSON(field.Name, Value.FieldByName(field.Name), 1))
			} else {
				res = append(res, printfJSON(field.Name, Value.FieldByName(field.Name), 0))
			}
		}
	} else {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("[", Type.Kind(), "]", v)
			}
		}()
		Type = Type.Elem()
		if Type.Kind() == reflect.Struct {
			Value = Value.Elem()
			for i := 0; i < Type.NumField(); i++ {
				field := Type.Field(i)
				if i != Type.NumField()-1 {
					res = append(res, printfJSON(field.Name, Value.FieldByName(field.Name), 1))
				} else {
					res = append(res, printfJSON(field.Name, Value.FieldByName(field.Name), 0))
				}
			}
		}
	}
	res = append(res, "}")
	for _, v := range res {
		fmt.Println(v)
	}
}

func printfJSON(key string, value reflect.Value, needComma int) string {
	jsonField := ""
	if value.Kind() != reflect.String {
		jsonField = fmt.Sprintf("\"%s\":%v", key, value)
	} else {
		jsonField = fmt.Sprintf("\"%s\":\"%v\"", key, value)
	}
	if needComma == 1 {
		return jsonField + ","
	} else {
		return jsonField
	}
}
