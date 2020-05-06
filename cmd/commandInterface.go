package cmd

type Command interface {
	Run(args []string) error
}
