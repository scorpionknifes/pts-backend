package model

import (
	"time"
)

// Story - Story for story table
type Story struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Count     int       `json:"count"`
	People    int       `json:"people"`
	Tags      string    `json:"tags"`
	Turns     []*Turn   `json:"turns"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Turn - Turn for turns table
type Turn struct {
	ID        int `json:"id"`
	UserID    int
	StoryID   int
	Story     *Story    `json:"story"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// User - User for users table
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Count - Custom count for number of users
type Count struct {
	Count  int
	People int
}
