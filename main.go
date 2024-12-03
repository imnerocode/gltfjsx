package main

import (
	"github.com/imnerocode/gltfjsx/gltfjsx"
)

func main() {

	err := gltfjsx.FormatToJSX()

	if err != nil {
		panic(err)
	}

}
