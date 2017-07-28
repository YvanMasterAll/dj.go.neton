package utils

import (
	"runtime"
	"time"
	"path/filepath"
	"os"
)

func LearnOS() string {
	return runtime.GOOS
}

func GetTime() string {
	timestamp := time.Now().Unix()
	return time.Unix(timestamp, 0).Format("2006-01-02 03:04:05 PM")
}

func GetCurrentPath() (string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}