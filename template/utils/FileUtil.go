package utils

import (
	"os"
	"io/ioutil"
)

func ReadFile(path string) (str string, err error) {
	fi, err := os.Open(path)
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	str = string(fd)
	return

}