package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/karrick/godirwalk"
)

func count(path string) {
	count := 0
	err := godirwalk.Walk(path, &godirwalk.Options{
		Callback: func(pathname string, de *godirwalk.Dirent) error {
			if de.IsDir() {
				return nil
			}
			if strings.HasSuffix(strings.ToLower(de.Name()), ".jpg") {
				count += 1
			}
			return nil
		}})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}

func main() {
	count(os.Args[1])
}
