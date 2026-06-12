package routes

import "time"

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL
	CustomShort
	Expiry
	XRateRemaining
	XRateLimitRest
}
