/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	gitraw "gitraw/pkg"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get a file from a repository",
		Long: `Get a file from a repository
	For example:
	$ gitraw get -r <username>/<repository_name> -b <branch> <filepath>
	`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			gitraw.DownloadContents(ctx, repositoryFlag, branchFlag, destDir, args)
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().StringVarP(&repositoryFlag, "repository", "r", "", "<username>/<repository_name>")
	getCmd.MarkPersistentFlagRequired("repository")

	getCmd.PersistentFlags().StringVarP(&destDir, "output-dir", "o", "", "Output directory for downloaded content")
	getCmd.MarkPersistentFlagRequired("output-dir")

	getCmd.Flags().StringVarP(&branchFlag, "branch", "b", "master", "branch")
}
