package account

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/account"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/users"
)

type Admin interface {
	SignIn(email, password string) (*model.Session, error)
	GetUser(userId string) (*model.Account, error)
	SignUp(username, email, password string) (*model.Account, error)
	PasswordReset(email string, recoveryUrl string) (*model.Token, error)
	UpdateVerification(secret, userId string) (*model.Token, error)
}

type admin struct {
	user    *users.Users
	account *account.Account
}

func NewAdminWithConfig(config *config.Config) Admin {
	adminClient := utils.NewAdminClient(config.Appwrite.ApiKey, utils.WithProject(config.Appwrite.ProjectID), utils.WithEndpoint(config.Appwrite.Endpoint))
	return &admin{
		user:    appwrite.NewUsers(*adminClient),
		account: appwrite.NewAccount(*adminClient),
	}
}

func NewAdmin(client *client.Client) Admin {
	return &admin{
		user:    appwrite.NewUsers(*client),
		account: appwrite.NewAccount(*client),
	}
}

func (s *admin) SignIn(email, password string) (*model.Session, error) {
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}
	if password == "" {
		return nil, fmt.Errorf("password cannot be empty")
	}
	session, err := s.account.CreateEmailPasswordSession(email, password)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}
	return &model.Session{Session: session}, nil
}

func (s *admin) GetUser(userId string) (*model.Account, error) {
	if userId == "" {
		return nil, fmt.Errorf("userId cannot be empty")
	}
	user, err := s.user.Get(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user %s: %w", userId, err)
	}
	return model.NewAccount(user), nil
}

func (s *admin) SignUp(username, email, password string) (*model.Account, error) {
	if username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}
	if password == "" {
		return nil, fmt.Errorf("password cannot be empty")
	}
	user, err := s.user.Create(
		id.Unique(),
		s.user.WithCreateEmail(email),
		s.user.WithCreatePassword(password),
		s.user.WithCreateName(username))
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	_, err = s.user.UpdatePrefs(user.Id, model.Prefs{
		AvatarID: "",
		Twitter:  "",
		Facebook: "",
		Tumblr:   "",
	})
	return model.NewAccount(user), nil
}

func (s *admin) PasswordReset(email string, recoveryUrl string) (*model.Token, error) {
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}
	if recoveryUrl == "" {
		return nil, fmt.Errorf("recoveryUrl cannot be empty")
	}
	token, err := s.account.CreateRecovery(email, recoveryUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create password recovery: %w", err)
	}
	return &model.Token{Token: token}, nil
}

func (s *admin) UpdateVerification(secret, userId string) (*model.Token, error) {
	if secret == "" {
		return nil, fmt.Errorf("secret cannot be empty")
	}
	if userId == "" {
		return nil, fmt.Errorf("userId cannot be empty")
	}
	token, err := s.account.UpdateVerification(userId, secret)
	if err != nil {
		return nil, fmt.Errorf("failed to update verification: %w", err)
	}
	return &model.Token{Token: token}, nil
}
