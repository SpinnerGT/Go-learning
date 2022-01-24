package main

import (
	"fmt"
	"reflect"
)

func main3() {
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go!"))
}
