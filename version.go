package main

var Version = "dev"

//Dood, what
var cmdVersion = &Command{
	Run:   runVersion,
	Usage: "version",
	Short: "Display current version",
	Long: `
Display current version

Examples:

  force version
`,
}

func init() {
}

func runVersion(cmd *Command, args []string) {
	ConsolePrintln(Version)
}
