package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type counter struct {
	dirCount  int
	fileCount int
}

// treeCmd represents the tree command
var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {

		counter := &counter{}
		if len(args) == 0 {
			args = []string{"."}
		}
		for _, arg := range args {
			err := tree(arg, "", counter)
			if err != nil {
				log.Printf("tree %s: %v\n", arg, err)
			}
		}
		fmt.Printf("\n%d directories, %d files\n", counter.dirCount, counter.fileCount)
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)
}

func tree(root, indent string, counter2 *counter) error {
	fi, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("could not stat %s: %v", root, err)
	}

	fmt.Println(fi.Name())
	if !fi.IsDir() {
		counter2.fileCount++
		return nil
	}

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("could not read dir %s: %v", root, err)
	} else {
		counter2.dirCount++
	}

	var names []string
	for _, fi := range fis {
		if fi.Name()[0] != '.' {
			names = append(names, fi.Name())
		}
	}

	for i, name := range names {
		add := "│   "
		if i == len(names)-1 {
			fmt.Printf(indent + "└── ")
			add = "    "
		} else {
			fmt.Printf(indent + "├── ")
		}

		if err := tree(filepath.Join(root, name), indent+add, counter2); err != nil {
			return err
		}
	}

	return nil
}
