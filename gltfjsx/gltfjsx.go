package gltfjsx

import (
	"fmt"

	"github.com/imnerocode/gltfjsx/constants"
	"github.com/imnerocode/gltfjsx/helpers"
	"github.com/imnerocode/gltfjsx/vo"
	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

func ParseGLBFromFile(pathFile string) *vo.ParseResponse {

	doc, err := gltf.Open(pathFile)
	if err != nil {
		return &vo.ParseResponse{Document: nil, IsParsed: false, Err: err}
	}

	return &vo.ParseResponse{Document: doc, IsParsed: true, Err: nil}
}

func FormatToJSX() error {
	pathFile := constants.PATH_FILE
	parseData := ParseGLBFromFile(pathFile)
	if parseData.Err != nil {
		return parseData.Err
	}
	if !parseData.IsParsed {
		return helpers.ErrParse
	}

	doc := parseData.Document
	var documentData vo.DocumentData

	for _, mesh := range doc.Meshes {
		var meshData vo.MeshData

		// Adding the name of the geometry.
		documentData.GeometryName = append(documentData.GeometryName, mesh.Name)
		for _, primitive := range mesh.Primitives {
			if primitive.Material != nil {
				meshData.Material = primitive.Material
			}
			if primitive.Indices != nil {
				var buff []byte
				indices, err := modeler.ReadAccessor(doc, doc.Accessors[*primitive.Indices], buff)
				if err != nil {
					return err
				}
				meshData.Indices = indices
			}

			var attributesData vo.AttributesData
			for key, value := range primitive.Attributes {
				switch key {
				case "POSITION":
					var buff [][3]float32
					position, err := modeler.ReadPosition(doc, doc.Accessors[value], buff)
					if err != nil {
						return err
					}
					attributesData.Position = position
				case "NORMAL":
					var buff [][3]float32
					normal, err := modeler.ReadNormal(doc, doc.Accessors[value], buff)
					if err != nil {
						return err
					}
					attributesData.Normal = normal
				case "TEXCOORD_0":
					var buff [][2]float32
					texCoord, err := modeler.ReadTextureCoord(doc, doc.Accessors[value], buff)
					if err != nil {
						return err
					}
					attributesData.TexCoord = texCoord
				}

			}
			meshData.Attributes = append(meshData.Attributes, attributesData)
		}
		documentData.Meshes = append(documentData.Meshes, meshData)
	}

	fmt.Printf("Mesh name: %s\n", documentData.GeometryName[0])
	fmt.Printf("Mesh Indices: %+v\n", documentData.Meshes[0].Indices)
	fmt.Printf("Mesh Material: %+v\n", documentData.Meshes[0].Material)
	fmt.Printf("Mesh Attributes: %+v\n", documentData.Meshes[0].Attributes)
	return nil
}
