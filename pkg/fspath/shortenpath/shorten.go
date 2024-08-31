package shortenpath

import (
	"path/filepath"
	"strings"
)

// Shorten takes a file path and returns a tuple.
// The path with the last directory removed and last directory that was removed.
func Shorten(path string) (string, string) {
	if path == "" {
		return "", ""
	}

	// Clean the path to ensure it's in a consistent format.
	cleanedPath := filepath.Clean(path)

	// Split the path into its components.
	pathSegments := strings.Split(cleanedPath, string(filepath.Separator))

	// Extract the last segment (the part to be removed).
	lastSegment := pathSegments[len(pathSegments)-1]

	// Reconstruct the path without the last segment.
	shortenedPath := filepath.Join(pathSegments[:len(pathSegments)-1]...)

	return shortenedPath, lastSegment
}
