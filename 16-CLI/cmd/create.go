/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/4lexRossi/pos-go/16-CLI/internal/database"
	"github.com/spf13/cobra"
)

func NewCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  `A longer description that spans multiple lines and likely contains examples`,
		RunE:  runCreate(categoryDb),
	}
}

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := NewCreateCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Category name")
	createCmd.Flags().StringP("description", "d", "", "Category description")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
