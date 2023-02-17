package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `Required:"true" Max:"10"`
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("Required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	sample := Sample{"Tangguh"}
	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("Required"))
	fmt.Println(sampleType.Field(0).Tag.Get("Max"))

	sample.Name = ""

	fmt.Println(isValid(sample))
}
