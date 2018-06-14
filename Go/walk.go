package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func count(path string) {
	count := 0
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(info.Name()), ".jpg") {
			count += 1
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}

func main() {
	count(os.Args[1])
}
