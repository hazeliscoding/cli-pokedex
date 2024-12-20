package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/hazeliscoding/godex/subcommands"
)

var styleBold = lipgloss.NewStyle().Bold(true)
var styleItalic = lipgloss.NewStyle().Italic(true)

func main() {

	flag.Usage = func() {
		fmt.Println("Welcome! This tool displays data about a selected Pokémon in the terminal!")

		fmt.Println(styleBold.Render("\nUSAGE:"))
		fmt.Println("\t", "godex [flag]")
		fmt.Println("\t", "godex [pokemon name] [flag]")
		fmt.Println("\t", "----------")
		fmt.Println("\t", styleItalic.Render("Example:"), "godex bulbasaur")

		fmt.Println(styleBold.Render("\nGLOBAL FLAGS:"))
		fmt.Println("\t", "-h, --help", "\t", "Shows the help menu")
		fmt.Print("\n")

		fmt.Println(styleBold.Render("POKEMON NAME FLAGS:"))
		fmt.Println("\t", "Add a flag after declaring a Pokémon's name for more details!")
		fmt.Print("\t", "--types\n\n")
	}

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Please declare a Pokémon's name after the CLI name")
		fmt.Println("Run 'godex --help' for more details")
		os.Exit(1)
	}

	subcommands.PokemonCommand()
}
