package main

import (
	"os"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func IsFileExist(filename string) bool {
	fileInfo, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !fileInfo.IsDir()
}
