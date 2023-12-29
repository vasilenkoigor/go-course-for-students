package main

import (
	"fmt"
	"lecture03_homework/file_copier"
	"os"
)

var flagsParser file_copier.FlagsParser = file_copier.UnixCmdFlagsParser{}
var fileCopier file_copier.FileCopier = file_copier.UnixDDCopier{}

func main() {
	opts, err := flagsParser.Parse()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "can not parse flags:", err)
		os.Exit(1)
	}

	if err := fileCopier.Copy(*opts); err != nil {
		fmt.Println(err)
	}
}
