package main

import (
	"fmt"
	"reflect"
)

type Organization struct {
	Name     string `label:"name"`
	Language string `label:"language"`
}

func main() {
	organization := Organization{
		Name:     "gumi",
		Language: "Go",
	}

	fmt.Println(reflect.TypeOf(organization).Field(0).Tag)
	fmt.Println(reflect.TypeOf(organization).Field(0).Tag.Get("label"))
	fmt.Println(reflect.TypeOf(organization).Field(1).Tag)
	fmt.Println(reflect.TypeOf(organization).Field(1).Tag.Get("label"))
}
