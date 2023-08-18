package utils

import (
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"
)

func ReadPEMFileToString(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return "", errors.New("failed to decode PEM block")
	}
	return base64.StdEncoding.EncodeToString(block.Bytes), nil
}
