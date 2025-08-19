package model

import (
	"math"
)

type SubscriptionPlan struct {
	ID       string
	Name     string
	Price    float64
	Duration string
	Discount float64
}

func (sp *SubscriptionPlan) GetPrice() float64 {
	if sp == nil || sp.Discount < 0 || sp.Discount > 1 {
		return 0
	}
	discountedPrice := sp.Price - (sp.Price * sp.Discount)
	return math.Round(discountedPrice*100) / 100
}

func NewSubscriptionPlan(id, name string, price float64, duration string, discount float64) *SubscriptionPlan {
	return &SubscriptionPlan{
		ID:       id,
		Name:     name,
		Price:    price,
		Duration: duration,
		Discount: discount,
	}
}

func NewSubscriptionPlans() *[]SubscriptionPlan {
	return &[]SubscriptionPlan{
		*NewSubscriptionPlan("1", "Monthly Plan", 10, "1 month", 0.0),
		*NewSubscriptionPlan("2", "Quarterly Plan", 30, "3 month", 0.0),
		*NewSubscriptionPlan("3", "Semi-annual Plan", 60, "6 month", 0.1),
		*NewSubscriptionPlan("4", "Yearly Plan", 120, "1 year", 0.2),
	}
}

type Category struct {
	Value string
	Name  string
	Href  string
}

func NewCategory(value, name, href string) *Category {
	return &Category{
		Value: value,
		Name:  name,
		Href:  href,
	}
}

func NewCategories() *map[string]Category {
	return &map[string]Category{
		"one":   *NewCategory("1", "Technology", "/category/technology"),
		"two":   *NewCategory("2", "Science", "/category/science"),
		"three": *NewCategory("3", "Business", "/category/business"),
		"four":  *NewCategory("4", "Entertainment", "/category/entertainment"),
	}
}
