package controllers

import "time"

type ticketAcceptance struct {
	ReadyBy time.Time `form:"readyBy" json:"readyBy" binding:"required" time_format:"rfc3339"`
}
