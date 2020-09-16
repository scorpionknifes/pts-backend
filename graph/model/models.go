package model

import (
	"time"
)

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

type Turn struct {
	ID        int `json:"id"`
	UserID    int
	User      *User `json:"user"`
	StoryID   int
	Story     *Story    `json:"story"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Count struct {
	Count  int
	People int
}
