package sqlite3

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"request/storage"
	"sync"
)

type NodeRepo struct {
	db *gorm.DB
	mu sync.Mutex
}

func NewRepo() *NodeRepo {
	repo := &NodeRepo{}
	repo.Init()
	return repo
}

func (s *NodeRepo) Init() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db != nil {
		return
	}
	db, err := gorm.Open(sqlite.Open("request-sqlite3.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&storage.Node{})
	if err != nil {
		panic(err)
	}
	s.db = db
}

func (s *NodeRepo) FirstLevelNodes() []storage.Node {
	var nodes []storage.Node
	s.db.Where("parent_id is null or parent_id = ''").Find(&nodes)
	return nodes
}

func (s *NodeRepo) QueryChildNodes(nodeId string) []storage.Node {
	var nodes []storage.Node
	s.db.Where("parent_id = ?", nodeId).Find(&nodes)
	return nodes
}

func (s *NodeRepo) HasChild(nodeId string) bool {
	var nodes []storage.Node
	s.db.Where("parent_id = ?", nodeId).Find(&nodes)
	return len(nodes) > 0
}

func (s *NodeRepo) QueryNode(nodeId string) storage.Node {
	var node storage.Node
	s.db.First(&node, nodeId)
	return node
}

func (s *NodeRepo) AddNode(node storage.Node) {
	s.db.Save(&node)
}

func (s *NodeRepo) DeleteNode(nodeId string) {
	s.db.Delete(&storage.Node{ID: nodeId})
}

func (s *NodeRepo) UpdateNode(node storage.Node) {
	s.db.Save(&node)
}
