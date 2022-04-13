package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	RecursiveDirectoryCrawl("./2021-obsidian")
}

func RecursiveDirectoryCrawl(dirName string) {
	files, err := ioutil.ReadDir(dirName)
	handleError(err)

	for _, f := range files {
		fileOrDir, err := os.Stat(dirName + "/" + f.Name())
		handleError(err)
		switch mode := fileOrDir.Mode(); {
		case mode.IsDir():
			RecursiveDirectoryCrawl(dirName + "/" + f.Name())
		case mode.IsRegular():
			s := strings.Split(f.Name(), ".")
			if len(s) > 1 && s[1] == "md" {
				dat, err := os.ReadFile(dirName + "/" + f.Name())
				handleError(err)
				handleNotion(dirName+"/"+f.Name(), string(dat))
			}

		}
	}
}

func handleNotion(path string, data string) {
	// fmt.Println("going to publish to notion")
	// categorise
	s := strings.ReplaceAll(path, "./2021-obsidian/2021/", "")
	fmt.Println(s)
	// fmt.Println(data)
}

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}
