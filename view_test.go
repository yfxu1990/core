package main

import "net/http"

type TestView struct {
	GeneralView
}

func (v*TestView) get(r http.Request, response http.Response)  {

}

func (v*TestView) post(r http.Request, response http.Response)  {

}

func (v*TestView) pos1(r http.Request, response http.Response)  {

}
