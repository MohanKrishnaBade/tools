package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

type openApp struct {
}

// openAppCmd represents the openApp command
var openAppCmd = &cobra.Command{
	Use:   "openApp",
	Short: "open apps that are already installed in your local machine.",
	Run: func(cmd *cobra.Command, args []string) {
		ThrowIf(openApp{}.Run(args))
	},
}

func init() {
	rootCmd.AddCommand(openAppCmd)
}

func (o openApp) Run(args []string) error {

	argLen := len(args)
	if argLen < 1 {
		return fmt.Errorf("couldn't process your request with %d arguments", argLen)
	} else if argLen == 1 {
		args = append(args, "./")
	}

	var name, cmd string
	switch args[0] {
	case "ps":
		cmd = "/usr/local/bin/pstorm"
	case "gs":
		cmd = "/usr/local/bin/goland"
	case "ws":
		cmd = "/usr/local/bin/webstorm"
	case "chrome":
		cmd = "/Applications/Google Chrome.app"
		name = "open"
	case ".":
		name = "open"
		cmd = "."
	default:
		cmd = "/Applications/" + strings.Title(args[0]) + ".app"
		name = "open"
	}

	if name != "" {
		if err := exec.Command(name, cmd).Start(); err != nil {
			return fmt.Errorf("%s", err)
		}

	} else if err := exec.Command(cmd, args[1]).Start(); err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}
