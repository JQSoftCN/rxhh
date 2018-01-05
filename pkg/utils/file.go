package utils

import (
	"os"
	"log"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Dir(dir string) string {
	if PathExists(dir) {
		return dir
	}
	//mk dir
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Println(err)
	}
	return dir
}
