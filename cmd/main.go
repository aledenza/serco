package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

var changedFiles = map[string]struct{}{"go_mod": struct{}{}, "go_sum": struct{}{}}

func fail(msg string) {
	fmt.Printf("\033[91m%s\033[0m", msg)
	os.Exit(1)
}

func success(msg string) {
	fmt.Printf("\033[92m%s\033[0m", msg)
}

//go:embed template/*
var template embed.FS

func main() {
	flag.Usage = func() {
		fmt.Println(`
Usage: serco dst

Creates new project from template

positional arguments:
	dst			Project folder

options:
	-help		Show this message and exit
	`)
	}
	var projectDir string
	flag.Parse()
	switch flag.NArg() {
	case 1:
		projectDir = flag.Arg(0)
	default:
		flag.Usage()
		return
	}
	if projectDir == "" {
		fail("project dir is unset")
	}
	err := fs.WalkDir(template, "template", func(fpath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		pth := strings.Replace(fpath, "template", projectDir, 1)
		if _, ok := changedFiles[d.Name()]; ok {
			pth = strings.Replace(pth, d.Name(), strings.Replace(d.Name(), "_", ".", 1), 1)
		}
		if d.IsDir() {
			err = os.Mkdir(pth, 0755)
			if err != nil {
				return err
			}
		} else {
			data, _ := template.ReadFile(fpath)
			err = os.WriteFile(pth, data, 0666)
			if err != nil {
				return err
			}
		}
		return nil
	},
	)
	if err != nil {
		fail(err.Error())
	}
	success("Project `" + projectDir + "` created")
}
