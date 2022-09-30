package md5

import (
	"fmt"
	"testing"
)

func TestStringToMd5(t *testing.T) {
	s := "simple-utils"
	md5, err := StringToMd5(s)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(md5)
}

func TestFileToMd5Example(t *testing.T) {
	md5, err := FileToMd5("/home/simple.txt")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(md5)
}
