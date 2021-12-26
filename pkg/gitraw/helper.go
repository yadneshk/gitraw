package gitraw

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	githubToken = "AUTH_GITHUB"
	client      *github.Client
)

type Repo struct {
	owner, repo, branch string
}

func AuthenticateGithub(ctx context.Context) error {
	token := os.Getenv(githubToken)
	if token == "" {
		return errors.New("missing value to environment variable AUTH_GITHUB")
		// TODO: Add authentication section from README
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)

	// consider authentication sucessful if we get user details
	_, _, err := client.Users.Get(ctx, "")

	if err != nil {
		return err
	}
	return nil
}

// Verify whether the repository and branch exists or not
func RepoExists(ctx context.Context, repository, branch string) (*Repo, error) {
	// list all repositories for the authenticated user
	err := AuthenticateGithub(ctx)
	if err != nil {
		return nil, errors.New(fmt.Sprint(err))
	}
	if !strings.Contains(repository, "/") {
		return nil, errors.New("repository format should be <owner>/<repository>")
	}
	repoInfo := strings.Split(repository, "/")

	repoObject := Repo{
		owner:  repoInfo[0],
		repo:   repoInfo[1],
		branch: branch,
	}

	_, _, err = client.Repositories.Get(ctx, repoObject.owner, repoObject.repo)
	if err != nil {
		return nil, errors.New(fmt.Sprint(err))
	}

	_, _, err = client.Repositories.GetBranch(ctx, repoObject.owner, repoObject.repo, repoObject.branch)
	if err != nil {
		return nil, errors.New(fmt.Sprint(err))
	}
	return &repoObject, nil
}

func GetRepositoryContentGetOptions(ref string) *github.RepositoryContentGetOptions {
	return &github.RepositoryContentGetOptions{Ref: ref}
}

func DownloadFile(destination string, fileContent *github.RepositoryContent) error {
	destPath := filepath.Join(destination, *fileContent.Name)
	f, err := os.Create(destPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	content, err := fileContent.GetContent()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Fprint(f, content)
	fmt.Println(destPath)

	return nil
}

func DownloadDir(repoObject *Repo, destination string, paths []string, ctx context.Context) error {

	for _, path := range paths {
		fileContent, dirContent, _, err := client.Repositories.GetContents(
			ctx, repoObject.owner, repoObject.repo, path,
			GetRepositoryContentGetOptions(repoObject.branch),
		)
		if err != nil {
			return err
		}
		// unempty fileContent would mean it's a file
		if fileContent != nil {
			DownloadFile(destination, fileContent)
		}

		// unempty dirContent would mean it's a directory
		for _, con := range dirContent {
			dirPath := filepath.Join(destination, *con.Name)
			if *con.Type == "dir" {
				err := os.Mkdir(dirPath, 0755)
				if err != nil {
					return err
				}
				DownloadDir(repoObject, dirPath, []string{*con.Path}, ctx)
			} else if *con.Type == "file" {
				DownloadFile(destination, con)
			}

		}
	}

	return nil
}
