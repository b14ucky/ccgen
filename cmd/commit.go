/*
Copyright © 2026 Dominik Meisner <meisnerd2003@gmail.com>

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
	"ccgen/git"
	"fmt"

	"github.com/spf13/cobra"
)

var giveDescription bool = false
var commitCmd = &cobra.Command{
	Use:   "commit [files...]",
	Short: "Generate Conventional Commit message from staged git diff using AI",
	Long: `Generate a Conventional Commit message based on currently staged changes.

The command reads the output of "git diff --staged" and sends it to the configured
AI model (Gemini API). Based on the diff, it generates a properly formatted
Conventional Commit title.

By default, only the commit title is generated.
If the --description (-d) flag is provided, a full commit message
(title + description) will be generated.

You can optionally pass file names as arguments to exclude them from
the generated diff.

Examples:

  ccgen commit
  ccgen commit -d
  ccgen commit file1.go file2.go`,
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var filesToExclude []string
		for _, f := range args {
			filesToExclude = append(filesToExclude, ":!"+f)
		}
		diff, err := git.GetCommitDiff(filesToExclude...)
		if err != nil {
			return err
		}
		result, _ := aiService.GetCommitMessage(diff, giveDescription)
		fmt.Println(result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().BoolVarP(&giveDescription, "description", "d", false, "if toggled commit description will be also generated")
}
