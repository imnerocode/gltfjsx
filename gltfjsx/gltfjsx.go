package gltfjsx

import (
	"os"
	"text/template"

	"github.com/imnerocode/gltfjsx/constants"
	"github.com/imnerocode/gltfjsx/helpers"
	"github.com/imnerocode/gltfjsx/templates"
	"github.com/imnerocode/gltfjsx/vo"
	"github.com/qmuntal/gltf"
)

func ParseGLBFromFile(pathFile string) *vo.ParseResponse {

	doc, err := gltf.Open(pathFile)
	if err != nil {
		return &vo.ParseResponse{Document: nil, IsParsed: false, Err: err}
	}

	return &vo.ParseResponse{Document: doc, IsParsed: true, Err: nil}
}

func FormatToJSX() (string, error) {
	doc := ParseGLBFromFile(constants.PATH_FILE)
	if doc.Err != nil {
		return "", doc.Err
	}

	if !doc.IsParsed {
		return "", helpers.ErrParse
	}

	templateJsx := templates.TemplateJSX()

	tmpl, err := template.New("test").Parse(templateJsx)

	if err != nil {
		return "", err
	}
	var docData vo.DocumentData

	for _, node := range doc.Document.Nodes {
		var nodeData vo.NodeData
		nodeData.Rotation = node.Rotation
		nodeData.Scale = node.Scale
		nodeData.Position = node.Translation

		if node.Mesh != nil {
			nodeData.MeshID = node.Mesh
		}

		docData.Nodes = append(docData.Nodes, nodeData)
	}

	for _, mesh := range doc.Document.Meshes {
		var meshData vo.MeshData
		meshData.Name = mesh.Name
		meshData.Material = mesh.Primitives[0].Material
		meshData.Indices = mesh.Primitives[0].Indices

		// Procesar atributos
		attributes := make(map[string]string)
		for key := range mesh.Primitives[0].Attributes {
			attributes[key] = "buffer" // Aquí puedes mapear a valores reales
		}
		meshData.Attributes = attributes

		docData.Meshes = append(docData.Meshes, meshData)
	}
	err = tmpl.Execute(os.Stdout, docData)
	if err != nil {
		return "", err
	}
	return "", nil
}
