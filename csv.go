package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// Reads from a csv file and returns a slice of UserEntities
func readFromFile(filePath string) []UserEntity {
	csvFile, err := os.Open(filePath)

	if err != nil {
		csvFile, _ = os.Create(filePath)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var userEntities []UserEntity

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		checkError("Failed to read lines in file\n", err)
		i, _ := strconv.ParseInt(line[1], 10, 64)
		userEntities = append(userEntities, UserEntity{
			ScreenName:        line[0],
			FollowedTimestamp: i,
		})
	}

	return userEntities
}
