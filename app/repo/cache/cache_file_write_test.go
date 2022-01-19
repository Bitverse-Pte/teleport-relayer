package cache

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCacheFileWriter_Write(t *testing.T) {
	cfgDirName := ".xibc-relayer"
	userDir, _ := os.UserHomeDir()
	homeDir := filepath.Join(userDir, cfgDirName)
	dir := "cache"
	filename := "teleport.json"
	writer := NewCacheFileWriter(homeDir, dir, filename)
	if err := writer.Write(1); err != nil {
		t.Fatal(err)
	}
}
