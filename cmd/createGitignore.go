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
var excludeFile string
var createGitignoreCmd = &cobra.Command{
	Use:   "createGitignore",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Aliases: []string{"c-gi"},
	Run: func(cmd *cobra.Command, args []string) {

		excludeFileContents, err := os.Open(excludeFile)

		if err != nil {
			fmt.Println("An error has occurred with reading" + excludeFile)
			return
		}

		defer excludeFileContents.Close()

		gitFile, err := os.Create(".gitignore")

		if err != nil {
			fmt.Println("An error occurred while creating .gitignore file")
			return
		}

		defer gitFile.Close()

		_, err = io.Copy(gitFile, excludeFileContents)
		if err != nil {
			fmt.Println("An error occurred when copying file contents to new file")
		}

		fmt.Println("Created a " + gitFile.Name() + " file")

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
