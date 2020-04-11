package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// openAppCmd represents the openApp command
var openAppCmd = &cobra.Command{
	Use:   "openApp",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		execute(args)
	},
}

func init() {
	rootCmd.AddCommand(openAppCmd)
}

func execute(args []string) {

	argLen := len(args)
	if argLen < 1 {
		fmt.Printf("Couldn't process your request with %d arguments", argLen)
	} else if argLen == 1 {
		args = append(args, "./")
	}

	cmd := ""
	switch args[0] {
	case "ps":
		cmd = "/usr/local/bin/pstorm"
	case "gl":
		cmd = "/usr/local/bin/goland"
	default:
		fmt.Println("argument one doesn't match with our list")
		return
	}

	if err := exec.Command(cmd, args[1]).Start(); err != nil {
		fmt.Printf("%s", err)
	}
}
