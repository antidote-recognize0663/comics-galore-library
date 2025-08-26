package config

import (
	"log"
	"sort"
)

type ApplicationConfig interface {
	GetEnv() string
	GetPort() int
	GetBaseUrl() string
	GetAppName() string
	GetJwtSecret() string
	GetSubscriptionPlans() *[]SubscriptionPlan
	GetCategories() *map[string]Category
	GetSortedCategories() []Category
	GetSortedCategoriesByName() []Category
	GetSubscriptionPlan(string) (SubscriptionPlan, bool)
	GetCategory(string) (Category, bool)
}

type applicationConfig struct {
	Env               string
	Port              int
	BaseUrl           string
	AppName           string
	JwtSecret         string
	SubscriptionPlans *[]SubscriptionPlan
	Categories        *map[string]Category
}

func (a *applicationConfig) GetEnv() string {
	return a.Env
}

func (a *applicationConfig) GetPort() int {
	return a.Port
}

func (a *applicationConfig) GetBaseUrl() string {
	return a.BaseUrl
}

func (a *applicationConfig) GetAppName() string {
	return a.AppName
}

func (a *applicationConfig) GetJwtSecret() string {
	return a.JwtSecret
}

func (a *applicationConfig) GetSubscriptionPlans() *[]SubscriptionPlan {
	return a.SubscriptionPlans
}

func (a *applicationConfig) GetCategories() *map[string]Category {
	return a.Categories
}

func (a *applicationConfig) GetSubscriptionPlan(id string) (SubscriptionPlan, bool) {
	if a.SubscriptionPlans == nil {
		return nil, false
	}
	for _, plan := range *a.SubscriptionPlans {
		if plan.GetID() == id {
			return plan, true
		}
	}
	return nil, false // Not found
}

func (a *applicationConfig) GetCategory(key string) (Category, bool) {
	if a.Categories == nil {
		return nil, false
	}
	category, ok := (*a.Categories)[key]
	return category, ok
}

func NewApplicationConfig() ApplicationConfig {
	var parseErr error
	config := &applicationConfig{
		Env:       GetEnv("ENV", "development"),
		Port:      GetIntEnv("PORT", 3000, &parseErr),
		AppName:   GetEnv("APP_NAME", "Comics Galore"),
		BaseUrl:   GetEnv("BASE_URL", "http://localhost:3000"),
		JwtSecret: GetEnv("JWT_SECRET", "PzBK+Wmb6LtK+8PfiLQ+dWLCsRnsTQm3v+He14YuZac="),
	}
	if parseErr != nil {
		log.Printf("error parsing integer environment variables: %v", parseErr)
	}
	return config
}

func (a *applicationConfig) GetSortedCategories() []Category {
	var sortedList []Category
	for _, category := range *a.Categories {
		sortedList = append(sortedList, category)
	}
	sort.Slice(sortedList, func(i, j int) bool {
		return sortedList[i].GetValue() < sortedList[j].GetValue()
	})
	return sortedList
}

func (a *applicationConfig) GetSortedCategoriesByName() []Category {
	if a.Categories == nil {
		return nil
	}
	categories := make([]Category, 0, len(*a.Categories))
	for _, category := range *a.Categories {
		categories = append(categories, category)
	}
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].GetName() < categories[j].GetName()
	})
	return categories
}
