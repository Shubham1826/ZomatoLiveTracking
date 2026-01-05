package main

type LocationEvent struct {
	PartnerID string  `json:"partner_id"`
	OrderID   string  `json:"order_id"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Timestamp int64   `json:"timestamp"`
}
