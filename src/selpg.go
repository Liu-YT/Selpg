package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"bufio"
	"os"
	"os/exec"
	"math"
	"github.com/spf13/pflag"
)

type selpg_args struct {
	startPage 	int
	endPage		int
	inFile		string
	pageLen		int
	pageType	bool	// true for -f, false for -lNumber
	outDestination	string
}

func main() {
	var args selpg_args

	// 定义并且获取参数
	getArgs(args)

	// 检查参数合法性
	checkArgs(args)

	// 执行命令
	processInput(args)
}

// 定义并且获取参数
func getArgs(args *selpg_args) {
	flag.IntVal(&(args.startPage), "s", "-1", "start page")
	flag.IntVal(&(args.endPage), "e", "-1", "end page")
	flag.IntVal(&(args.pageLen), "l", "72", "the length of page")
	flag.BoolVar(&(args.pageType), "f", false, "page type")
	flag.StringVar(&(args.outDestination), "d", "", "print destination")
	flag.Parse()

	other := flag.Args()	// 其余参数
	if(len(other) > 0) {
		args.inFile = other[0]
	}
	else {
		args.inFile = ""
	}
}

// 检查参数合法性
func checkArgs(args *selpg_args) {
	if args.startPage == -1 || args.endPage == -1 {
		os.Stderr.Write([]byte("You shouid input like selpg -sNumber -eNumber ... \n"))
		os.Exit(0)
	}

	if args.startPage < 1 || args.startPage > math.MaxInt32 {
		os.Stderr.Write([]byte("You should input valid start page\n"))
		os.Exit(0)
	}

	if args.endPage < 1 ||  args.endPage > math.MaxInt32 || args.endPage < startPage{
		os.Stderr.Write([]byte("You should input valid end page\n"))
		os.Exit(0)
	}

	if (!args.pageType) && (args.pageLen < 1 || args.pageLen > math.MaxInt32) {
		os.Stderr.Write([]byte("You should input valid page length\n"))
		os.Exit(0)
	}
}

// 执行命令
func processInput(args *selpg_args) {

	// read the file
	var reader *bufio.Reader

	if args.inFile == "" {
		reader = bufio.NewReader(os.Stdin)
	} 
	else {
		fileIn, err := os.Open(args.inFile)
		defer fileIn.Close()
		if err != nil {
			os.Stderr.Write([]byte("Open file [%s] error\n", args.inFile))
			os.Exit(0)
		}
		reader = bufio.NewReader(fileIn)
	}

	// output the file
	if args.outDestination == "" {
		// 输出到当前命令行
		outputCurrent(reader, args);
	}
	else {
		// 输出到目的地
		outputToDest(reader, args);
	}
}

// 输出到当前命令行
func outputCurrent(reader *bufio.Reader, args *selpg_args) {
	writer := bufio.NewWriter(os.Stdout)

	lineCtr = 0
	pageCtr = 1
	if args.pageType == true {
		// -f page type
		for {
			char, errR := reader.ReadByte()
			if err != nil {
				if errR == io.EOF {
					break
				}
				else {
					os.Stderr.Write([]byte("Read byte from reader fail\n"))
					os.Exit(0)
				}
			}
			if pageCtr >= args.startPage && pageCtr <= args.endPage {
				errW := writer.WriteByte(char)
				if errW != nil {
					os.Stderr.Write([]byte("Write byte to out fail\n"))
					os.Exit(0)
				}
				writer.Flush()	// 刷新缓存流，输出到控制台
			}
			if char == '\f' {
				pageCtr++
			}
		}
	}
	else {
		// -lNumber page type
		// page len is 72 or the number in -lNumber
		for{
			strLine, errR := reader.ReadString('\n')
			if err != nil {
				if errR == io.EOF {
					break
				}
				else {
					os.Stderr.Write([]byte("Read string line from reader fail\n"))
					os.Exit(0)
				}
			}

			lineCtr++

			if pageCtr >= args.startPage && pageCtr <= args.endPage {
				errW := writer.Write(strLine)
				if errW != nil {
					os.Stderr.Write([]byte("Write string line to out fail\n"))
					os.Exit(0)
				}
				writer.Flush()	// 刷新缓存流，输出到控制台
			}
			if lineCtr == args.pageLen {
				lineCtr = 0
				pageCtr++ 
			}
		}
	}

	writer.Close()

	checkPageNum(args, pageCtr)
}

// 输出到指定目的地	
func outputToDest(reader *bufio.Reader, args *selpg_args) {
	
}

// 检查开始页号与结束页号的合理性
func checkPageNum(args *selpg_args, pageCtr int) {

	if pageCtr < args.startPage {
		os.Stderr.writer([]byte("Start page is bigger than the total page num"))
		os.exit(0)
	}

	if pageCtr > args.endPage {
		os.Stderr.writer([]byte("End page is bigger than the total page num"))
		os.exit(0)
	}
}