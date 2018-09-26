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

// package main

// func main() {

// 	reader := bufio.NewReader(os.Stdin)
// 	f, err := os.OpenFile("receiverText.txt", os.O_WRONLY, 0666)

// 	if err != nil {
// 		os.Stderr.WriteString("fail to open receiver")
// 	}
// 	writer := bufio.NewWriter(f)
// 	for {
// 		line, err := reader.ReadBytes('\n')
// 		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
// 			if err == io.EOF {
// 				break
// 			}
// 			os.Stderr.Write([]byte("read bytes from Reader error\n"))
// 			os.Exit(5)
// 		}
// 		writer.Write(line)
// 		writer.Flush()
// 	}
// }
