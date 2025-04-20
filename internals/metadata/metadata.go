package metadata

import (
	"fmt"
	"time"
)

// FormatSize returns a human-readable file size.
func FormatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

// FormatTime returns a formatted timestamp.
func FormatTime(t time.Time) string {
	return t.Format("Jan 02 2006 15:04")
}
