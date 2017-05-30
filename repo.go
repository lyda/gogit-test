//
// repo.go
// Copyright (C) 2017 kevin <kevin@phrye.com>
//
// Distributed under terms of the GPL license.
//

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/speedata/gogit"
)

type Repo struct {
	Author        *gogit.Signature
	CommitMessage string
}

var Repos map[string]*Repo

func getRepos(repo_path string, info os.FileInfo, err error) error {
	if info != nil && info.IsDir() && strings.HasSuffix(repo_path, ".git") && repo_path != ".git" {
		if repo, e := readRepo(repo_path); e == nil {
			Repos[repo_path] = repo
			fmt.Printf("Adding '%s'", repo_path)
		} else {
			fmt.Printf("Error adding '%s' (%s)", repo_path, e)
		}
		return filepath.SkipDir
	}
	return nil
}

func readRepo(repo_path string) (*Repo, error) {
	repo := new(Repo)

	if r, err := gogit.OpenRepository(repo_path); err == nil {
		var ref *gogit.Reference
		var c *gogit.Commit
		var e error

		ref, e = r.LookupReference("HEAD")
		if e == nil {
			c, e = r.LookupCommit(ref.Oid)
		} else {
			fmt.Printf("Error reading repo '%s' (%s)\n",
				repo_path, e)
		}
		if e == nil {
			repo.Author = c.Author
			repo.CommitMessage = c.CommitMessage
		}
	}
	return repo, nil
}

func ReadRepos(base string) {
	Repos = make(map[string]*Repo)
	filepath.Walk(base, getRepos)
}

func main() {
	ReadRepos(os.Args[1])
	for repo_path, repo_data := range Repos {
		fmt.Printf("%s - %s\n", repo_path, repo_data.Author.Name)
	}
}
