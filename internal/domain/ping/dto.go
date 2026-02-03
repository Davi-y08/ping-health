package ping

import "github.com/google/uuid"

type CreatePingDto struct {
	StatusCode 	int    		`json:"status_code"`
	Status     	bool   		`json:"status"`
	URL        	string 		`json:"url"`
	MonitorID  	uuid.UUID 	`json:"monitor_id"`
}