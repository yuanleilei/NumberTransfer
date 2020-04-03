package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"./chinese2digits"
	"./rmblower2upper"
	"github.com/axgle/mahonia"
)

const (
	OldReplaceFlag = "零点"
	NewReplaceFlag = "0."
)

// 修改替换原文件中所有中文数字转为阿拉伯数字
func ReplaceFileLine(path string) {

	var result = ""

	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("read file err=", err)
		return
	}

	decoder := mahonia.NewDecoder("gb18030")
	if decoder == nil {
		fmt.Println("文件编码不存在!")
		return
	}

	r := bufio.NewReader(f)

	for {
		if buf, err := r.ReadBytes('\n'); err != nil {
			if err == io.EOF {
				break
			}
		} else {
			tmpStr := string(buf)

			str := decoder.ConvertString(tmpStr)

			// 处理零点->0.
			if strings.Contains(str, OldReplaceFlag) {
				str = strings.ReplaceAll(str, OldReplaceFlag, NewReplaceFlag)
			}

			dirStr := chinese2digits.TakeNumberFromString(str, false, false).(string)

			result += dirStr

		}
	}
	// 关闭文件
	f.Close()

	fmt.Println(result)

	// 打开文件
	fw, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) // os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		fmt.Println("write file err=", err)
		return
	}
	w.Flush()
	// 关闭文件
	fw.Close()
}

// 人民币小写转大写
func RMBLowerCaseToUpperCase(num float64) string {
	return rmblower2upper.ConvertNumToCny(12345.78)
}

func main() {
	path := "./result.txt"
	ReplaceFileLine(path)

	RMBLowerCaseToUpperCase(12345.78)
}
