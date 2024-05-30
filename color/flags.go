package color

import (
	"errors"
	"flag"
	"strings"
)

// ColorFlag function parses command-line arguments to extract color options, text to be colored, and banner file name.
func ColorFlag() (string, string, string, string, error) {
	// Define flag for color option
	var colorFlag string
	flag.StringVar(&colorFlag, "color", "", "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"\n")
	flag.Parse()

	// Initialize variables for colorized letters, input text, and banner file
	colorizeLetters := ""
	inputText := ""
	bannerFile := ""

	// Determine the number of arguments and parse accordingly
	if len(flag.Args()) == 1 {
		// One argument: input text
		inputText = flag.Arg(0)
	} else if len(flag.Args()) == 2 {
		// Two arguments: colorize letters and input text or bannerfile
		if flag.Arg(1) != "thinkertoy" && flag.Arg(1) != "standard" && flag.Arg(1) != "shadow" {
			inputText = flag.Arg(1)
		} else {
			bannerFile = flag.Arg(1)
			inputText = flag.Arg(0)
		}
		colorizeLetters = flag.Arg(0)
	} else if len(flag.Args()) == 3 {
		// Three arguments: colirize letters, input text and bannerfile
		inputText = flag.Arg(1)
		colorizeLetters = flag.Arg(0)
		bannerFile = flag.Arg(2)
	} else {
		// Invalid number of arguments
		return "", "", "", "", errors.New("usage: go run . [OPTION] [STRING]\nex: go run . --color=<color> <letters to be colored> \"something\"")
	}

	// Convert color flag and banner file name to lowercase for consistency
	return strings.ToLower(colorFlag), colorizeLetters, inputText, strings.ToLower(bannerFile), nil
}
