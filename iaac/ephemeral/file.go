package ephemeral

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func pathToCurrentDir() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Err = %v", err)
		return "", err
	}
	return pwd, nil
}

func loadFile(fileRelPath string) ([]byte, error) {
	dir, err := pathToCurrentDir()
	if err != nil {
		return nil, err
	}
	filePath := filepath.Join(dir, fileRelPath)
	readmeBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return readmeBytes, nil
}
