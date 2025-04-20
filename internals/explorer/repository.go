package explorer

import "time"

// ItemType indicates whether a Node is a file or directory.
//go:generate stringer -type=ItemType

type ItemType int

const (
    File ItemType = iota
    Directory
)

func (t ItemType) String() string {
    switch t {
    case File:
        return "File"
    case Directory:
        return "Directory"
    default:
        return "Unknown"
    }
}

// Node represents a file or directory in the tree.
type Node struct {
    Name     string
    Path     string
    Type     ItemType
    Size     int64
    Mode     string    // permissions string
    Created  time.Time // using ModTime for both Created and Modified
    Modified time.Time
    Children []*Node
}

// Repository defines lazy-loading FS operations.
type Repository interface {
    List(path string) ([]*Node, error)
}