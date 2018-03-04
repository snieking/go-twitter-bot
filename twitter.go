// This file is supposed to handle all logic when it comes to communicating
// with the Twitter API. It creates the clients, and implements all needed
// operations towards the API.
package main

import (
	"log"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var client twitter.Client
var limitTracker LimitTracker

// Creates and returns the twitter client that will be used to perform
// actions towards the Twitter API.
func createConnection(twitterConf TwitterAccess) {
	config := oauth1.NewConfig(twitterConf.ConsumerKey, twitterConf.ConsumerSecret)
	token := oauth1.NewToken(twitterConf.AccessToken, twitterConf.AccessSecret)

	// http.Client will automatically authorize requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client = *twitter.NewClient(httpClient)
}

// Follows a provided user.
func follow(user string) {
	preventReachingLimit()
	_, _, err := client.Friendships.Create(&twitter.FriendshipCreateParams{
		ScreenName: user,
		Follow:     newTrue(),
	})
	logError(err)
}

// Unfollows a provided user.
func unfollow(user string) {
	preventReachingLimit()
	_, _, err := client.Friendships.Destroy(&twitter.FriendshipDestroyParams{
		ScreenName: user,
	})
	checkError("Failed to unfollow\n", err)
}

// List all the followers of a provided user.
func listFollows(user string) []string {
	users := []string{}
	preventReachingLimit()
	friends, _, error := client.Friends.List(&twitter.FriendListParams{ScreenName: user, Count: 1000})
	checkError("Failed to fetch friends\n", error)
	for _, element := range friends.Users {
		users = append(users, element.ScreenName)
	}
	return users
}

// Search for tweets based on a provided topic and returns as many users
// who wrote tweets as it can find based on the provided topic and limit
func searchTweets(value string, limit int) []UserEntity {
	preventReachingLimit()
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: value,
		Count: limit,
	})

	checkError("Failed to search for tweets\n", err)

	var users []UserEntity
	for _, element := range search.Statuses {
		if element.Lang == "en" {
			users = append(users, UserEntity{
				ScreenName: element.User.ScreenName,
				UserID:     element.User.ID,
			})
		}
	}

	return users
}

// Gets information of who a user follows.
// Returns a map where the key is the ID for easier and quicker lookup.
func getMapOfFollowedUsers(user string) map[int64]bool {
	preventReachingLimit()
	friends, _, err := client.Friends.IDs(&twitter.FriendIDParams{ScreenName: user})
	checkError("Failed to fetch followed users\n", err)
	m := make(map[int64]bool)
	for _, element := range friends.IDs {
		m[element] = true
	}
	return m
}

// Twitter has a limitation where you cannot perform more than 15 operations per limit window.
// A limit window is started when you perform your first operation.
// This function tracks the number of operations that have been performed in the active window
// and if we go over it, it will sleep until the window is over.
func preventReachingLimit() {
	if limitTracker.WindowStarted.IsZero() {
		limitTracker = LimitTracker{Operations: 0, WindowStarted: time.Now()}
	} else if limitTracker.Operations > opsBeforeSleep {
		windowStartInNano := limitTracker.WindowStarted.Nanosecond()
		nowInNano := time.Now().Nanosecond()
		nanosSinceStarted := nowInNano - windowStartInNano
		shouldSleepForNanos := (time.Duration(15) * time.Minute) - (time.Duration(nanosSinceStarted) * time.Nanosecond)
		log.Printf("Sleeping for about %d minutes", shouldSleepForNanos/60000000000.0)
		time.Sleep(time.Duration(shouldSleepForNanos) * time.Nanosecond)
	}

	if limitTracker.WindowStarted.Add(15 * time.Minute).Before(time.Now()) {
		limitTracker = LimitTracker{Operations: 0, WindowStarted: time.Now()}
	} else {
		limitTracker.Operations++
	}

}

// Config holds configuration from the user
type Config struct {
	TwitterName   string
	Interests     []string
	TwitterAccess TwitterAccess
}

// TwitterAccess holds the keys, and secrets for the Twitter API
type TwitterAccess struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

// UserEntity holds data for a user that we followed
type UserEntity struct {
	ScreenName        string `json:"screenName"`
	UserID            int64  `json:"userID"`
	FollowedTimestamp int64  `json:"followedTimestamp"`
}

type LimitTracker struct {
	Operations    int
	WindowStarted time.Time
}
