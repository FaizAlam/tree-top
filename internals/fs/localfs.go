package fs

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/faizalam/tree-top/internals/explorer"
)

// LocalFS implements explorer.Repository using the local filesystem.
type LocalFS struct{}

// NewLocalFS constructs a new LocalFS.
func NewLocalFS() *LocalFS {
	return &LocalFS{}
}

// List reads directory entries for lazy-loading.
func (l *LocalFS) List(dir string) ([]*explorer.Node, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	nodes := make([]*explorer.Node, 0, len(entries))
	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			continue
		}
		node := &explorer.Node{
			Name:     e.Name(),
			Path:     filepath.Join(dir, e.Name()),
			Type:     fileType(e, info),
			Size:     info.Size(),
			Mode:     info.Mode().String(),
			Created:  info.ModTime(),
			Modified: info.ModTime(),
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func fileType(e fs.DirEntry, info fs.FileInfo) explorer.ItemType {
	if info.IsDir() {
		return explorer.Directory
	}
	return explorer.File
}
