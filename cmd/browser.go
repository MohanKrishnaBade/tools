package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

type browser struct {
}

var browserCmd = &cobra.Command{
	Use:   "browser",
	Short: "open an URL in the default browser",
	Run: func(cmd *cobra.Command, args []string) {
		ThrowIf(browser{}.Run(args))
	},
}

func init() {
	rootCmd.AddCommand(browserCmd)
}

const (
	http  = "http://www."
	https = "https://www."
	com   = ".com"
)

func (b browser) Run(args []string) error {
	url := args[0]

	if !strings.Contains(url, http) {
		url = http + url
	} else if !strings.Contains(url, https) {
		url = https + url
	}

	if !strings.Contains(url, com) {
		url += com
	}
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		return err
	}

	return nil
}
