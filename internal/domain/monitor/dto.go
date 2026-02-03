package monitor

import "github.com/google/uuid"

type CreateMonitorDto struct {
	URL      	string 	`json:"url"`
	Interval 	int    	`json:"interval"`
	UserID   	uuid.UUID 	`json:"user_id"`
}