package cmd

import (
	"log"
	"os/exec"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
)

// closeAppCmd represents the closeApp command
var closeAppCmd = &cobra.Command{
	Use:   "closeApp",
	Short: "it helps to close any app.",
	Run: func(cmd *cobra.Command, args []string) {
		closeApp(args)
	},
}

func init() {
	rootCmd.AddCommand(closeAppCmd)
}

func closeApp(args []string) {

	if len(args) > 0 {
		jetBrainsIdes := [3]string{"phpstorm", "goland", "webstorm"}
		if !itemExists(jetBrainsIdes, args[0]) {
			args[0] = strings.Title(args[0])
		}
		out, err := exec.Command("/bin/sh", "-c", "ps -ax  | grep "+args[0]).Output()
		if err != nil {
			log.Fatal(err)
		}

		processes := strings.Split(string(out[:]), "??")
		if len(processes) > 0 {
			err := exec.Command("/bin/sh", "-c", "kill "+strings.TrimSpace(processes[0])).Start()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
