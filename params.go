package main

import "reflect"

type SimpleUser struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Addr   string `json:"addr"`
	Extern interface{} `json:"extern"`
}

/*
*/
func ParseParams(data interface{})interface{}{
	t := reflect.TypeOf(data)
	kind := t.Kind()
	if kind != reflect.Slice && kind != reflect.String{
		return nil
	}
	rv := reflect.New(t)
	return nil
}
