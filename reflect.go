// reflect
package main

import (
	"fmt"
	"reflect"
)

type user struct {
	id   int
	name string
	age  int
}

type manager struct {
	user
	title string
}

func main() {
	m := manager{user{1, "ok", 12}, "php"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
	fmt.Println(t.Field(0))
	fmt.Println(t.FieldByIndex([]int{1}))
	/*
		reflect.StructField{Name:"user", PkgPath:"",
		Type:(*reflect.rtype)(0x4aa2e0),
		Tag:"", Offset:0x0, Index:[]int{0},
		Anonymous:true}

		reflect.StructField{Name:"name", PkgPath:"main",
		Type:(*reflect.rtype)(0x48dc80),
		Tag:"", Offset:0x8, Index:[]int{1},
		Anonymous:false}

		{user  main.user  0 [0] true}

		{title main string  32 [1] false}
	*/
}
