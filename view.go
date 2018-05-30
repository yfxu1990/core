package main

import (
	"net/http"
	"fmt"
	"reflect"
	"delaycmd/common"
)

type CallBack func (data interface{}, req http.Request)(*http.Response,int)

type BaseMember struct {
	Method string `json:"method"`
	Handle CallBack
}

type BaseView struct {
	router string
	SerializeClass interface{}
	/*method*/
	MethodSet map[string]BaseMember
}

func (v *BaseView) Handler(method string,req http.Request) (*http.Response,int) {
	value,ok := v.MethodSet[method];if !ok{
		err := fmt.Sprint(method,"is not exist!")
		panic(err)
		return nil, -1
	}
	return value.Handle(v.SerializeClass,req)
}

type GeneralView struct {
	BaseView
}

type FrameWork struct {
	Views map[string]*GeneralView
}

func _createView(url string, t reflect.Type)  {

}

