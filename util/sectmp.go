package util

import (
	"crypto/rand"
	"fmt"
	"os"
	"path"
)

// Create a secure temporary directory and return its full path.
func SecTempDir() (string, error) {
	basePath := os.TempDir()
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	// Prepend PID to make it easier to cleanup stale temp files on crashes.
	tmpDir := path.Join(basePath, fmt.Sprintf("%d-%x", os.Getpid(), buf))
	if err := os.Mkdir(tmpDir, os.ModeDir|os.ModeTemporary|0700); err != nil {
		return "", nil
	}
	return tmpDir, nil
}
