package storage

import (
	"database/sql"
	"time"
)

type Node struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
	Name      string       `gorm:"name"`
	ParentId  string       `gorm:"parent_id"`
	Path      string       `gorm:"path"`
}

func (n *Node) TableName() string {
	return "nodes"
}
