/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/bct/utils"
	"github.com/spf13/cobra"
)

// createCustomDomainCmd represents the createCustomDomain command
var name = ""
var useDDD = false
var createNewDomainCmd = &cobra.Command{
	Use:   "createNewDomain",
	Short: "Create a new domain with a custom name",
	Long: `Use this function to create a new custom domain with a given name. 
	This domain can autogenerate a chosen design philosophy or architecture as well.
	
	The current supported optional design patterns/architectures are:

	 - DDD
	
	`,
	Aliases: []string{"c-nd"},
	Run: func(cmd *cobra.Command, args []string) {

		if useDDD == false {

			utils.CreateDir(name)
			fmt.Println("Created simple custom domain:")

		} else {

			utils.CreateDomain(name)
			utils.CreateApplication(name)
			utils.CreateInfrastructure(name)
			utils.CreateEntrypoints(name)

		}
	},
}

func init() {
	rootCmd.AddCommand(createNewDomainCmd)
	createNewDomainCmd.Flags().StringVarP(
		&name,
		"name",
		"n",
		"newDomain",
		"Add a name for your new domain",
	)
	createNewDomainCmd.Flags().BoolVarP(
		&useDDD,
		"useDDD",
		"d",
		true,
		"Set the autogenerate for the domain to DDD style",
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCustomDomainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCustomDomainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
