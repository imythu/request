package storage

type NodeRepo interface {
	Init()
	FirstLevelNodes() []Node
	QueryChildNodes(nodeId string) []Node
	HasChild(nodeId string) bool
	QueryNode(nodeId string) Node
	AddNode(node Node)
	DeleteNode(nodeId string)
	UpdateNode(node Node)
}
