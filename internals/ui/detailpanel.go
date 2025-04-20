package ui

import (
	"fmt"
	"strings"

	"github.com/faizalam/tree-top/internals/metadata"

	"github.com/faizalam/tree-top/internals/explorer"

	"github.com/rivo/tview"
)

type DetailPanel struct {
	view *tview.TextView
}

func NewDetailPanel() *DetailPanel {
	tv := tview.NewTextView()
	tv.SetDynamicColors(true)
	tv.SetBorder(true)
	tv.SetTitle(" Details ")
	return &DetailPanel{view: tv}
}

func humanPermsFromString(modeStr string) string {
	if len(modeStr) < 10 {
		return "invalid mode string"
	}
	// Helper to decode a triad like "rwx"
	decode := func(triad string) string {
		var acts []string
		if triad[0] == 'r' {
			acts = append(acts, "read")
		}
		if triad[1] == 'w' {
			acts = append(acts, "write")
		}
		// treat x, s, t all as execute bits
		if triad[2] == 'x' || triad[2] == 's' || triad[2] == 't' {
			acts = append(acts, "execute")
		}
		if len(acts) == 0 {
			return "no permissions"
		}
		return strings.Join(acts, " & ")
	}

	owner := decode(modeStr[1:4])
	group := decode(modeStr[4:7])
	others := decode(modeStr[7:10])

	return fmt.Sprintf(
		"Owner: %s; Group: %s; Others: %s",
		owner, group, others,
	)
}

func (d *DetailPanel) Update(node *explorer.Node) {
	d.view.Clear()
	fmt.Fprintf(d.view, "[yellow]Permissions:[white] %s (%s)\n", node.Mode, humanPermsFromString(node.Mode))
	fmt.Fprintf(d.view, "[yellow]Name:[white] %s\n", node.Name)
	fmt.Fprintf(d.view, "[yellow]Type:[white] %s\n", node.Type)
	fmt.Fprintf(d.view, "[yellow]Size:[white] %s\n", metadata.FormatSize(node.Size))
	fmt.Fprintf(d.view, "[yellow]Modified:[white] %s\n", metadata.FormatTime(node.Modified))
	fmt.Fprintf(d.view, "[yellow]Path:[white] %s\n", node.Path)
}
