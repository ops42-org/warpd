package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(githubCmd)
}

var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Github helper commands",
	Long:  `Various helpers for github actions`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println(version)
	// },
}
