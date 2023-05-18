package util

import (
	"fmt"
	"os"
)

func bytesToKilobytes(bytes int64) string {
	kilobytes := bytes / 1024
	return fmt.Sprintf("%dKB", kilobytes)
}

func FileSize(fileDir string) string {
	fi, err := os.Stat(fileDir)
	if err != nil {
		fmt.Println(err)
	}
	size := fi.Size()
	return bytesToKilobytes(size)
}
