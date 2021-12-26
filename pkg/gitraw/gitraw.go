package gitraw

import (
	"context"
	"fmt"
	"os"
)

func ListContents(ctx context.Context, repository, branch string, paths []string) {
	repoObject, err := RepoExists(ctx, repository, branch)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(paths) == 0 {
		paths = append(paths, "/")
	}
	for _, path := range paths {
		_, dirContent, _, err := client.Repositories.GetContents(
			ctx, repoObject.owner, repoObject.repo, path,
			GetRepositoryContentGetOptions(repoObject.branch),
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(dirContent) > 0 {
			for _, con := range dirContent {
				if *con.Type == "dir" {
					*con.Name += "/"
				}
				fmt.Println(*con.Name)
			}
		}
	}
}

func DownloadContents(ctx context.Context, repository, branch, destination string, paths []string) {
	repoObject, err := RepoExists(ctx, repository, branch)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(paths) == 0 {
		paths = append(paths, "/")
	}
	err = DownloadDir(repoObject, destination, paths, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
