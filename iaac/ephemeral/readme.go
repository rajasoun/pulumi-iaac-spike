package ephemeral

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func pathToCurrentDir() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Err = %v", err)
		return "", err
	}
	return pwd, nil
}

func info(ctx *pulumi.Context) error {
	readmeBytes, err := loadReadmeFile()
	if err != nil {
		return err
	}

	ctx.Export("readme", pulumi.String(string(readmeBytes)))
	return nil
}

func loadReadmeFile() ([]byte, error) {
	dir, err := pathToCurrentDir()
	if err != nil {
		return nil, err
	}
	filePath := filepath.Join(dir, "/iaac/ephemeral/README.md")
	readmeBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return readmeBytes, nil
}
