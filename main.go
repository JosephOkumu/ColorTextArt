package main

import (
	"fmt"
	"os"

	"asciiart/color"
)

// main function reads a banner file, creates a map of ASCII art, validates user input,
// and prints the corresponding ASCII art to the console.
func main() {
	// Check if color flag is not provided correctly, i.e. provided without equal sign.
	args := os.Args
	for _, v := range args {
		if v == "--color" || v == "-color" {
			fmt.Print("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"\n")
			return
		}
	}

	// Get the flag values for color, letters to colorize, input text and banner file name. Handle possible errors.
	colorFlag, colorizeLetters, inputText, banner, err := color.ColorFlag()
	check(err)

	// Get ANSI format string to colorize ASCII-art in the terminal.
	colorCode, err := color.SetColor(colorFlag)
	check(err)

	// Read banner file
	if banner == "" {
		banner = "standard"
	}
	bannerFile, err := color.ReadBannerFile("./banners/" + banner + ".txt")
	check(err)

	// Create map of ASCII art
	ASCIIArtMap, err := color.MapCreator(bannerFile)
	check(err)

	// Get ASCII art corresponding to input text
	asciiArt, err := color.ArtRetriever(inputText, colorCode, colorizeLetters, ASCIIArtMap)
	check(err)

	// Print the ASCII art to the terminal
	fmt.Print(asciiArt)
}

// Handle errors
func check(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", e)
		os.Exit(1)
	}
}
