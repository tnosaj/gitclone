package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/pcarranza/sh-tools/git"
)

func main() {

	goflag := flag.Bool("g", false, "Flag to signify that its a go repo")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("One argument is required")
	}

	repo := args[0]
	fmt.Println(repo)

	u, err := git.Parse(repo)
	if err != nil {
		panic(err)
	}
	var path []string
	path = append(path, os.Getenv("HOME"))

	if *goflag {
		path = append(path, "go/src")
	} else {
		path = append(path, "GIT")

	}
	path = append(path, u.ToGoPath())
	directory := strings.Join(path, "/")
	directory = strings.TrimSuffix(directory, ".git")
	var cmd *exec.Cmd
	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		cmd = exec.Command("git", "pull")
		cmd.Dir = directory

	} else {
		os.MkdirAll(directory, os.ModePerm)
		cmd = exec.Command("git", "clone", repo, directory)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("could not pull git repo %s into %s: %s\n", repo, directory, err)
	}
	fmt.Println(directory)
}
