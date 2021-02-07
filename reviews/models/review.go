package models

import "time"

// Review represent an anon review from some website
type Review struct {
	Id      int64
	Stars   int       // 1 - 5
	Comment string    // max 400 chars
	Date    time.Time // created at
}

// CreateReviewCMD command to create a new review
type CreateReviewCMD struct {
	Stars   int    `json:"stars"`
	Comment string `json:"comment"`
}
