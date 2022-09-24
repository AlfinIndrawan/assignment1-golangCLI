package models

import "time"

type Data struct {
	Water int `json:"water"`
	Air   int `json:"wind"`
}

type DataStatus struct {
	WaterStatus string `json:"waterStatus"`
	AirStatus   string `json:"airStatus"`
}

type Response struct {
	Data      Data       `json:"data"`
	Status    DataStatus `json:"status"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
