package cache

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"

	"github.com/teleport-network/teleport-relayer/app/types"
)

func TestCacheFileWriter_Write(t *testing.T) {
	cfgDirName := ".xibc-relayer"
	userDir, _ := os.UserHomeDir()
	homeDir := filepath.Join(userDir, cfgDirName)
	dir := "cache"
	filename := "teleport.json"
	writer := NewCacheFileWriter(homeDir, dir, filename)
	pts := []types.PacketDetail{}
	pts = append(pts, types.PacketDetail{
		Sequence:  1,
		SrcChain:  "rinkeby",
		DestChain: "teleport",
	})
	pts = append(pts, types.PacketDetail{
		Sequence:  1,
		SrcChain:  "rinkeby",
		DestChain: "teleport",
	})
	pts = append(pts, types.PacketDetail{
		Sequence:  1,
		SrcChain:  "rinkeby",
		DestChain: "teleport",
	})
	pts = append(pts, types.PacketDetail{
		Sequence:  1,
		SrcChain:  "rinkeby",
		DestChain: "teleport",
	})
	writer.WriteErrRelay(pts,true)
}

func TestCacheFileWriter_Read(t *testing.T) {
	cfgDirName := ".xibc-relayer"
	userDir, _ := os.UserHomeDir()
	homeDir := filepath.Join(userDir, cfgDirName)
	dir := "cache"
	filename := "teleport.json"
	writer := NewCacheFileWriter(homeDir, dir, filename)
	p, err := writer.GetErrRelay()
	require.NoError(t, err)
	fmt.Println(p)
	writer.WriteErrRelay(p, false)
}
