package pars

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

func FileToString(filePath string) ([]string, error) {
	content, err := ReadFile(filePath)
	listStrings := strings.Split(string(content), "\n")
	return listStrings, err
}

func ReadFile(filePath string) ([]byte, error) {
	log.Printf("Extract data from file `%v`", filePath)
	content, err := os.ReadFile(filePath)
	return content, err
}

func JSON[T any](storage *T, r io.Reader) error {
	content, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(content, &storage); err != nil {
		return err
	}
	return nil
}
