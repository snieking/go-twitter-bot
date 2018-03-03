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
	defer csvFile.Close()

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

// Writes the list of user entities to file in order to keep track of
// who to later unfollow and at what time.
func writeListOfFollowsToFile(userEntities []UserEntity) {
	file, err := os.Create("follows.csv")
	checkError("Cannot create file\n", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range userEntities {
		timestamp := strconv.FormatInt(value.FollowedTimestamp, 10)
		strWrite := []string{value.ScreenName, timestamp}
		err := writer.Write(strWrite)
		writer.Flush()
		checkError("Cannot write to file", err)
	}
}
