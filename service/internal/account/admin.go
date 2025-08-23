package account

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/id"
)

type Admin interface {
	SignIn(email, password string) (*model.Session, error)
	GetUser(userId string) (*model.User, error)
	SignUp(username, email, password string) (*model.Account, error)
	PasswordReset(email string, recoveryUrl string) (*model.Token, error)
	UpdateVerification(secret, userId string) (*model.Token, error)
}

type admin struct {
	apiKey    string
	endpoint  string
	projectID string
}

func NewAdminWithConfig(config *config.Config) Admin {
	return &admin{
		apiKey:    config.Appwrite.ApiKey,
		endpoint:  config.Appwrite.Endpoint,
		projectID: config.Appwrite.ProjectID,
	}
}

func NewAdmin(options ...Option) Admin {
	_config := &Config{
		endpoint:  "https://fra.cloud.appwrite.io/v1",
		projectID: "6510a59f633f9d57fba2",
	}
	for _, option := range options {
		option(_config)
	}
	return &admin{
		apiKey:    _config.apiKey,
		endpoint:  _config.endpoint,
		projectID: _config.projectID,
	}
}

func (s *admin) SignIn(email, password string) (*model.Session, error) {
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}
	if password == "" {
		return nil, fmt.Errorf("password cannot be empty")
	}
	account := appwrite.NewAccount(*utils.NewAdminClient(s.apiKey, utils.WithProject(s.projectID), utils.WithEndpoint(s.endpoint)))
	session, err := account.CreateEmailPasswordSession(email, password)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}
	return &model.Session{Session: session}, nil
}

func (s *admin) GetUser(userId string) (*model.User, error) {
	if userId == "" {
		return nil, fmt.Errorf("userId cannot be empty")
	}
	users := appwrite.NewUsers(*utils.NewAdminClient(s.apiKey, utils.WithProject(s.projectID), utils.WithEndpoint(s.endpoint)))
	user, err := users.Get(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user %s: %w", userId, err)
	}
	//TODO: Don't forget to map Preferences
	return &model.User{User: user}, nil
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
	users := appwrite.NewUsers(*utils.NewAdminClient(s.apiKey, utils.WithProject(s.projectID), utils.WithEndpoint(s.endpoint)))
	user, err := users.Create(
		id.Unique(),
		users.WithCreateEmail(email),
		users.WithCreatePassword(password),
		users.WithCreateName(username))
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return model.NewAccount(user), nil
}

func (s *admin) PasswordReset(email string, recoveryUrl string) (*model.Token, error) {
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}
	if recoveryUrl == "" {
		return nil, fmt.Errorf("recoveryUrl cannot be empty")
	}
	account := appwrite.NewAccount(*utils.NewAdminClient(s.apiKey, utils.WithProject(s.projectID), utils.WithEndpoint(s.endpoint)))
	token, err := account.CreateRecovery(email, recoveryUrl)
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
	account := appwrite.NewAccount(*utils.NewAdminClient(s.apiKey, utils.WithProject(s.projectID), utils.WithEndpoint(s.endpoint)))
	token, err := account.UpdateVerification(userId, secret)
	if err != nil {
		return nil, fmt.Errorf("failed to update verification: %w", err)
	}
	return &model.Token{Token: token}, nil
}
