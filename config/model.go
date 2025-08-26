package config

import (
	"math"
)

type SubscriptionPlan interface {
	GetID() string
	GetName() string
	GetPrice() float64
	GetDuration() string
}

type subscriptionPlan struct {
	ID       string
	Name     string
	Price    float64
	Duration string
	Discount float64
}

func (sp *subscriptionPlan) GetID() string       { return sp.ID }
func (sp *subscriptionPlan) GetName() string     { return sp.Name }
func (sp *subscriptionPlan) GetDuration() string { return sp.Duration }
func (sp *subscriptionPlan) GetPrice() float64 {
	if sp == nil || sp.Discount < 0 || sp.Discount >= 1 {
		return sp.Price
	}
	discountedPrice := sp.Price - (sp.Price * sp.Discount)
	return math.Round(discountedPrice*100) / 100
}

func NewSubscriptionPlan(id, name string, price float64, duration string, discount float64) SubscriptionPlan {
	return &subscriptionPlan{
		ID:       id,
		Name:     name,
		Price:    price,
		Duration: duration,
		Discount: discount,
	}
}

func NewSubscriptionPlans() []SubscriptionPlan {
	return []SubscriptionPlan{
		NewSubscriptionPlan("1", "Monthly Plan", 10, "1 month", 0.0),
		NewSubscriptionPlan("2", "Quarterly Plan", 30, "3 months", 0.0),
		NewSubscriptionPlan("3", "Semi-annual Plan", 60, "6 months", 0.1),
		NewSubscriptionPlan("4", "Yearly Plan", 120, "1 year", 0.2),
	}
}

type Category interface {
	GetValue() string
	GetName() string
	GetHref() string
}

type category struct {
	Value string
	Name  string
	Href  string
}

func (c *category) GetValue() string { return c.Value }
func (c *category) GetName() string  { return c.Name }
func (c *category) GetHref() string  { return c.Href }

func NewCategory(value, name, href string) Category {
	return &category{
		Value: value,
		Name:  name,
		Href:  href,
	}
}

func NewCategories() map[string]Category {
	return map[string]Category{
		"misc":        NewCategory("0", "Misc", "/category/misc"),
		"hentai":      NewCategory("1", "Hentai / Doujinshi", "/category/hentai"),
		"erotica":     NewCategory("2", "Mature / Erotica", "/category/erotica"),
		"yuri":        NewCategory("3", "Yuri / GL (Girls' Love)", "/category/yuri"),
		"yaoi":        NewCategory("4", "Yaoi / BL (Boys' Love)", "/category/yaoi"),
		"adult-scifi": NewCategory("5", "Adult Sci-Fi & Fantasy", "/category/adult-sci-fi"),
		"milf":        NewCategory("7", "MILF", "/category/milf"),
		"interracial": NewCategory("8", "Interracial", "/category/interracial"),
		"big-tits":    NewCategory("9", "Big Tits", "/category/big-tits"),
		"giantess":    NewCategory("10", "Giantess", "/category/giantess"),
		"ai-porn":     NewCategory("11", "AI Porn", "/category/ai-porn"),
		"3d-porn":     NewCategory("12", "3D Porn", "/category/3d-porn"),
		"futanari":    NewCategory("13", "Futanari", "/category/futanari"),
		"site-rip":    NewCategory("14", "SiteRips", "/category/siterips"),
	}
}
