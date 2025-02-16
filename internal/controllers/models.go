package controllers

import "time"

type TicketAcceptance struct {
	ReadyBy time.Time `form:"readyBy" json:"readyBy" binding:"required" time_format:"rfc3339"`
}
