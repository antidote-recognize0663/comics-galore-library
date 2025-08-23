package user

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/models"
)

type User interface {
	AddLabel(userId, label string) (*model.User, error)
	RemoveLabel(userId, label string) (*model.User, error)
}

type user struct {
	endpoint string
	project  string
	apiKey   string
}

type Config struct {
	endpoint string
	project  string
	apiKey   string
}

type Option func(*Config)

func WithEndpoint(endpoint string) Option {
	return func(cfg *Config) {
		cfg.endpoint = endpoint
	}
}
func WithProject(project string) Option {
	return func(cfg *Config) {
		cfg.project = project
	}
}
func WithApiKey(apiKey string) Option {
	return func(cfg *Config) {
		cfg.apiKey = apiKey
	}
}

func NewUser(options ...Option) User {
	cfg := &Config{
		endpoint: "",
		project:  "",
		apiKey:   "",
	}
	for _, option := range options {
		option(cfg)
	}
	return &user{
		endpoint: cfg.endpoint,
		project:  cfg.project,
		apiKey:   cfg.apiKey,
	}
}

func NewUserWithConfig(cfg *config.Config) User {
	return &user{
		endpoint: cfg.Appwrite.Endpoint,
		project:  cfg.Appwrite.ProjectID,
		apiKey:   cfg.Appwrite.ApiKey,
	}
}

func mapUser(user *models.User) (*model.User, error) {
	var prefs model.Preferences
	if err := user.Decode(&prefs); err != nil {
		return nil, err
	}
	return &model.User{
		User:        user,
		Preferences: &prefs,
	}, nil
}

func (s *user) AddLabel(userId, label string) (*model.User, error) {
	userDB := appwrite.NewUsers(*utils.NewAdminClient(s.apiKey, utils.WithEndpoint(s.endpoint), utils.WithProject(s.project), utils.WithEndpoint(s.endpoint)))
	fetchedUser, err := userDB.Get(userId)
	if err != nil {
		return nil, fmt.Errorf("GetUser error for userId '%s': %v", userId, err)
	}
	containsSubscriber := false
	for _, l := range fetchedUser.Labels {
		if l == label {
			containsSubscriber = true
			break
		}
	}
	if !containsSubscriber {
		user, err := userDB.UpdateLabels(userId, append(fetchedUser.Labels, "subscriber"))
		if err != nil {
			return nil, fmt.Errorf("AddLabel error for userId '%s': %v", userId, err)
		}
		return mapUser(user)
	}
	return mapUser(fetchedUser)
}

func (s *user) RemoveLabel(userId, label string) (*model.User, error) {
	userDB := appwrite.NewUsers(*utils.NewAdminClient(s.apiKey, utils.WithEndpoint(s.endpoint), utils.WithProject(s.project), utils.WithEndpoint(s.endpoint)))
	fetchedUser, err := userDB.Get(userId)
	if err != nil {
		return nil, fmt.Errorf("GetUser error for userId '%s': %v", userId, err)
	}
	fetchedUser.Labels = utils.Filter(fetchedUser.Labels, func(l string) bool {
		return l != label
	})
	user, err := userDB.UpdateLabels(userId, fetchedUser.Labels)
	if err != nil {
		return nil, fmt.Errorf("RemoveLabel error for userId '%s': %v", userId, err)
	}
	return mapUser(user)
}
