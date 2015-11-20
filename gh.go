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
		user_and_repo := match[1]

		out, err := exec.Command("git", "symbolic-ref", "--short", "-q", "HEAD").Output()
		branch_path := ""
		if err == nil {
			branch := strings.TrimSpace(string(out))
			if branch != "master" {
				branch_path = fmt.Sprintf("/tree/%s", branch)
			}
		}

		if len(match) > 1 {
			gh_url := fmt.Sprintf("https://github.com/%s%s", user_and_repo, branch_path)
			if _, err := exec.Command("open", gh_url).Output(); err == nil {
				os.Exit(0)
			}
		}
	}

	os.Exit(1)
}
