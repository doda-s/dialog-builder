package cmd

import (
	"dodas/dialogbuilder/internal/core"
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
	cwd := core.Directory.Cwd

	fmt.Println("Cwd: " + cwd)

	fileRelativePath, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}

	fmt.Println("File relative path: " + fileRelativePath)

	var fileAbsolutePath string

	if strings.Contains(cwd, fileRelativePath) {
		fileFullPath := filepath.Join(cwd, fileRelativePath)
		fileAbsolutePath, err = filepath.Abs(fileFullPath)
		if err != nil {
			return err
		}
		parseToJson(fileAbsolutePath)
		
		return nil
	}

	fileAbsolutePath, err = filepath.Abs(fileRelativePath)

	fmt.Println("Absolute path: " + fileAbsolutePath)
	
	if err != nil {
		return err
	}

	parseToJson(fileAbsolutePath)

	return nil
}

func parseToJson(absolutePath string) {
	data, err := os.ReadFile(absolutePath)
	if err != nil {
		panic(err)
	}

	var scopeFile dialog.Scope
	err = yaml.Unmarshal(data, &scopeFile)
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(scopeFile, "", "\t")
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
