package cmd

import (
	"dodas/dialogbuilder/internal/dialog"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Convert dialog file.",
	RunE:  execParseCmd,
}

func init() {
	parseCmd.Flags().String("format", "", "Define file format. [json/yaml]")
	parseCmd.Flags().StringP("file", "f", "", "File name to parse.")
	RootCmd.AddCommand(parseCmd)
}

func execParseCmd(cmd *cobra.Command, args []string) error {
	format, err := cmd.Flags().GetString("format")
	if err != nil {
		return err
	}

	fileName, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fileFullPath := filepath.Join(cwd, fileName)
	fileAbsolutePath, err := filepath.Abs(fileFullPath)
	if err != nil {
		return err
	}

	fmt.Println("File name: " + fileName)
	fmt.Println("File absolute path: " + fileAbsolutePath)
	fmt.Println("File extension: " + format)

	parseToJson(fileAbsolutePath)

	return nil
}

func parseToJson(absolutePath string) {
	data, err := os.ReadFile(absolutePath)
	if err != nil {
		panic(err)
	}

	var scopeFile dialog.ScopeFile
	err = yaml.Unmarshal(data, &scopeFile)
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(scopeFile, "", " ")
	if err != nil {
		panic(err)
	}

	path := strings.Split(absolutePath, "/")
	fileName := path[len(path)-1]
	fileName = strings.Split(fileName, ".")[0]
	err = os.WriteFile(fileName+".json", jsonData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("File '"+strings.Split(fileName, ".")[0]+"' created.")
}
