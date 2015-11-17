package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	if _, err := os.Stat(".git"); err == nil {
		out, _ := exec.Command("git", "config", "remote.origin.url").Output()
		repo_url := strings.TrimSpace(string(out))

		re := regexp.MustCompile(`github\.com[:\/](.*?)(\.git)?$`)
		match := re.FindStringSubmatch(repo_url)

		if len(match) > 1 {
			gh_url := fmt.Sprintf("https://github.com/%s", match[1])
			if _, err := exec.Command("open", gh_url).Output(); err == nil {
				os.Exit(0)
			}
		}
	}

	os.Exit(1)
}
