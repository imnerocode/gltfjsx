package main

import (
	"fmt"

	"github.com/imnerocode/gltfjsx/gltfjsx"
	"github.com/imnerocode/gltfjsx/vo"
)

func main() {

	pathFile := "D:/Art/Models/Basic/Ali.glb"
	parseResponse := new(vo.ParseResponse)

	parseResponse = gltfjsx.ParseGLBFromFile(pathFile)
	if parseResponse.Err != nil {
		panic(parseResponse.Err)
	}
	if !parseResponse.IsParsed {
		panic(fmt.Errorf("error trying to parse file"))
	}
	err := gltfjsx.FormatToJSX()

	if err != nil {
		panic(err)
	}

}
