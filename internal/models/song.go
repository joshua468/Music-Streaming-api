package models

import (
	"time"
)

type Song struct {
	Title        string    `json:"title"`
	Artist       string    `json:"artist"`
	Genre        string    `json:"genre"`
	DateReleased time.Time `json:"dateReleased"`
	Listeners    float64   `json:"listeners"`
	Followers    string    `json:"followers"`
	Download     string    `json:"download"`
	Listen       string    `json:"listen"`
}
