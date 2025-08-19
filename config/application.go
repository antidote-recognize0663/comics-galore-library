package config

import (
	"github.com/antidote-recognize0663/comics-galore-library/config/utils"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"log"
	"sort"
)

type ApplicationConfig struct {
	Env               string
	Port              int
	BaseUrl           string
	AppName           string
	JwtSecret         string
	SubscriptionPlans *[]model.SubscriptionPlan
	Categories        *map[string]model.Category
}

func NewApplicationConfig() *ApplicationConfig {
	var parseErr error
	config := &ApplicationConfig{
		Env:       utils.GetEnv("ENV", "development"),
		Port:      utils.GetIntEnv("PORT", 3000, &parseErr),
		BaseUrl:   utils.GetEnv("BASE_URL", "http://localhost:3000"),
		AppName:   utils.GetEnv("APP_NAME", "Comics Galore"),
		JwtSecret: utils.GetEnv("JWT_SECRET", "PzBK+Wmb6LtK+8PfiLQ+dWLCsRnsTQm3v+He14YuZac="),
	}
	if parseErr != nil {
		log.Printf("error parsing integer environment variables: %w", parseErr)
	}
	return config
}

func (a *ApplicationConfig) SortedCategories() []model.Category {
	var sortedList []model.Category
	for _, category := range *a.Categories {
		sortedList = append(sortedList, category)
	}
	sort.Slice(sortedList, func(i, j int) bool {
		return sortedList[i].Value < sortedList[j].Value
	})
	return sortedList
}
