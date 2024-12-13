package main

type AvailabilityAPIResponse struct {
	RequestID string             `json:"requestId"`
	Results   []AvailabilityData `json:"results"`
}

type AvailabilityData struct {
	Date                string `json:"date"`
	Available           int    `json:"available"`
	ClosedToArrival     int    `json:"closed_to_arrival"`
	ClosedToDeparture   int    `json:"closed_to_departure"`
	AvailableForCheckin int    `json:"available_for_checkin"`
	MinNights           int    `json:"minNights"`
	MaxNights           int    `json:"maxNights"`
}

type PriceAPIResponse struct {
	RequestID string      `json:"requestId"`
	Results   []PriceData `json:"results"`
}

type PriceData struct {
	Date     string  `json:"date"`
	Price    float64 `json:"price"`
	PriceEUR float64 `json:"price_eur"`
	PriceUSD float64 `json:"price_usd"`
}

type APIResponse struct {
	RequestID           string             `json:"request_id"`
	OccupancyPercentage map[string]float64 `json:"occupancy_percentage"`
	AverageRate         float64            `json:"average_rate"`
	HighestRate         float64            `json:"highest_rate"`
	LowestRate          float64            `json:"lowest_rate"`
}
