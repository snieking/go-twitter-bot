// Handles all the business logic for the twitter bot.
package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

// Starts the bot, creating the twitter connection as well as performing
// all the actions until an fatal exception occurs.
func startBot() {
	createConnection(config.TwitterAccess)

	if unfollowAll {
		unfollowAllFromUserAndExit(config.TwitterName)
	}

	for {
		followEntries := readFromFile("follows.csv")

		// Unfollow all previous followed users
		if clean {
			cleanFollowListAndExit(followEntries)
		}

		unfollowOldUsers(followEntries)
		followEntries = followNewUsers(followEntries)
		writeListOfFollowsToFile(followEntries)

		log.Printf("Done with follow/unfollow, sleeping for %d minutes", sleepTime)
		time.Sleep(time.Duration(sleepTime) * time.Minute)
	}
}

// Cleans the provided list of follow entries, meaning that all of them
// will be unfollowed.
func cleanFollowListAndExit(followEntries []FollowEntry) {
	log.Printf("Unfollowing all followed users in list")
	for index, element := range followEntries {
		unfollow(element.ScreenName)
		if index != 0 && index%19 == 0 {
			log.Printf("Sleeping for 10 min to prevent spam")
			time.Sleep(15 * time.Minute)
		}
		log.Printf("[%d] Unfollowed: %s", index, element.ScreenName)
	}
	os.Exit(3)
}

// Unfollows previous users that we followed.
// Users are considered old if they are older than the configured hours.
func unfollowOldUsers(followEntries []FollowEntry) {
	log.Printf("Checking if anyone needs to be unfollowed")
	for index, element := range followEntries {
		if element.Timestamp < makeTimestampHoursBeforeNow(followHours) {
			unfollow(element.ScreenName)
			followEntries = remove(followEntries, index)
			log.Printf("[%d] Unfollowed: %s", index, element.ScreenName)
		} else {
			log.Printf("[%d] user %s isn't due for unfollow yet", index, element.ScreenName)
		}
	}
}

// Unfollows all users of a provided twitterName.
// Recursively calls itself until finished.
func unfollowAllFromUserAndExit(twitterName string) {
	users := listFollows(twitterName)
	if len(users) < 1 {
		os.Create("follows.csv")
		os.Exit(3)
	} else {
		for index, element := range users {
			unfollow(element)
			// If more than 20 friends were to be returned for some reason
			if index != 0 && index%25 == 0 {
				log.Printf("Sleeping for %d min to prevent spam", sleepTime)
				time.Sleep(time.Duration(sleepTime) * time.Minute)
			}
			log.Printf("[%d] Unfollowed: %s", index, element)
		}
		log.Printf("Sleeping for %d min to prevent spam", sleepTime)
		time.Sleep(time.Duration(sleepTime) * time.Minute)
		unfollowAllFromUserAndExit(twitterName)
	}
}

// Follows all the users in the provided list of follow entries.
func followNewUsers(followEntries []FollowEntry) []FollowEntry {
	log.Printf("\nSearching for new users to follow")
	users := searchTweets(randomElementFromSlice(config.Interests), 10)
	for index, element := range users {
		followEntries = append(followEntries, FollowEntry{
			ScreenName: element,
			Timestamp:  makeTimestamp(),
		})
		follow(element)
		log.Printf("[%d] followed: %s", index, element)
	}

	return followEntries
}

// Writes the list of follow entries to file in order to keep track of
// who to later unfollow and at what time.
func writeListOfFollowsToFile(followEntries []FollowEntry) {
	file, err := os.Create("follows.csv")
	checkError("Cannot create file\n", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range followEntries {
		timestamp := strconv.FormatInt(value.Timestamp, 10)
		strWrite := []string{value.ScreenName, timestamp}
		err := writer.Write(strWrite)
		writer.Flush()
		checkError("Cannot write to file", err)
	}
}
