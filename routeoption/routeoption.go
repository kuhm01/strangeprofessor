package routeoption

import "fmt"

func RouteOption(args []string) {
	arg := args[1]
	switch arg {
	case "-h", "help":
		printHelpOption()
	case "-c", "start":
		if len(args) == 2 {
			fmt.Printf("Not input file. professorc <command> [argument]\n")
		} else {
			Interpreter(args[2])
		}
	default:
		fmt.Printf("Professorlang. Wrong Option\n")
	}
}

func printHelpOption() {
	fmt.Printf("Professorlang is a tool for managing Professorlang source code.\n\nUsage:\n\n")
	fmt.Printf("         professorc <command> [argument]\n\n")
	fmt.Printf("The commands are:\n\n")
	fmt.Printf("         -h help       print Information\n")
	fmt.Printf("         -c start      Start Interpreting\n")
	fmt.Printf("\n\n")
}
