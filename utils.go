package main

import (
	"os"
	"strconv"
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

func StrToInt(param string) int {
	result, err := strconv.Atoi(param)
	PanicIfError(err)
	return result
}
