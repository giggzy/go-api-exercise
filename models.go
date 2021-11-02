package main

import "sync"

type Services struct {
	Lock     sync.RWMutex
	Services []Service `json:"services"`
}
type Service struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	VersionCount int       `json:"versionCount"`
	URL          string    `json:"url"`
	Versions     []Version `json:"versions"`
}

type Version struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// not used yet, intention is to pass metadata on pagination to frontend
type Meta struct {
	Page     int
	PageSize int
	//total int
}
