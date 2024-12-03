package vo

import "github.com/qmuntal/gltf"

type ParseResponse struct {
	Document *gltf.Document
	IsParsed bool
	Err      error
}

type DocumentData struct {
	GeometryName []string
	Meshes       []MeshData
}

type MeshData struct {
	Material   *int
	Indices    any
	Attributes AttributesData
	Node       []NodeData
}
type NodeData struct {
	Rotation [4]float64
	Scale    [3]float64
	Position [3]float64
	MeshID   *int
}

type AttributesData struct {
	Normal   [][3]float32
	Position [][3]float32
	TexCoord [][2]float32
	Indices  any
}
type AttributesMain struct {
	Normal   []float32
	Position []float32
	TexCoord [][2]float32
	Indices  any
}
