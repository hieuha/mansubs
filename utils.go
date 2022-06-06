package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func readLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()

	lines := make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines, sc.Err()
}

func parseTech(techFile string) []Target {
	var targets []Target
	target := Target{}
	lines, err := readLines(techFile)
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range lines {
		line := strings.Split(line, "|")
		u := strings.ReplaceAll(line[1], "~", "")
		u2, _ := url.Parse(u)
		subdomain := u2.Host
		tech := strings.TrimSpace(strings.Split(line[2], "/")[0])
		target.Subdomain = subdomain
		target.Technology = tech
		targets = append(targets, target)
	}
	return targets
}
