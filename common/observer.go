package common

import (
	"time"
)

type Observer struct {
	Latitude  float64
	Longitude float64
	Location  *time.Location
}
