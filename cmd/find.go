package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tools/Sort"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var printCollection []string

// findCmd represents the path command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "gets the file or folder path that you are looking for",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) > 0 {
			err := path(".", args[0])

			if len(printCollection) == 0 {
				fmt.Printf("no match found with the given string :: %s\n", args[0])
			} else {
				Sort.BSort{}.SortByLength(printCollection)
				for _, v := range printCollection {
					fmt.Println(v)
				}
				fmt.Printf("\n%d matches found with a given string::%s\n", len(printCollection), args[0])
			}

			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}

func path(root, fileName string) error {
	fi, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("could not stat %s: %v", root, err)
	}
	if !fi.IsDir() {
		//printIfMatch(root, fileName)
		return nil
	}

	files, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("could not readDir %s: %v", root, err)
	}

	var fileCollection []string
	for _, fi := range files {
		if fi.Name()[0] != '.' {
			fileCollection = append(fileCollection, fi.Name())
		}
	}

	for _, file := range fileCollection {

		printIfMatch(root, file, fileName)
		if err := path(filepath.Join(root, file), fileName); err != nil {
		}
	}

	return nil
}

func printIfMatch(root, str, str1 string) {
	if strings.Contains(strings.TrimSuffix(str, filepath.Ext(str)), str1) {
		printCollection = append(printCollection, "./"+filepath.Join(root, str))
	}
}
