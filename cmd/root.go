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
	"ccgen/ai"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var aiService *ai.Service

var rootCmd = &cobra.Command{
	Use:   "ccgen",
	Short: "AI-powered Conventional Commit generator",
	Long: `ccgen is a lightweight CLI tool that leverages AI to automatically generate
Conventional Commit messages and pull requests descriptions from your git diffs`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString("api_key")
		fmt.Println("Using API key:", apiKey)
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func Execute() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initAiConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ccgen.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ccgen")
	}

	viper.AutomaticEnv()

	fmt.Println(viper.ConfigFileUsed())

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "No config file found:", err)
		os.Exit(1)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())

}

func initAiConfig() {
	service, err := ai.New(
		viper.GetString("api_key"),
		viper.GetString("model"),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error with GeminiAPI:", err)
	}

	aiService = service
}
