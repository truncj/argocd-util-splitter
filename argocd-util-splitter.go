package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var lines, kind, name []string

func main() {

	var (
		src = flag.String("src", "/home/backup.yaml", "path of argocd backup yaml")
		dst = flag.String("dst", "/tmp/", "output directory")
	)

	flag.Parse()

	yamlFile, err := ioutil.ReadFile(*src)
	if err != nil {
		fmt.Println("Unable to read source file -", err)
		os.Exit(1)
	}

	manifests := strings.Split(string(yamlFile), "---\n")

	reKind := regexp.MustCompile("kind: (.*)")
	reName := regexp.MustCompile("  name: (.*)")

	for _, manifest := range manifests {

		scanner := bufio.NewScanner(strings.NewReader(manifest))

		for scanner.Scan() {
			line := scanner.Text()
			lines = append(lines, line)

			if kind == nil {
				kind = reKind.FindStringSubmatch(line)
			}

			if name == nil {
				name = reName.FindStringSubmatch(line)
			}
		}

		text := strings.Join(lines, "\n")
		path := filepath.Join(*dst, fmt.Sprintf("%s-%s.yaml", name[1], kind[1]))

		err := ioutil.WriteFile(path, []byte(text), 0644)
		if err != nil {
			fmt.Println("Unable to write file -", err)
			os.Exit(1)
		}

		// reset manifest details to nil
		lines, name, kind = nil, nil, nil
	}

	fmt.Println(fmt.Sprintf("Success - ArgoCD individual backup manifests have been written to %s", *dst))
	os.Exit(0)
}
