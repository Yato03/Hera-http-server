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
	path, err := getDirectoryPathFromConfiguration()
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
	path, err := getDirectoryPathFromConfiguration()
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

func ReadHTML(relativePath string) (string, error) {
	path, err := getContentPathFromConfiguration()
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

func MakeConfigurationFile(directoryPath string, contentPath string) error {
	file, err := os.Create(CONFIGURATION_FILE)
	if err != nil {
		return err
	}
	defer file.Close()

	if directoryPath != "" {
		_, err = file.WriteString("directory:" + directoryPath + "\n")
		if err != nil {
			return err
		}
	}

	if contentPath != "" {
		_, err = file.WriteString("content:" + contentPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func getConfiguration() (map[string]string, error) {
	file, err := os.ReadFile(CONFIGURATION_FILE)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(file), "\n")
	configuration := make(map[string]string)

	for _, line := range lines {
		keyValue := strings.Split(line, ":")
		if len(keyValue) == 2 {
			configuration[keyValue[0]] = keyValue[1]
		}
	}

	return configuration, nil
}

func getDirectoryPathFromConfiguration() (string, error) {
	configuration, err := getConfiguration()
	if err != nil {
		return "", err
	}

	if configuration["directory"] == "" {
		return "", errors.New("directory not found in configuration file")
	}

	return configuration["directory"], nil
}

func getContentPathFromConfiguration() (string, error) {
	configuration, err := getConfiguration()
	if err != nil {
		return "", err
	}

	if configuration["content"] == "" {
		return "", errors.New("content not found in configuration file")
	}

	return configuration["content"], nil
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
