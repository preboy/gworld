package game

import (
	"io/ioutil"
)

func ReadFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}