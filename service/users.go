package services

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/users"
)

type UserService interface {
	AddSubscriberToUserLabels(userId string) (*model.User, error)
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

func (s *userService) AddSubscriberToUserLabels(userId string) (*model.User, error) {
	userDB := users.New(*s.client)
	fetchedUser, err := userDB.Get(userId)
	if err != nil {
		return nil, fmt.Errorf("UpdateLabels error for userId '%s': %v", userId, err)
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
