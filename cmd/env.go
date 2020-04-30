package cmd

import "os"

func config() {
	_ = os.Setenv("YAML_PATH", "/Users/mohankrishnareddybade/go/src/github.com/tools/test.yaml")
	_ = os.Setenv("JSON_PATH", "/Users/mohankrishnareddybade/go/src/github.com/tools/test.json")
}
