package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type Connections struct {
	Connections []struct {
		Host     string `json:"host_name"`
		UserName string `json:"user_name"`
		DbName   string `json:"db_name"`
		Password string `json:"password"`
		Port     string `json:"port"`
	} `json:"connections"`
}

type dbConfig struct {
	Hostname string
	Username string
	Port     string
}

func (c *dbConfig) Parse(data []byte) error {
	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}
	return nil
}

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(args); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
	config()
}

func run(args []string) error {

	if len(args) == 1 {
		jsonFilePath, yamlFilePath := os.Getenv("JSON_PATH"), os.Getenv("YAML_PATH")
		fi, err := os.Stat(jsonFilePath)
		if err != nil {
			return err
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			return fmt.Errorf("we are not handling the directories for now")
		case mode.IsRegular():

			jsonFile, err := os.Open(jsonFilePath)
			if err != nil {
				return err
			}
			byteValue, _ := ioutil.ReadAll(jsonFile)
			var result Connections
			_ = json.Unmarshal(byteValue, &result)

			for _, db := range result.Connections {
				if strings.Contains(db.UserName, args[0]) {
					data, err := ioutil.ReadFile(yamlFilePath)
					if err != nil {
						return err
					}
					var config dbConfig
					if err := config.Parse(data); err != nil {
						return err
					}

					finalData := string(data)

					finalData = strings.Replace(finalData, config.Hostname, db.Host, -1)
					finalData = strings.Replace(finalData, config.Username, db.UserName, -1)
					finalData = strings.Replace(finalData, config.Port, db.Port, -1)

					fmt.Println(finalData)
					return ioutil.WriteFile(yamlFilePath, []byte(finalData), 0)
				}
			}
		}
		return nil
	}
	return fmt.Errorf("you should pass atleast one argumanet to process")
}
