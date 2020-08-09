package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GetFileSize(path string) (int, error) {
	fi, err := os.Stat(path)
	if err == nil {
		return int(fi.Size()), nil
	}

	return 0, err
}

func JoinString(s1, s2 string) []string {
	s3 := s1 + s2
	s4 := fmt.Sprintf("%s%s", s1, s2)
	s5 := strings.Join([]string{s1, s2}, "")

	var b1 bytes.Buffer
	b1.WriteString(s1)
	b1.WriteString(s2)
	s6 := b1.String()

	var b2 strings.Builder
	b2.WriteString(s1)
	b2.WriteString(s2)
	s7 := b2.String()

	return []string{s3, s4, s5, s6, s7}
}
