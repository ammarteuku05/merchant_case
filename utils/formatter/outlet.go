package formatter

import (
	"merchant-service/entity"
	"time"
)

type OutletFormat struct {
	ID         string `json:"id"`
	OutletName string `json:"outlet_name"`
	Picture    string `json:"picture"`
	UserID     string `json:"user_id"`
}

type OutletDeleteFormat struct {
	Message    string    `json:"message"`
	TimeDelete time.Time `json:"time_delete"`
}

func FormatOutlet(outlet entity.Outlet) OutletFormat {
	var formatOutlet = OutletFormat{
		ID:         outlet.Id,
		OutletName: outlet.OutletName,
		Picture:    outlet.Picture,
		UserID:     outlet.UserId,
	}

	return formatOutlet
}

func FormatDeleteOutlet(msg string) OutletDeleteFormat {
	var deleteFormat = OutletDeleteFormat{
		Message:    msg,
		TimeDelete: time.Now(),
	}

	return deleteFormat
}
