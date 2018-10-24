// 模拟打印机

package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	file, err := os.OpenFile("../data/printerText.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	for {
		line, errR := reader.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			} else {
				os.Stderr.Write([]byte("Read bytes from reader fail\n"))
				os.Exit(0)
			}
		}
		_, errW := writer.Write(line)
		if errW != nil {
			os.Stderr.Write([]byte("Write bytes to file fail\n"))
			os.Exit(0)
		}
		writer.Flush()
	}
}