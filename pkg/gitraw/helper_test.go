package gitraw

import (
	"context"
	"testing"

	"github.com/google/go-github/github"
)

var (
	repository = "yadneshk/k8s-cluster-deploy"
	branch     = "main"
)

func TestAuthenticateGithub(t *testing.T) {
	var want error

	got := AuthenticateGithub(context.Background())
	want = nil

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}

func TestRepoExists(t *testing.T) {

	got, _ := RepoExists(context.Background(), repository, branch)
	want := Repo{
		owner:  "yadneshk",
		repo:   "k8s-cluster-deploy",
		branch: "main",
	}
	if *got != want {
		t.Errorf("got %q wanted %q", *got, want)
	}

}

func TestGetRepositoryContentGetOptions(t *testing.T) {
	got := GetRepositoryContentGetOptions("main")
	want := github.RepositoryContentGetOptions{
		Ref: "main",
	}
	if *got != want {
		t.Errorf("got %q wanted %q", *got, want)
	}
}
