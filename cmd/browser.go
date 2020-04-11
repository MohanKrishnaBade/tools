package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var browserCmd = &cobra.Command{
	Use:   "browser",
	Short: "opens any URL in your default browser",
	Run: func(cmd *cobra.Command, args []string) {
		openBrowser(args[0])
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

func openBrowser(url string) {

	// append the common front part like web protocol stuff.
	if !strings.Contains(url, http) {
		url = http + url
	} else if !strings.Contains(url, https) {
		url = https + url
	}
	// append the common tail part.
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
		log.Fatal(err)
	}
}
