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

	fmt.Printf("Name Geometry: %s\n", parseResponse.Document.Nodes[0].Name)
	fmt.Printf("Attributes Geometry: %v\n", parseResponse.Document.Meshes[0].Primitives[0].Attributes)
	fmt.Printf("Extensions Geometry: %v\n", parseResponse.Document.Meshes[0].Primitives[0].Extensions)
	fmt.Printf("Rotation Geometry: %v\n", parseResponse.Document.Nodes[0].Rotation)
	fmt.Printf("Mesh Data: %v\n", parseResponse.Document.Meshes[0].Primitives)
	fmt.Printf("Scale Geometry: %v\n", parseResponse.Document.Nodes[0].Scale)

	_, err := gltfjsx.FormatToJSX()

	if err != nil {
		panic(err)
	}
}
