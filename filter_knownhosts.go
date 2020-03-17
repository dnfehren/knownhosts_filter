package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "os/exec"
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

    args := os.Args[1:]

    if len(args) == 0 {
        fmt.Println("need some filter strings")
        os.Exit(1)
    }

    if err != nil {
        log.Fatal(err)
    }

    inFile, err := os.Open(knownhostsPath)
    if err != nil {
        log.Fatal(err)
    }

    bakPath := knownhostsPath + ".bak"
    bakFile, err := os.Create(bakPath)
    if err != nil {
        log.Fatal(err)
    }

    _, err = io.Copy(bakFile, inFile)
    if err != nil {
        log.Fatal(err)
    }
    defer bakFile.Close()

    newPath := knownhostsPath + ".new"
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

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    newFile.Close()

    // already tired of all the extra file stuff, punting here

    cmd := exec.Command("cp", "-f", newPath, knownhostsPath)
    cmd.Run()
    fmt.Println("cleaned ", cleanedNum)
}
