package main

// FollowEntry holds data for a user that we followed
type FollowEntry struct {
	ScreenName string `json:"screenName"`
	Timestamp  int64  `json:"timestamp"`
}
