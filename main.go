package main

import (
	"ConvertDate/DateConverter"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {

	// Validate the arguments are correct, display the usage if not
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "Usage: ConvertDate inputFile inputFormat outputFile outputFormat\nDate format components:")
		var tags []string
		for tag, _ := range DateConverter.FormatMap {
			tags = append(tags, tag)
		}
		sort.Strings(tags)

		for _, tag := range tags {
			fmt.Fprintf(os.Stderr, "   %s - %s\n", tag, DateConverter.FormatMap[tag].Description)
		}
		os.Exit(1)
	}

	inputFile, inputFormat, outputfile, outputFormat := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	// Create a new converter
	dc, err := DateConverter.NewDateConverter(inputFormat, outputFormat)

	if err != nil {
		log.Fatalf("[FATAL] Failed to initialize DateConverter: %v\n", err)
	}

	// Process the file
	if err = dc.ConvertFile(inputFile, outputfile); err != nil {
		log.Fatalf(`[FATAL] Failed to convert dates for "%s" to "%s": %v`, inputFile, outputfile, err)
	}

}
