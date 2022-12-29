package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"strconv"
)

func main() {
	/*
		keywordptr := flag.String("keyword", "", "string flag")
		lengthptr := flag.Int("length", 16, "string flag")
		boolptr := flag.Bool("yesorno", false, "string flag")
	*/
	flag.Parse()

	keyword, offset, length := getArgs()
	password := generate(keyword, length, offset)

	template := "关键字:【%s】\n起始坐标：【%d】\n密码长度：【%d】\n所生成的密码为：\n%s\n"
	fmt.Printf(template, keyword, offset, length, password)
}

func getArgs() (string, int, int) {
	args := flag.Args()
	keyword := ""
	var length, offset int
	var err error
	if flag.NArg() > 0 {
		keyword = args[0]
	}
	if flag.NArg() > 1 {
		offset, err = strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
	} else {
		offset = 0
	}
	if flag.NArg() > 2 {
		length, err = strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}
	} else {
		length = 16
	}

	return keyword, offset, length
}

func generate(keyword string, length int, offset int) string {
	sha1 := sha1.New()
	io.WriteString(sha1, keyword)
	hexenc := hex.EncodeToString(sha1.Sum(nil))
	b64enc := base64.StdEncoding.EncodeToString([]byte(hexenc))
	return b64enc[offset : offset+length]
}
