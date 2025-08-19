package services

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/users"
)

type UserService interface {
	AddSubscriberToLabels(userId string) (*model.User, error)
	RemoveSubscriberFromLabels(userId string) (*model.User, error)
}

type userService struct {
	client *client.Client
}

func NewUserService(client *client.Client) UserService {
	return &userService{
		client: client,
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

func (s *userService) AddSubscriberToLabels(userId string) (*model.User, error) {
	userDB := users.New(*s.client)
	fetchedUser, err := userDB.Get(userId)
	if err != nil {
		return nil, fmt.Errorf("GetUser error for userId '%s': %v", userId, err)
	}
	containsSubscriber := false
	for _, label := range fetchedUser.Labels {
		if label == "subscriber" {
			containsSubscriber = true
			break
		}
	}
	if !containsSubscriber {
		user, err := userDB.UpdateLabels(userId, append(fetchedUser.Labels, "subscriber"))
		if err != nil {
			return nil, fmt.Errorf("UpdateLabels error for userId '%s': %v", userId, err)
		}
		return mapUser(user)
	}
	return mapUser(fetchedUser)
}

func (s *userService) RemoveSubscriberFromLabels(userId string) (*model.User, error) {
	userDB := users.New(*s.client)
	fetchedUser, err := userDB.Get(userId)
	if err != nil {
		return nil, fmt.Errorf("GetUser error for userId '%s': %v", userId, err)
	}
	fetchedUser.Labels = utils.Filter(fetchedUser.Labels, func(label string) bool {
		return label != "subscriber"
	})
	user, err := userDB.UpdateLabels(userId, fetchedUser.Labels)
	if err != nil {
		return nil, fmt.Errorf("UpdateLabels error for userId '%s': %v", userId, err)
	}
	return mapUser(user)
}
