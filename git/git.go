package git

import (
	log "github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// CreateBranch create a new branch on the current HEAD
func CreateBranch(repoPath string, branchName string) {
	branch := createReference(repoPath, branchName, "heads")
	repo := getRepo(repoPath)
	w, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
	}
	err = w.Checkout(&git.CheckoutOptions{
		Branch: branch.Name(),
	})
	if err != nil {
		log.Fatal(err)
	}
}

// CreateTag create a new tag on the current HEAD
func CreateTag(repoPath string, tagName string) {
	createReference(repoPath, tagName, "tags")
}

func createReference(repoPath string, refName string, refType string) *plumbing.Reference {
	repo := getRepo(repoPath)
	headRef, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}
	refStr := plumbing.ReferenceName("refs/" + refType + "/" + refName)
	ref := plumbing.NewHashReference(refStr, headRef.Hash())
	err = repo.Storer.SetReference(ref)
	if err != nil {
		log.Fatal(err)
	}
	return ref
}

func getRepo(repoPath string) *git.Repository {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}
