package main

import "testing"

// TODO: ADD MALFORMED RECEIPTS TO TESTS
func TestReceiptValidation(t *testing.T) {

	t.Run("Receipt 1", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []Item{
				{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
				{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
				{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
				{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
			},
			Total: "35.35",
		}

		if !receipt.Retailer.valid() {
			t.Fatal("Retailer not valid")
		}

		if !receipt.PurchaseDate.valid() {
			t.Fatal("Purchase Date not valid")
		}

		if !receipt.PurchaseTime.valid() {
			t.Fatal("Purchase Time not valid")
		}

		if !receipt.Items.valid() {
			t.Fatal("Purchase Date not valid")
		}

		if !receipt.Total.valid() {
			t.Fatal("Total not valid")
		}
	})

	t.Run("Receipt 2", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "M&M Corner Market",
			PurchaseDate: "2022-03-20",
			PurchaseTime: "14:33",
			Items: []Item{
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
			},
			Total: "9.00",
		}

		if !receipt.Retailer.valid() {
			t.Fatal("Retailer not valid")
		}

		if !receipt.PurchaseDate.valid() {
			t.Fatal("Purchase Date not valid")
		}

		if !receipt.PurchaseTime.valid() {
			t.Fatal("Purchase Time not valid")
		}

		if !receipt.Items.valid() {
			t.Fatal("Purchase Date not valid")
		}

		if !receipt.Total.valid() {
			t.Fatal("Total not valid")
		}
	})

	t.Run("Bad Retailer", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "=++++===",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []Item{
				{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
				{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
				{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
				{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
			},
			Total: "35.35",
		}

		if receipt.Retailer.valid() {
			t.Fatal("Retailer should not be valid")
		}
	})

	t.Run("Bad Purchase Date", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01foo",
			PurchaseTime: "13:01",
			Items: []Item{
				{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
				{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
				{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
				{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
			},
			Total: "35.35",
		}

		if receipt.PurchaseDate.valid() {
			t.Fatal("Purchase date should not be valid")
		}
	})

	t.Run("Bad Purchase Time", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13.01bar",
			Items: []Item{
				{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
				{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
				{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
				{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
			},
			Total: "35.35",
		}

		if receipt.PurchaseTime.valid() {
			t.Fatal("Purchase time should not be valid")
		}
	})

	t.Run("Bad Item", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []Item{
				{ShortDescription: "Mountain Dew 12PK", Price: "6.49bux"},
				{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
				{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
				{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
				{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
			},
			Total: "35.35",
		}

		if receipt.Items.valid() {
			t.Fatal("Items should not be valid")
		}

		if receipt.Items[0].valid() {
			t.Fatal("First item should not be valid")
		}

		if receipt.Items[0].Price.valid() {
			t.Fatal("First item price should not be valid")
		}
	})

	t.Run("Item empty", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items:        []Item{},
			Total:        "35.35",
		}

		if receipt.Items.valid() {
			t.Fatal("Items should not be valid")
		}
	})

	t.Run("Bad Total", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []Item{
				{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
				{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
				{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
				{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
			},
			Total: "35.+++/",
		}

		if receipt.Total.valid() {
			t.Fatal("Total should not be valid")
		}
	})

}

func TestComputePrice(t *testing.T) {

	t.Run("Receipt 1", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []Item{
				{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
				{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
				{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
				{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
			},
			Total: "35.35",
		}

		if receipt.computePoints() != 28 {
			t.Fatal("Compute points is wrong")
		}
	})

	t.Run("Receipt 2", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "M&M Corner Market",
			PurchaseDate: "2022-03-20",
			PurchaseTime: "14:33",
			Items: []Item{
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
			},
			Total: "9.00",
		}

		if receipt.computePoints() != 109 {
			t.Log(receipt.computePoints())
			t.Fatal("Compute points is wrong")
		}
	})

	t.Run("Receipt 2, @ 4pm", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "M&M Corner Market",
			PurchaseDate: "2022-03-20",
			PurchaseTime: "16:00",
			Items: []Item{
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
			},
			Total: "9.00",
		}

		if receipt.computePoints() != 99 {
			t.Log(receipt.computePoints())
			t.Fatal("Compute points is wrong")
		}
	})

	t.Run("Receipt 2, @ 2pm", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "M&M Corner Market",
			PurchaseDate: "2022-03-20",
			PurchaseTime: "14:00",
			Items: []Item{
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
			},
			Total: "9.00",
		}

		if receipt.computePoints() != 99 {
			t.Log(receipt.computePoints())
			t.Fatal("Compute points is wrong")
		}
	})
	t.Run("Receipt 2, odd day", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "M&M Corner Market",
			PurchaseDate: "2022-03-21",
			PurchaseTime: "14:33",
			Items: []Item{
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
			},
			Total: "9.00",
		}

		if receipt.computePoints() != 115 {
			t.Log(receipt.computePoints())
			t.Fatal("Compute points is wrong")
		}
	})

	t.Run("Receipt 2, with two items", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "M&M Corner Market",
			PurchaseDate: "2022-03-20",
			PurchaseTime: "14:33",
			Items: []Item{
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
			},
			Total: "9.00",
		}

		if receipt.computePoints() != 104 {
			t.Log(receipt.computePoints())
			t.Fatal("Compute points is wrong")
		}
	})

	t.Run("Receipt 2, description is multiple of 3", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "M&M Corner Market",
			PurchaseDate: "2022-03-20",
			PurchaseTime: "14:33",
			Items: []Item{
				{ShortDescription: "GatoradeX", Price: "2.25"},
				{ShortDescription: "GatoradeX", Price: "2.25"},
				{ShortDescription: "GatoradeX", Price: "2.25"},
				{ShortDescription: "GatoradeX", Price: "2.25"},
			},
			Total: "9.00",
		}

		if receipt.computePoints() != 113 {
			t.Log(receipt.computePoints())
			t.Fatal("Compute points is wrong")
		}
	})

	t.Run("Receipt 2, total is not round dollar & total is not multiple of 0.25", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "M&M Corner Market",
			PurchaseDate: "2022-03-20",
			PurchaseTime: "14:33",
			Items: []Item{
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
			},
			Total: "9.01",
		}

		if receipt.computePoints() != 34 {
			t.Log(receipt.computePoints())
			t.Fatal("Compute points is wrong")
		}
	})

	t.Run("Receipt 2, retailer name has extra character", func(t *testing.T) {
		receipt := Receipt{
			Retailer:     "M&M Corner MarketX",
			PurchaseDate: "2022-03-20",
			PurchaseTime: "14:33",
			Items: []Item{
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
				{ShortDescription: "Gatorade", Price: "2.25"},
			},
			Total: "9.00",
		}

		if receipt.computePoints() != 110 {
			t.Log(receipt.computePoints())
			t.Fatal("Compute points is wrong")
		}
	})
}
