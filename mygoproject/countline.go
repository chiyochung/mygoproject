package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	path := "C:\\Users\\lili0\\go"
	key := ".go"
	t := time.Now()
	fmt.Println(myReadLine(path, key))
	fmt.Println(time.Since(t))
}

func myReadLine(path, key string) (sum int) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, fileInfo := range files {
			newPath := path + "\\" + fileInfo.Name()
			if fileInfo.IsDir() {
				sum += myReadLine(newPath, key)
			} else if strings.Contains(fileInfo.Name(), key) {
				file, err := os.Open(newPath)
				if err == nil {
					fileBuf := bufio.NewReader(file)
					lineSum := 0
					for {
						_, err := fileBuf.ReadString('\n')
						lineSum++
						if err == io.EOF {
							break
						}
					}
					file.Close()
					sum += lineSum
				}
			}
		}
	}
	return sum

}
