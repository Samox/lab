package cmd

import (
	"github.com/spf13/cobra"
)

// mrCmd represents the mr command
var mergeRequestCmd = &cobra.Command{
	Use:   "merge-request [remote [branch]]",
	Short: mrCreateCmd.Short,
	Long:  mrCreateCmd.Long,
	Args:  mrCreateCmd.Args,
	Run:   runMRCreate,
}

func init() {
	mergeRequestCmd.Flags().AddFlagSet(mrCreateCmd.Flags())
	RootCmd.AddCommand(mergeRequestCmd)
}
