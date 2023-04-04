package service

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func RandomInt(len int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len) + 1
}

func RandomFile(folderPath string) (string, error) {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return "", fmt.Errorf("folder %s is empty", folderPath)
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(files))
	randomFile := files[randomIndex].Name()

	return randomFile, nil
}
