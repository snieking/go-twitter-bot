// The purpose of this application is to follow new users in hope that they will
// return the favor and follow you back. After a set of time the program will unfollow
// the user. And hopefully, the user that followed you forgets to unfollow you.
package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var clean, unfollowAll bool
var followHours, sleepTime, opsBeforeSleep int
var config Config

// Parses flags and reads the configuration file.
func init() {
	flag.BoolVar(&clean, "clean", false, "Cleans all previous follows from the bot")
	flag.BoolVar(&unfollowAll, "unfollowAll", false, "Unfollows all your friends")
	flag.IntVar(&followHours, "followHours", 6, "How many hours to follow users")
	flag.IntVar(&sleepTime, "sleepTime", 15, "Time in minutes to sleep between each circle")
	flag.IntVar(&opsBeforeSleep, "opsBeforeSleep", 10, "Number of operations before sleeping")
	flag.Parse()

	filePath := getPath("config.json")

	file, err1 := ioutil.ReadFile(filePath)
	if err1 != nil {
		checkError("Error while reading file\n", err1)
		os.Exit(1)
	}

	err2 := json.Unmarshal(file, &config)
	if err2 != nil {
		log.Fatal("error:", err2)
		os.Exit(1)
	}
}

// Starts the bot.
func main() {
	startBot()
}
