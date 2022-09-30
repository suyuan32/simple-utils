package md5

import (
	"crypto/md5"
	"io"
	"os"

	"encoding/hex"
)

// StringToMd5 returns md5 from a string
func StringToMd5(str string) (string, error) {
	h := md5.New()
	_, err := h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil)), err
}

// FileToMd5 returns md5 from a file
func FileToMd5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return string(h.Sum(nil)), nil
}
