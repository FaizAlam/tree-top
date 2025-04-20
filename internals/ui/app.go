// internal/ui/app.go
package ui

import (
	"context"

	"github.com/faizalam/tree-top/internals/explorer"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	app      *tview.Application
	treeView *TreeView
	detail   *DetailPanel
	layout   *tview.Flex
}

func NewTviewApp(service *explorer.Service) *App {
	app := tview.NewApplication()
	treeView := NewTreeView(service)
	detail := NewDetailPanel()

	// initial render, no selection override
	treeView.Build("")
	detail.Update(service.Tree())

	// setup layout
	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(treeView.view, 0, 1, true).
		AddItem(detail.view, 8, 0, false)

	// capture keys for expand/collapse and preserve selection
	treeView.view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		node := treeView.view.GetCurrentNode()
		if node == nil {
			return event
		}
		ref, _ := node.GetReference().(*explorer.Node)
		selPath := ref.Path

		switch event.Key() {
		case tcell.KeyRight:
			if ref.Type == explorer.Directory {
				service.Expand(ref.Path)
				treeView.Build(selPath)
			}
			return nil
		case tcell.KeyLeft:
			if ref.Type == explorer.Directory {
				service.Collapse(ref.Path)
				treeView.Build(selPath)
			}
			return nil
		case tcell.KeyRune:
			if event.Rune() == 'q' || event.Rune() == 'Q' {
				app.Stop()
				return nil
			}
		}
		return event
	})

	// update detail panel on selection change
	treeView.view.SetChangedFunc(func(node *tview.TreeNode) {
		if ref, ok := node.GetReference().(*explorer.Node); ok {
			detail.Update(ref)
		}
	})

	return &App{app: app, treeView: treeView, detail: detail, layout: layout}
}

func (a *App) Run(ctx context.Context) error {
	return a.app.SetRoot(a.layout, true).Run()
}
