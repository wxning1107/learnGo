package utils

import "os"

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
