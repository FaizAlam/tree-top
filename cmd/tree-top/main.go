package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/faizalam/tree-top/internals/explorer"
	"github.com/faizalam/tree-top/internals/fs"
	"github.com/faizalam/tree-top/internals/ui"
)

func main() {
	ctx := context.Background()

	// Determine working directory
	dirFlag := flag.String("dir", "", "directory to explore (default to current directory)")
	flag.Parse()
	var dir string
	if *dirFlag != "" {
		dir = *dirFlag
	} else {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("failed to get working directory: %v", err)
		}
		dir = wd
	}

	info, err := os.Stat(dir)
	if err != nil {
		log.Fatalf("invalid directory %q: %v", dir, err)
	}
	if !info.IsDir() {
		log.Fatalf("%q is not a directory", dir)
	}

	// Initialize filesystem repository
	repo := fs.NewLocalFS()

	// Initialize explorer service
	svc := explorer.NewService(repo)

	// Load root directory
	if err := svc.LoadRoot(dir); err != nil {
		log.Fatalf("failed to load root directory: %v", err)
	}

	// Initialize and run TUI app
	app := ui.NewTviewApp(svc)
	if err := app.Run(ctx); err != nil {
		log.Fatalf("application error: %v", err)
	}
}
