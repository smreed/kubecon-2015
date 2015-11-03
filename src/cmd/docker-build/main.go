package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	repo, path, tag := detectRepoPathAndTag()
	dockerfiles := findDockerfiles()
	for dockerfile, image := range mapDockerfileToRepo(repo, path, tag, dockerfiles...) {
		_, err := runCmdAndGetOutput("docker", "build", "-t", image, "-f", dockerfile, ".")
		if err != nil {
			log.Println("error building", image, "from", dockerfile)
			log.Fatal(err)
		}
	}
}

func detectRepoPathAndTag() (repo, path, tag string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting working dir", err)
	}

	repo, err = runCmdAndGetOutput("git", "config", "--get", "remote.origin.url")
	if err != nil {
		log.Fatal("error detecting git repo", err)
	}

	if index := strings.LastIndex(repo, ":"); index != -1 {
		repo = repo[index+1:]
	}

	if strings.HasSuffix(repo, ".git") {
		repo = repo[:len(repo)-4]
	}

	path, err = runCmdAndGetOutput("git", "rev-parse", "--show-toplevel")
	if strings.HasPrefix(wd, path) {
		path = wd[len(path)+1:]
		path = strings.Replace(path, "/", "-", -1)
	} else {
		log.Fatal("Current directory is not child of top level", wd, path)
	}

	tag, err = runCmdAndGetOutput("git", "rev-parse", "HEAD")
	if err != nil {
		log.Fatal("error detecting git HEAD revision", err)
	}

	if len(tag) > 7 {
		tag = tag[:7]
	}

	return repo, path, tag
}

func runCmdAndGetOutput(name string, arg ...string) (string, error) {
	fmt.Println(">", name, strings.Join(arg, " "))
	cmd := exec.Command(name, arg...)

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

func findDockerfiles() []string {
	matches, err := filepath.Glob("Dockerfile*")
	if err != nil {
		log.Fatal("error finding Dockerfiles", err)
	}
	return matches
}

func mapDockerfileToRepo(base, path, tag string, dockerfile ...string) map[string]string {
	m := make(map[string]string)
	for _, f := range dockerfile {
		m[f] = generateRepoName(base, path, tag, f)
	}
	return m
}

func generateRepoName(base, path, tag, dockerfile string) string {
	base = "docker.io/" + base
	if path != "" {
		base = base + "-" + path
	}
	suffix := dockerfile
	if index := strings.LastIndex(suffix, "."); index != -1 {
		suffix = suffix[index+1:]
	} else {
		return base + ":" + tag
	}
	return base + "-" + suffix + ":" + tag
}
