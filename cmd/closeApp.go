package cmd

import (
	"os/exec"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
)

type closeApp struct {
}

// closeAppCmd represents the closeApp command
var closeAppCmd = &cobra.Command{
	Use:   "closeApp",
	Short: "close any open app.",
	Run: func(cmd *cobra.Command, args []string) {
		ThrowIf(closeApp{}.Run(args))
	},
}

func init() {
	rootCmd.AddCommand(closeAppCmd)
}

func (c closeApp) Run(args []string) error {

	if len(args) > 0 {
		jetBrainsIdes := [3]string{"phpstorm", "goland", "webstorm"}
		if !itemExists(jetBrainsIdes, args[0]) {
			args[0] = strings.Title(args[0])
		}
		out, err := exec.Command("/bin/sh", "-c", "ps -ax  | grep "+args[0]).Output()
		if err != nil {
			return err
		}

		processes := strings.Split(string(out[:]), "??")
		if len(processes) > 0 {
			err := exec.Command("/bin/sh", "-c", "kill "+strings.TrimSpace(processes[0])).Start()
			if err != nil {
				return err
			}
		}
	}

	return nil
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
