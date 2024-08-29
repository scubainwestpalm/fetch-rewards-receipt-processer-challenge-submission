package models

import (
	"errors"
	"math"
	"strings"
	"time"

	"fetch.dev/receipt-processor-challenge-submission/utils"
	"github.com/google/uuid"
)

type Receipt struct {
	Retailer     string  `json:"retailer" binding:"required"`
	PurchaseDate string  `json:"purchaseDate" binding:"required"`
	PurchaseTime string  `json:"purchaseTime" binding:"required"`
	Items        []Item  `json:"items" binding:"required"`
	Total        float64 `json:"total,string" binding:"required"`
	Points       int
}

var receipts map[string]Receipt = make(map[string]Receipt)

func (r *Receipt) Process() (string, error) {
	var id = uuid.NewString()

	purchaseDay, hours, mins, err := r.parseTimeDataFromDateAndTimeStrings()

	if err != nil {
		return "", err
	}

	r.Points = 0

	// +1 for every alphanumeric char in retailer name
	r.Points += utils.CountAlphanumericCharsInString(r.Retailer)

	// +50 if total is whole number
	if utils.Float64IsWholeNumber(r.Total) {
		r.Points += 50
	}

	// +25 if total is multiple of .25
	if utils.Float64IsMultipleOf(r.Total, 0.25) {
		r.Points += 25
	}

	// +5 Points for every two items in receipt
	r.Points += ((len(r.Items) / 2) * 5)

	/* If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned */
	for _, item := range r.Items {
		trimedDesc := strings.Trim(item.ShortDescription, " ")
		trimmedLenIsMultipleOfThree := utils.IntIsMultipleOf(len(trimedDesc), 3)
		if trimmedLenIsMultipleOfThree {
			r.Points += int(math.Ceil(item.Price * 0.2))
		}
	}

	// +6 if day in purchase date is odd
	if purchaseDay%2 != 0 {
		r.Points += 6
	}

	// +10 if time of purchase is > 2pm < 4pm
	if ((hours == 14 && mins > 0) || hours > 14) && hours < 16 {
		r.Points += 10
	}

	receipts[id] = *r
	return id, nil
}

func (r *Receipt) parseTimeDataFromDateAndTimeStrings() (purchaseDay int, hours int, mins int, err error) {
	// Assumes +0000 UTC
	parsedTime, err := time.Parse(time.RFC3339, strings.Join([]string{r.PurchaseDate, "T", r.PurchaseTime, ":00Z"}, ""))

	if err != nil {
		return 0, 0, 0, err
	}

	parsedHours, parsedMins, _ := parsedTime.Clock()

	return parsedTime.Day(), parsedHours, parsedMins, nil
}

func GetReceiptById(id string) (Receipt, error) {
	if id == "" {
		return Receipt{}, errors.New("no id specified for receipt")
	}

	var r, exists = receipts[id]

	if !exists {
		return Receipt{}, errors.New("receipt does not exist under specified id")
	}

	return r, nil
}
