package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage 'deb dist.tar.gz'")
		return
	}
	source := os.Args[1]
	version := debVersion(version(source))
	name := fmt.Sprintf("go_%s_%s.deb", version, debArch())
	input, err := os.Open(source)
	if err != nil {
		fmt.Println("Could not open source", err)
	}
	defer input.Close()
	result, err := os.Create(name)
	if err != nil {
		fmt.Println("Could not create deb", err)
		return
	}
	defer result.Close()
	createDeb(version, input, result)
}

func version(file string) string {
	start := 2
	end := strings.Index(file, ".linux-")
	return file[start:end]
}