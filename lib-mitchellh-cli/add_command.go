package main

type AddCommand struct {
}

func (c *AddCommand) Synopsis() string {
	return "Add todo task to list"
}

func (c *AddCommand) Help() string {
	return "Usage: todo add [option]"
}

func (c *AddCommand) Run(args []string) int {
	// TODO
	return 0
}
