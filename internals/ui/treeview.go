package ui

import (
	"github.com/faizalam/tree-top/internals/explorer"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TreeView struct {
	view    *tview.TreeView
	service *explorer.Service
}

func NewTreeView(service *explorer.Service) *TreeView {
	tv := tview.NewTreeView()
	tv.SetGraphics(true)
	tv.SetBorder(true)
	tv.SetTitle(" tree-htop ")
	return &TreeView{view: tv, service: service}
}

func (t *TreeView) Build(selectedPath string) {
	rootNode := t.buildNode(t.service.Tree())
	rootNode.SetExpanded(true)
	t.view.SetRoot(rootNode)
	// t.view.SetCurrentNode(rootNode)
	if selectedPath != "" {
		t.setCurrentByPath(selectedPath)
	} else {
		t.view.SetCurrentNode(rootNode)
	}
}

func (t *TreeView) buildNode(n *explorer.Node) *tview.TreeNode {
	tn := tview.NewTreeNode(n.Name).
		SetReference(n).
		SetSelectable(true)
	if n.Type == explorer.Directory {
		tn.SetColor(tcell.ColorGreen)
		// tn.SetAttributes(tcell.AttrReverse)
	}
	for _, child := range n.Children {
		tn.AddChild(t.buildNode(child))
	}
	return tn
}

func (t *TreeView) setCurrentByPath(path string) {
	var found *tview.TreeNode
	var walk func(node *tview.TreeNode)
	walk = func(node *tview.TreeNode) {
		if found != nil {
			return
		}
		if ref, ok := node.GetReference().(*explorer.Node); ok && ref.Path == path {
			found = node
			return
		}
		for _, c := range node.GetChildren() {
			walk(c)
		}
	}
	walk(t.view.GetRoot())
	if found != nil {
		t.view.SetCurrentNode(found)
	}
}
