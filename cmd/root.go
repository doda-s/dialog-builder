package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Run: execRootCmd,
}

func execRootCmd (cmd *cobra.Command, args []string) {
	fmt.Println("No action required.")
}