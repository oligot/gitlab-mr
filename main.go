package main

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	// configuration
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.config/gitlab-mr")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	if !viper.IsSet("token") {
		log.Fatalf("No token in the config file")
	}

	// parse git remote url
	remote, err := exec.Command("git", "remote", "get-url", "origin").Output()
	if err != nil {
		log.Fatal(err)
	}
	r := regexp.MustCompile(`@([\w.]+):([\w/-]+)\/([\w-]+)\.`)
	matched := r.FindStringSubmatch(string(remote))
	group, name := matched[2], matched[3]

	// git branch
	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		log.Fatal(err)
	}

	// merge request
	git := gitlab.NewClient(nil, viper.GetString("token"))
	options := &gitlab.ListProjectMergeRequestsOptions{
		SourceBranch: gitlab.String(strings.TrimSuffix(string(branch), "\n")),
	}
	mergeRequests, _, err := git.MergeRequests.ListProjectMergeRequests(group+"/"+name, options)
	if err != nil {
		log.Fatal(err)
	}
	if len(mergeRequests) == 1 {
		open.Start(mergeRequests[0].WebURL)
	} else {
		fmt.Println("No merge request found")
	}
}
