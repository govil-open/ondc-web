package models

type OnSubscribe struct {
	SubscriberId string `json:"subscriber_id"`
	Challenge    string `json:"challenge"`
}
