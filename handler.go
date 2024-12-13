package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func FetchAvailability(roomID string) ([]AvailabilityData, error) {
	url := "https://airbnb-listings.p.rapidapi.com/v2/listingAvailabilityFull?id=" + roomID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("x-rapidapi-key", "f9df31e10dmshd92a56e4655fbaap18456bjsnb782c8e06c84")
	req.Header.Add("x-rapidapi-host", "airbnb-listings.p.rapidapi.com")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse AvailabilityAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return apiResponse.Results, nil
}

func FetchPrice(roomID string) ([]PriceData, error) {
	url := "https://airbnb-listings.p.rapidapi.com/v2/listingPricesFull?id=" + roomID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-rapidapi-key", "f9df31e10dmshd92a56e4655fbaap18456bjsnb782c8e06c84")
	req.Header.Add("x-rapidapi-host", "airbnb-listings.p.rapidapi.com")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponse PriceAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.Results, nil
}

func CalculateMetrics(roomId string, availability []AvailabilityData, prices []PriceData) APIResponse {
	occupancy := make(map[string]float64)
	now := time.Now()
	next30Days := now.AddDate(0, 0, 30)

	monthData := make(map[string][]AvailabilityData)
	for _, day := range availability {
		date, _ := ParseDate(day.Date)

		if date.After(now) && date.Before(now.AddDate(0, 5, 0)) {
			month := date.Format("2006-01")
			monthData[month] = append(monthData[month], day)
		}
	}

	for month, days := range monthData {
		totalDays := len(days)
		availableDays := 0

		for _, day := range days {
			if day.Available > 0 {
				availableDays++
			}
		}

		occupancy[month] = (float64(availableDays) / float64(totalDays)) * 100
	}

	var totalRate, highestRate, lowestRate float64
	var count int
	lowestRate = 1e9

	for _, price := range prices {
		date, _ := ParseDate(price.Date)

		if date.After(now) && date.Before(next30Days) {
			totalRate += price.Price

			if price.Price > highestRate {
				highestRate = price.Price
			}

			if price.Price < lowestRate {
				lowestRate = price.Price
			}

			count++
		}
	}

	averageRate := totalRate / float64(count)

	return APIResponse{
		RequestedRoomID:          roomId,
		OccupancyPercentage:      occupancy,
		AverageRateForNext30days: averageRate,
		HighestRateForNext30days: highestRate,
		LowestRateForNext30days:  lowestRate,
	}
}
