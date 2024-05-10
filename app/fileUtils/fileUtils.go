package fileUtils

import (
	"bytes"
	"compress/gzip"
	"errors"
	"os"
	"strings"
)

const (
	CONFIGURATION_FILE = ".config"
)

func ReadFile(relativePath string) (string, error) {
	path, err := getPathFromConfiguration()
	if err != nil {
		return "", err
	}
	absolutePath := path + "/" + relativePath
	file, err := os.ReadFile(absolutePath)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func WriteFile(relativePath string, content string) error {
	path, err := getPathFromConfiguration()
	if err != nil {
		return err
	}
	absolutePath := path + "/" + relativePath
	err = os.WriteFile(absolutePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func MakeConfigurationFile(path string) error {
	file, err := os.Create(CONFIGURATION_FILE)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("path:" + path)
	if err != nil {
		return err
	}
	return nil
}

func getPathFromConfiguration() (string, error) {
	file, err := os.ReadFile(CONFIGURATION_FILE)
	if err != nil {
		return "", err
	}

	pathKeyValue := strings.Split(string(file), ":")

	if len(pathKeyValue) != 2 {
		return "", errors.New("INVALID CONFIGURATION FILE")
	}

	return string(pathKeyValue[1]), nil
}

func CleanConfiguration() error {
	err := os.Remove(CONFIGURATION_FILE)
	if err != nil {
		return err
	}
	return nil
}

func Gzip(content string) (string, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	// Setting the Header fields is optional.
	zw.Name = "Encoded by gzip"
	zw.Comment = "Encoded by Heras"

	_, err := zw.Write([]byte(content))
	if err != nil {
		return "", err
	}

	// Make sure to close the gzip writer to flush any remaining data
	err = zw.Close()
	if err != nil {
		return "", err
	}

	return buf.String(), nil

}
