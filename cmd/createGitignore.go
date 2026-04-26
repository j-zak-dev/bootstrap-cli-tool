/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// createGitignoreCmd represents the createGitignore command
var excludeFile = ""
var createGitignoreCmd = &cobra.Command{
	Use:     "createGitignore",
	Short:   "Create a new .gitignore file",
	Long:    `Create a new .gitignore file with the option to pass in a file path with a pre-configured template for the .gitignore file.`,
	Aliases: []string{"c-gi"},
	Run: func(cmd *cobra.Command, args []string) {

		gitFile, err := os.Create(".gitignore")

		if err != nil {
			fmt.Println("An error occurred while creating .gitignore file")
			return
		}

		if len(excludeFile) > 0 {
			excludeFileContents, err := os.Open(excludeFile)

			if err != nil {
				fmt.Println("An error has occurred with reading" + excludeFile)
				return
			}

			_, err = io.Copy(gitFile, excludeFileContents)
			if err != nil {
				fmt.Println("An error occurred when copying file contents to new file")
			}

			defer excludeFileContents.Close()
			defer gitFile.Close()
			fmt.Println("Created a " + gitFile.Name() + " file")
		}
	},
}

func init() {
	rootCmd.AddCommand(createGitignoreCmd)
	createGitignoreCmd.Flags().StringVarP(
		&excludeFile,
		"file", // long flag: --file
		"f",    // short flag: -f
		"",     // default value
		"input file path",
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createGitignoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createGitignoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
