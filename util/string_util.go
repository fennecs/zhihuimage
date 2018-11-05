package util

import "strings"

func Trim(dir string) string {
	rootDir := dir
	rootDir = strings.Trim(rootDir, "'")
	rootDir = strings.Trim(rootDir, "\"")
	return rootDir
}