package user

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
)

type User interface {
	AddLabel(userId, label string) (*model.Account, error)
	RemoveLabel(userId, label string) (*model.Account, error)
}

type user struct {
	client *client.Client
}

func NewUser(client *client.Client) User {
	if client == nil {
		panic("appwrite client is mandatory")
	}
	return &user{
		client: client,
	}
}

func NewUserWithConfig(cfg *config.Config) User {
	return &user{
		client: utils.NewAdminClient(cfg.Appwrite.ApiKey, utils.WithProject(cfg.Appwrite.ProjectID), utils.WithEndpoint(cfg.Appwrite.Endpoint)),
	}
}

func (s *user) AddLabel(userId, label string) (*model.Account, error) {
	userDB := appwrite.NewUsers(*s.client)
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
		userAccount, err := userDB.UpdateLabels(userId, append(fetchedUser.Labels, "subscriber"))
		if err != nil {
			return nil, fmt.Errorf("AddLabel error for userId '%s': %v", userId, err)
		}
		return model.NewAccount(userAccount), nil
	}
	return model.NewAccount(fetchedUser), nil
}

func (s *user) RemoveLabel(userId, label string) (*model.Account, error) {
	database := appwrite.NewUsers(*s.client)
	fetchedUser, err := database.Get(userId)
	if err != nil {
		return nil, fmt.Errorf("GetUser error for userId '%s': %v", userId, err)
	}
	fetchedUser.Labels = utils.Filter(fetchedUser.Labels, func(l string) bool {
		return l != label
	})
	userAccount, err := database.UpdateLabels(userId, fetchedUser.Labels)
	if err != nil {
		return nil, fmt.Errorf("RemoveLabel error for userId '%s': %v", userId, err)
	}
	return model.NewAccount(userAccount), nil
}
