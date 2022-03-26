package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./myDownloadImage_00_00_34")
	if err != nil {
		log.Fatal(err)
	}
	fileList := make([]bool, 5001)
	for _, f := range files {
		i, _ := strconv.Atoi(strings.Split(f.Name(), ".")[0])
		fileList[i] = true
	}
	for i, v := range fileList {
		if i != 0 && i <= 4500 && !v {
			// missing files
			fmt.Println(i)
		}
	}
}
