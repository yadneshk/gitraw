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
	gitraw "gitraw/pkg/gitraw"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List files and directories of a repository",
		Long: `List files and directories of a repository.

For example:
$ gitraw list -r <username>/<repository_name> -b <branch>
`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			gitraw.ListContents(ctx, repositoryFlag, branchFlag, args)
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringVarP(&repositoryFlag, "repository", "r", "", "<username>/<repository_name>")
	listCmd.MarkPersistentFlagRequired("repository")

	listCmd.Flags().StringVarP(&branchFlag, "branch", "b", "master", "branch")
}
