package main

import (
	"fmt"
	"github.com/eurozulu/4mat/parser"
	"github.com/eurozulu/commandgo"
	"log"
	"os"
)

func main() {
	commandgo.AddFlag(&parser.FlagRecursive, "recursive", "r")

	cmd := commandgo.Commands{
		"yaml": ToYaml,
		"json": JsonCommand.ToJson,

		"": showHelp,
	}
	if err := cmd.Run(os.Args...); err != nil {
		log.Fatalln(err)
	}
}

func showHelp() {
	fmt.Printf("%s <format> [<filepath>[ <filepath> ...]]\n", os.Args[0])
	fmt.Println("\t<format>\t\tThe output format. Can be:")
	fmt.Println("\t\tjson")
	fmt.Println("\t\tyaml")
	fmt.Println("\t\tcsv")
	fmt.Println("\t\txml")
	fmt.Println("\t<filepath>\t\tpath to a text file to ingest")
	fmt.Println("\t\tfilepath can be one or more, space seperated, file paths, each read sequentially")
	fmt.Println("\t\ta dash '-' as a filepath indicates to read from the standard input.")
	fmt.Println("\t\tIf no filepaths are given, the dash is assumed, reading (or blocking) on the stdin")
	fmt.Println()
	os.Exit(-1)
}
