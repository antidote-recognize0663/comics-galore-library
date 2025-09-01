package account

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/models"
	"resty.dev/v3"
)

type Session interface {
	DeleteCurrentSession(secret string) error
	GetAccount(secret string) (*model.Account, error)
	GetAccount2(secret string) (*model.Account, error)
	UpdatePreferences(secret string, prefs *model.Prefs) (*model.Account, error)
	CreateVerification(secret string, verificationUrl string) (*model.Token, error)
	VerifyAccount(secret, userId string) (*model.Token, error)
	UpdateName(secret, name string) (*model.Account, error)
	UpdateEmail(secret, email, password string) (*model.Account, error)
	UpdatePassword(secret, oldPassword, newPassword string) (*model.Account, error)
	GetPrefs(secret string) (*model.Prefs, error)
}

type session struct {
	endpoint  string
	projectID string
}

func (s *session) GetAccount2(secret string) (*model.Account, error) {
	if secret == "" {
		return nil, fmt.Errorf("secret cannot be empty")
	}
	account := appwrite.NewAccount(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.endpoint), utils.WithProject(s.projectID)))
	data, err := account.Get()
	if err != nil {
		return nil, fmt.Errorf("error getting account: %v", err)
	}
	var response model.Account
	if err := data.Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding account: %v", err)
	}
	return &response, nil
}

func (s *session) GetPrefs(secret string) (*model.Prefs, error) {
	if secret == "" {
		return nil, fmt.Errorf("secret cannot be empty")
	}
	account := appwrite.NewAccount(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.endpoint), utils.WithProject(s.projectID)))
	prefs, err := account.GetPrefs()
	if err != nil {
		return nil, fmt.Errorf("error getting prefs: %v", err)
	}
	var preferences model.Prefs
	if err := prefs.Decode(&preferences); err != nil {
		return nil, fmt.Errorf("failed to decode prefs: %v", err)
	}
	return &model.Prefs{
		Tumblr:   preferences.Tumblr,
		Twitter:  preferences.Twitter,
		AvatarID: preferences.AvatarID,
		Facebook: preferences.Facebook,
	}, nil
}

type Config struct {
	endpoint  string
	projectID string
}

type Option func(*Config)

func WithEndpoint(endpoint string) Option {
	return func(c *Config) {
		c.endpoint = endpoint
	}
}

func WithProject(projectID string) Option {
	return func(c *Config) {
		c.projectID = projectID
	}
}

func (s *session) UpdateName(secret, name string) (*model.Account, error) {
	if secret == "" {
		return nil, fmt.Errorf("secret cannot be empty")
	}
	account := appwrite.NewAccount(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.endpoint), utils.WithProject(s.projectID)))
	user, err := account.UpdateName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to update username %w", err)
	}
	return model.NewAccount(user), nil
}

func (s *session) UpdateEmail(secret, email, password string) (*model.Account, error) {
	if secret == "" {
		return nil, fmt.Errorf("secret cannot be empty")
	}
	account := appwrite.NewAccount(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.endpoint), utils.WithProject(s.projectID)))
	user, err := account.UpdateEmail(email, password)
	if err != nil {
		return nil, fmt.Errorf("failed to update email %w", err)
	}
	return model.NewAccount(user), nil
}

func (s *session) UpdatePassword(secret, oldPassword, newPassword string) (*model.Account, error) {
	if secret == "" {
		return nil, fmt.Errorf("secret cannot be empty")
	}
	account := appwrite.NewAccount(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.endpoint), utils.WithProject(s.projectID)))
	user, err := account.UpdatePassword(newPassword, account.WithUpdatePasswordOldPassword(oldPassword))
	if err != nil {
		return nil, fmt.Errorf("failed to update password %w", err)
	}
	return model.NewAccount(user), nil
}

func NewSessionWithConfig(config *config.Config) Session {
	return &session{
		endpoint:  config.Appwrite.Endpoint,
		projectID: config.Appwrite.ProjectID,
	}
}

func NewSession(options ...Option) Session {
	cfg := &Config{
		endpoint:  "https://fra.cloud.appwrite.io/v1",
		projectID: "6510a59f633f9d57fba2",
	}
	for _, option := range options {
		option(cfg)
	}
	return &session{
		endpoint:  cfg.endpoint,
		projectID: cfg.projectID,
	}
}

func (s *session) DeleteCurrentSession(secret string) error {
	if secret == "" {
		return fmt.Errorf("secret cannot be empty")
	}
	account := appwrite.NewAccount(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.endpoint), utils.WithProject(s.projectID)))

	_, err := account.DeleteSession("current")
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}
	return nil
}

func (s *session) GetAccount(secret string) (*model.Account, error) {
	if secret != "" {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("X-Appwrite-Session", secret).
			SetHeader("X-Appwrite-Project", s.projectID).
			SetHeader("X-Appwrite-Response-Format", "1.6.0").
			SetResult(&models.User{}).
			Get(s.endpoint + "/account")
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve account information: %w", err)
		}
		data := resp.Result().(*models.User)
		return model.NewAccount(data), nil
	}
	return nil, fmt.Errorf("secret is empty")
}

func (s *session) UpdatePreferences(secret string, prefs *model.Prefs) (*model.Account, error) {
	request := map[string]interface{}{"prefs": *prefs}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Appwrite-Session", secret).
		SetHeader("X-Appwrite-Project", s.projectID).
		SetHeader("X-Appwrite-Response-Format", "1.6.0").
		SetBody(request).
		SetResult(&models.User{}).
		Patch(s.endpoint + "/account/prefs")
	if err != nil {
		return nil, fmt.Errorf("failed to update preferences: %w", err)
	}
	data := resp.Result().(*models.User)
	return model.NewAccount(data), nil
}

func (s *session) CreateVerification(secret string, verificationUrl string) (*model.Token, error) {
	if secret == "" {
		return nil, fmt.Errorf("secret cannot be empty")
	}
	account := appwrite.NewAccount(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.endpoint), utils.WithProject(s.projectID)))
	token, err := account.CreateVerification(verificationUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create verification link: %w", err)
	}
	return &model.Token{Token: token}, nil
}

func (s *session) VerifyAccount(secret, userId string) (*model.Token, error) {
	account := appwrite.NewAccount(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.endpoint), utils.WithProject(s.projectID)))
	token, err := account.UpdateVerification(userId, secret)
	if err != nil {
		return nil, fmt.Errorf("failed to verify account %s: %w", userId, err)
	}
	return &model.Token{Token: token}, nil
	//TODO: increment number of registered user on admin backend
}
