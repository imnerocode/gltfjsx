package vo

import "github.com/qmuntal/gltf"

type ParseResponse struct {
	Document *gltf.Document
	IsParsed bool
	Err      error
}

type DocumentData struct {
	GeometryName []string
	Nodes        []NodeData
	Meshes       []MeshData
}

type NodeData struct {
	Rotation [4]float64
	Scale    [3]float64
	Position [3]float64
	MeshID   *int
}

type MeshData struct {
	Material   *int
	Indices    any
	Attributes []AttributesData
}

type AttributesData struct {
	Normal   [][3]float32
	Position [][3]float32
	TexCoord [][2]float32
}
