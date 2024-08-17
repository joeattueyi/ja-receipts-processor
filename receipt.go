package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type validable interface {
	valid() (valid bool)
}
type retailer string

func (v *retailer) valid() (valid bool) {
	m, err := regexp.Match("^[\\w\\s\\-&]+$", []byte(*v))
	if !m || err != nil {
		return false
	}
	return true
}

type purchaseDate string

func (v *purchaseDate) valid() (valid bool) {
	m, err := regexp.Match("^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\\d|3[01])$", []byte(*v))
	if !m || err != nil {
		return false
	}

	return true
}

type purchaseTime string

func (v *purchaseTime) valid() (valid bool) {
	m, err := regexp.Match("^([01]\\d|2[0-3]):[0-5]\\d$", []byte(*v))
	if !m || err != nil {
		return false
	}

	return true
}

type itemArray []Item

func (arr *itemArray) valid() (valid bool) {
	if len(*arr) < 1 {
		return false
	}

	for _, v := range *arr {
		if !v.valid() {
			return false
		}
	}
	return true
}

type total string

func (v *total) valid() (valid bool) {
	m, err := regexp.Match("^\\d+\\.\\d{2}$", []byte(*v))
	if !m || err != nil {
		return false
	}

	return true
}

type shortDescription string

func (v *shortDescription) valid() (valid bool) {
	m, err := regexp.Match("^[\\w\\s\\-]+$", []byte(*v))
	if !m || err != nil {
		return false
	}
	return true
}

type price string

func (v *price) valid() (valid bool) {
	m, err := regexp.Match("^\\d+\\.\\d{2}$", []byte(*v))
	if !m || err != nil {
		return false
	}

	return true
}

type Item struct {
	ShortDescription shortDescription `json:"shortDescription"`
	Price            price            `json:"price"`
}

func (v *Item) valid() (valid bool) {
	return v.ShortDescription.valid() && v.Price.valid()
}

type Receipt struct {
	Retailer     retailer     `json:"retailer"`
	PurchaseDate purchaseDate `json:"purchaseDate"`
	PurchaseTime purchaseTime `json:"purchaseTime"`
	Items        itemArray    `json:"items"`
	Total        total        `json:"total"`
}

func ValidateReceipt(v *Validator, receipt *Receipt) {
	v.Check(receipt.Retailer.valid(), "retailer", "retailer not valid")
	v.Check(receipt.PurchaseDate.valid(), "purchase date", "purchase date not valid")
	v.Check(receipt.PurchaseTime.valid(), "purchase time", "purchase time not valid")
	v.Check(receipt.Items.valid(), "items", "items not valid")
	if !receipt.Items.valid() {
		for _, item := range receipt.Items {
			v.Check(item.ShortDescription.valid(), "short description", "short description not valid")
			v.Check(item.Price.valid(), "price", "price not valid")
		}
	}
	v.Check(receipt.Total.valid(), "total", "total not valid")
}

func (r *Receipt) computePoints() (points int) {
	// alphanumeric characters in retailer
	for _, c := range r.Retailer {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			points++
		}
	}

	total, _ := strconv.ParseFloat(string(r.Total), 64)
	total100 := int(total * 100)

	// round dollar
	if (total100 % 100) == 0 {
		points += 50
	}

	// multiple of 0.25
	if (total100 % 25) == 0 {
		points += 25
	}

	// Every two items on receipt
	points += 5 * (int(len(r.Items) / 2))

	// f the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned
	for _, item := range r.Items {
		description := strings.TrimSpace(string(item.ShortDescription))
		if len(description)%3 == 0 {
			price, _ := strconv.ParseFloat(string(item.Price), 64)
			points += int(math.Ceil(0.2 * price))
		}
	}

	// odd day
	if int(r.PurchaseDate[len(r.PurchaseDate)-1])%2 == 1 {
		points += 6
	}

	// purchase time after 2pm & before 4pm
	purchaseHour, _ := strconv.ParseInt(string(r.PurchaseTime[0:2]), 10, 64)
	purchaseMinute, _ := strconv.ParseInt(string(r.PurchaseTime[3:]), 10, 64)
	if (purchaseHour == 14 && purchaseMinute > 0) || purchaseHour == 15 {
		points += 10
	}

	return points
}
