package explorer

import (
	"fmt"
	"sync"
)

// Service orchestrates the directory tree and expansion state.
type Service struct {
	repo     Repository
	mu       sync.RWMutex
	root     *Node
	expanded map[string]bool
}

// NewService creates a new Explorer service.
func NewService(repo Repository) *Service {
	return &Service{
		repo:     repo,
		expanded: make(map[string]bool),
	}
}

// LoadRoot initializes the tree from the given path.
func (s *Service) LoadRoot(path string) error {
	nodes, err := s.repo.List(path)
	if err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.root = &Node{Name: path, Path: path, Type: Directory, Children: nodes}
	s.expanded[path] = true
	return nil
}

// Expand loads and expands the node at the given path.
func (s *Service) Expand(path string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.expanded[path] {
		return nil
	}
	node := s.findNode(s.root, path)
	if node == nil {
		return fmt.Errorf("node not found: %s", path)
	}
	children, err := s.repo.List(path)
	if err != nil {
		return err
	}
	node.Children = children
	s.expanded[path] = true
	return nil
}

// Collapse collapses the node at the given path.
func (s *Service) Collapse(path string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.expanded[path] {
		return nil
	}
	node := s.findNode(s.root, path)
	if node == nil {
		return fmt.Errorf("node not found: %s", path)
	}
	node.Children = nil
	delete(s.expanded, path)
	return nil
}

// Tree returns the root node of the tree.
func (s *Service) Tree() *Node {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.root
}

// findNode recursively searches for a node by path.
func (s *Service) findNode(current *Node, path string) *Node {
	if current.Path == path {
		return current
	}
	for _, child := range current.Children {
		if n := s.findNode(child, path); n != nil {
			return n
		}
	}
	return nil
}
