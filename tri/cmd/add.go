/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Cwagne17/tri/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: addRun,
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority of the item")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(datafile)
	if err != nil {
		cmd.PrintErrf("%v\n", err)
	}

	for _, arg := range args {
		item := todo.Item{Text: arg}
		item.SetPriority(priority)

		items = append(items, item)
	}
	
	err = todo.SaveItems(datafile, items)
	if err != nil {
		cmd.PrintErrf("%v\n", err)
	}
}
