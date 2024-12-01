package vo

import "github.com/qmuntal/gltf"

type ParseResponse struct {
	Document *gltf.Document
	IsParsed bool
	Err      error
}

type DocumentData struct {
	GeometryName string
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
	Name       string
	Material   *int
	Indices    *int
	Attributes map[string]string
}
