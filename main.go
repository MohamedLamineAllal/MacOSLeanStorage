package main

import "github.com/mohamedlamineallal/MacosLeanStorage/cmd"

// main is the entry point of the MacosLeanStorage application.
// It delegates execution to the cobra root command.
func main() {
	cmd.Execute()
}
