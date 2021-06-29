package main

import (
	"log"
	"reflect"
)

type Shape interface {
	area() int
	sides() int
}

type Square struct {
	Height int `json:"height" require:"true"`
	Weight int `json:"weight" require:"false"`
}

func (s *Square) area() int {
	return s.Height * s.Weight
}

func (s *Square) sides() int {
	return 2
}

func checkDataUsingReflection(v interface{}) {
	valueType := reflect.TypeOf(v)

	log.Println(valueType.Name())
	log.Println(valueType.Kind())

	totalField := valueType.NumField()
	log.Println(totalField)

	for i := 0; i < totalField; i++ {
		log.Println(valueType.Field(i))
		value := valueType.Field(i).Tag.Get("require")
		if value != "" {
			log.Println("field name: ", valueType.Field(i).Name)
			log.Println("Value of field by required: ", value)
		}
	}

}

func main() {

	square1 := Square{12, 12}
	defer func(s Shape) {
		println("area data2 is: ", s.area())
		println("sides data2 is: ", s.sides())
	}(&square1)

	square2 := new(Square)
	square2.Height = 3
	square2.Weight = 12
	defer func(s Shape) {
		println("area data is: ", s.area())
		println("sides data is: ", s.sides())
	}(square2)

	checkDataUsingReflection(square1)
}
