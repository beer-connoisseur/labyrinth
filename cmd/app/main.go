package main

import (
	"fmt"
	"os"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/infrastructure"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use --help")
		return
	}

	switch os.Args[1] {
	case "-h", "--help":
		infrastructure.PrintHelp()
		return
	case "-V", "--version":
		fmt.Println("maze-app version", version)
		return
	case "generate":
		infrastructure.HandleGenerate(os.Args[2:])
	case "solve":
		infrastructure.HandleSolve(os.Args[2:])
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		fmt.Println("Use --help")
	}
}
