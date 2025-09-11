package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
        Use:   "parse",
        Short: "Convert dialog file.",
        RunE: execParseCmd,
    }

func init() {
	parseCmd.Flags().String("format", "", "Define file format. [json/yaml]")
	parseCmd.Flags().StringP("file", "f", "", "File name to parse.")
	RootCmd.AddCommand(parseCmd)
}

func execParseCmd (cmd *cobra.Command, args []string) error {
	format, err := cmd.Flags().GetString("format")
	if err != nil {
		return err
	}

	file_name, err := cmd.Flags().GetString("file") 
	if err != nil {
		return  err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	file_full_path := filepath.Join(cwd, file_name)
	file_absolute_path, err := filepath.Abs(file_full_path)
	if err != nil {
		return err
	}

	fmt.Println("File name: "+ file_name)
	fmt.Println("File absolute path: "+ file_absolute_path)
	fmt.Println("File extension: "+ format)
	
	return nil
}