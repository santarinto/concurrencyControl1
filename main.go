package main

import (
	"concurrencyControl1/cmd"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "scan":
		cmd.RunScan(os.Args[2:])
	case "help", "-h", "--help":
		printHelp()
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Network Scanner Tool")
	fmt.Println("====================")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  concurrencyControl1 <command> [options]")
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println("  scan        Scan local network for active IPv4 hosts")
	fmt.Println("  help        Show this help message")
	fmt.Println()
	fmt.Println("For more information about a command, use: concurrencyControl1 <command> --help")
}
