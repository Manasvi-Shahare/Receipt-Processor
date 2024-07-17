package receipt

import (
	"time"
	"strings"
	"math"
	"regexp"
	"strconv"
	"errors"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ProcessResponse struct {
	ID string `json:"id"`
}

type PointsResponse struct {
	Points int `json:"points"`
}

var receipts = make(map[string]Receipt)
var points = make(map[string]int)

func GenerateID() string {
	return strings.ReplaceAll(time.Now().Format("20060102150405.000"), ".", "")
}

func CalculatePoints(receipt Receipt) int {
	points := 0

	alnumCount := len(regexp.MustCompile(`[a-zA-Z0-9]`).FindAllString(receipt.Retailer, -1))
	points += alnumCount

	total := receipt.Total
	if strings.HasSuffix(total, ".00") {
		points += 50
	}

	if totalFloat, err := strconv.ParseFloat(total, 64); err == nil && math.Mod(totalFloat, 0.25) == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			if price, err := strconv.ParseFloat(item.Price, 64); err == nil {
				points += int(math.Ceil(price * 0.2))
			}
		}
	}

	if date, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil {
		if date.Day()%2 != 0 {
			points += 6
		}
	}

	if purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
		if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
			points += 10
		}
	}

	return points
}

func StoreReceipt(id string, receipt Receipt) {
	receipts[id] = receipt
}

func GetPoints(id string) (int, error) {
	points, ok := points[id]
	if !ok {
		return 0, errors.New("receipt not found")
	}
	return points, nil
}

func StorePoints(id string, pointsVal int) {
	points[id] = pointsVal
}
