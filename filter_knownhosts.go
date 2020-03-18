package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
)

func containsAll(host string, criteria []string) bool {

	for _, criterion := range criteria {
		if !strings.Contains(host, criterion) {
			return false
		}
	}
	return true
}

func main() {
	knownhostsPath, err := homedir.Expand("~/.ssh/known_hosts")
	bakPath := knownhostsPath + ".bak"

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("need some filter strings")
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
	}

	os.Rename(knownhostsPath, bakPath)

	inFile, err := os.Open(bakPath)
	if err != nil {
		log.Fatal(err)
	}

	newPath := knownhostsPath
	newFile, err := os.Create(newPath)
	if err != nil {
		log.Fatal(err)
	}

	inFile.Seek(0, 0)
	scanner := bufio.NewScanner(inFile)
	cleanedNum := 0
	for scanner.Scan() {
		if containsAll(scanner.Text(), args) == true {
			cleanedNum = cleanedNum + 1
			continue
		}
		newFile.WriteString(scanner.Text())
	}
	newFile.Sync()
	inFile.Close()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	newFile.Close()

	fmt.Println("cleaned ", cleanedNum)
}
