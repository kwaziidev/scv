package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

var dir map[string][]string
var mode []string
var modeDir map[string][]string

func init() {
	dir = make(map[string][]string)
	dir["api"] = []string{"api"}
	dir["assets"] = []string{"assets"}
	dir["build"] = []string{"build/ci", "build/package"}
	dir["cmd"] = []string{"cmd"}
	dir["configs"] = []string{"configs"}
	dir["deployments"] = []string{"deployments"}
	dir["docs"] = []string{"docs"}
	dir["examples"] = []string{"examples"}
	dir["githooks"] = []string{"githooks"}
	dir["init"] = []string{"init"}
	dir["internal"] = []string{"internal/app", "internal/pkg"}
	dir["pkg"] = []string{"pkg"}
	dir["scripts"] = []string{"scripts"}
	dir["test"] = []string{"test"}
	dir["third_party"] = []string{"third_party"}
	dir["tools"] = []string{"tools"}
	dir["vendor"] = []string{"vendor"}
	dir["web"] = []string{"web/app", "web/static", "web/template"}
	dir["website"] = []string{"website"}

	mode = []string{"min", "std", "full"}

	modeDir = make(map[string][]string)
	modeDir["min"] = []string{"cmd", "configs", "docs", "internal", "pkg", "test"}
	modeDir["std"] = []string{"api", "assets", "cmd", "configs", "docs", "init", "internal", "pkg", "test"}
	modeDir["full"] = []string{}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("please input project name!")
	}

	projectName := args[1]
	if projectName == "" {
		log.Fatal("please input project name!")
	}

	if strings.ContainsAny(projectName, "./") {
		log.Fatal(`project name can not contains "." or "/"`)
	}

	wd, _ := os.Getwd()
	dstDir := filepath.Join(wd, "", projectName)
	if PathExists(dstDir) {
		log.Fatalf("%s already exists", projectName)
	}

	selectMode := ""
	prompt := &survey.Select{
		Message: "Choose a mode:",
		Options: mode,
	}
	survey.AskOne(prompt, &selectMode)

	if selectMode == "" {
		return
	}

	var wantMake []string
	keys := modeDir[selectMode]
	if len(keys) == 0 {
		for _, v := range dir {
			wantMake = append(wantMake, v...)
		}
	} else {
		for _, v := range keys {
			wantMake = append(wantMake, dir[v]...)
		}
	}

	for _, v := range wantMake {
		os.MkdirAll(projectName+"/"+v, os.ModePerm)
	}
	fmt.Println("generate success!")
}

func PathExists(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			return true
		} else {
			return false
		}
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
