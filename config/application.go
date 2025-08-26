package config

import (
	"log"
	"sort"
)

type ApplicationConfig interface {
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

func NewApplicationConfig() ApplicationConfig {
	var parseErr error
	config := &applicationConfig{
		Env:       GetEnv("ENV", "development"),
		Port:      GetIntEnv("PORT", 3000, &parseErr),
		BaseUrl:   GetEnv("BASE_URL", "http://localhost:3000"),
		AppName:   GetEnv("APP_NAME", "Comics Galore"),
		JwtSecret: GetEnv("JWT_SECRET", "PzBK+Wmb6LtK+8PfiLQ+dWLCsRnsTQm3v+He14YuZac="),
	}
	if parseErr != nil {
		log.Printf("error parsing integer environment variables: %v", parseErr)
	}
	return config
}

func (a *applicationConfig) SortedCategories() []Category {
	var sortedList []Category
	for _, category := range *a.Categories {
		sortedList = append(sortedList, category)
	}
	sort.Slice(sortedList, func(i, j int) bool {
		return sortedList[i].GetValue() < sortedList[j].GetValue()
	})
	return sortedList
}
