package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	repo, tag := detectRepoAndTag()
	image := "docker.io/" + repo + ":" + tag
	fmt.Println(image)
}

func detectRepoAndTag() (repo, tag string) {
	repo, err := runCmdAndGetOutput("git", "config", "--get", "remote.origin.url")
	if err != nil {
		log.Fatal("error detecting git repo", err)
	}

	if index := strings.LastIndex(repo, ":"); index != -1 {
		repo = repo[index+1:]
	}

	if strings.HasSuffix(repo, ".git") {
		repo = repo[:len(repo)-4]
	}

	tag, err = runCmdAndGetOutput("git", "rev-parse", "HEAD")
	if err != nil {
		log.Fatal("error detecting git HEAD revision", err)
	}

	if len(tag) > 7 {
		tag = tag[:7]
	}

	return repo, tag
}

func runCmdAndGetOutput(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}
