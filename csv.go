package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// Reads from a csv file and returns a slice of FollowEntries
func readFromFile(filePath string) []FollowEntry {
	csvFile, err := os.Open(filePath)

	if err != nil {
		csvFile, _ = os.Create(filePath)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var followEntries []FollowEntry

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		checkError("Failed to read lines in file\n", err)
		i, _ := strconv.ParseInt(line[1], 10, 64)
		followEntries = append(followEntries, FollowEntry{
			ScreenName: line[0],
			Timestamp:  i,
		})
	}

	return followEntries
}
