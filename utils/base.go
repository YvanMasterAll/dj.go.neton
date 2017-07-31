package utils

import (
	"runtime"
	"time"
	"path/filepath"
	"os"
	"strings"
	"log"
	"github.com/kardianos/osext"
	"os/exec"
	"bytes"
)

var (
	L = "neton.log"
	C = "neton.json"
	E = ""
)

func init() {
	f, err := GetExecName()
	if err != nil {
		log.Fatal(err)
	}
	L = f + ".log"
	C = f + ".json"

	e, err := GetCurrentPath()
	if err != nil {
		log.Fatal(err)
	}
	E = e
}

func LearnOS() string {
	return runtime.GOOS
}

func GetTime() string {
	timestamp := time.Now().Unix()
	return time.Unix(timestamp, 0).Format("2006-01-02 03:04:05 PM")
}

func GetCurrentPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}

func GetParentDirectory(dir string) string {
	return Substr(dir, strings.LastIndex(dir, "/"), len(dir) - strings.LastIndex(dir, "/"))
}

func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func GetExecName() (string, error) {
	fullexecpath, err := osext.Executable()
	if err != nil {
		return "", err
	}

	_, execname := filepath.Split(fullexecpath)
	ext := filepath.Ext(execname)
	name := execname[:len(execname)-len(ext)]

	return name, nil
}

func ExecCommand(c string, i string) string{
	cmd := exec.Command(c)

	in := bytes.NewBuffer(nil)
	cmd.Stdin = in//绑定输入
	var out bytes.Buffer
	cmd.Stdout = &out //绑定输出

	go func() {
		in.WriteString(i)//写入你的命令，可以有多行，"\n"表示回车
	}()

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}

	return out.String()
}