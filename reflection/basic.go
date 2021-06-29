package main

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

type Student struct {
	Name   string `json:"name" required:"true"`
	Number int    `json:"number" required:"true" max:"8"`
}

func main() {

	std := Student{
		Name:   "J I Rakib",
		Number: 80023583,
	}

	validation := checkValidation(std)
	if len(validation) > 0 {
		for val := range validation {
			log.Println(val)
		}
	}

}

func checkValidation(std interface{}) []string {
	var exception string
	obj := reflect.TypeOf(std)

	if obj.Kind() == reflect.Struct {
		for i := 0; i < obj.NumField(); i++ {

			name := obj.Field(i).Name
			fieldValue := reflect.ValueOf(std).Field(i)

			required := obj.Field(i).Tag.Get("required")
			if required != "false" {
				if fieldValue.String() == "" {
					exception += name + " required value can not be null."
				}
			}

			max := obj.Field(i).Tag.Get("max")
			if max != "" {
				data, err := strconv.Atoi(max)
				if err != nil {
					exception += name + " data parsing error in max value."
				}
				preString := strconv.Itoa(int(fieldValue.Int()))
				length := len(preString)

				if data != length {
					exception += name + " value do not meet requirement."
				}
			}

		}
	}
	err := strings.Split(exception, ".")
	if len(err) <= 0 {
		return make([]string, 0)
	}
	log.Println(err)
	return err
}
