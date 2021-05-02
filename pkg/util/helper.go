package util

import (
	"os"
	"time"
)

func FindElementExist(key string, el []string) bool {
	for _, item := range el {
		if item == key {
			return true
		}
	}
	return false
}

func GeneratedFileName() string {
	currentTime := time.Now()
	currentTimeString := currentTime.Format("20060102150405")
	return currentTimeString
}

func CheckDirectoryExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true

}
