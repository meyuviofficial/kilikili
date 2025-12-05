package utils

import (
	"math/rand"
	"os"
)

func MakeDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, 0755)
	}
	return nil
}

func RandomNameGenerator(nameList []string) string {
	return nameList[rand.Intn(len(nameList))]
}