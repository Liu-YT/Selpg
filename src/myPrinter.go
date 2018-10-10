// 模拟打印机

package main

import (
    "os"
    "io"
    "bufio"
)

func main() {
	var reader *bufio.Reader
    file,err := os.Open("../data/printerTxt.txt")
    if err != nil {
    	panic(err)
    }

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
		_, errW := file.Write(line)
		if errW != nil {
			os.Stderr.Write([]byte("Write bytes to file fail\n"))
			os.Exit(0)
		}
    }
    file.Close()
}