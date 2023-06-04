package utils

import "os"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	// 文件或者目录存在
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
