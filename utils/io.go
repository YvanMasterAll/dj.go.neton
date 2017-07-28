package utils

import (
	"os"
)

func fileIsExist(f string) (bool) {
	r := true;
	if _, err := os.Stat(f); os.IsNotExist(err) {
		r = false;
	}
	return r;
}

func Append(f string, i string) (bool) {
	if(!fileIsExist(f)) {
		_, err := os.Create(f)
		if err != nil {
			return false
		}
	}

	fi, err := os.OpenFile(f, os.O_WRONLY, 0644)
	if err != nil {
		return false
	}

	//查找文件末尾的偏移量
	n, _ := fi.Seek(0, os.SEEK_END)
	_, err = fi.WriteAt([]byte(i + "\r\n"), n)

	defer fi.Close()

	return true
}