/*
Copyright © 2024 Christopher Wagner
*/
package cmd

import (
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a todo appliation",
	Long: `Tri is a todo application that allows you to create, read, update, and delete tasks.
	It's designed to be simple and easy to use so you can get more done.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var datafile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	home, err := homedir.Dir()
	if err != nil {
		log.Println("Could not find home directory")
	}

	rootCmd.PersistentFlags().StringVar(&datafile, "datafile", home+string(os.PathSeparator)+".tridos.json", "the file to store todo items")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


