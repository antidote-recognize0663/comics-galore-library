package account

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/account"
	"resty.dev/v3"
)

type Session interface {
	DeleteCurrentSession(secret string) error
	GetAccount(secret string) (*model.User, error)
	UpdatePreferences(secret string, prefs *model.Preferences) (*model.User, error)
	CreateVerification(secret string, verificationUrl string) (*model.Token, error)
	VerifyAccount(secret, userId string) (*model.Token, error)
	UpdateName(secret, name string) (*model.Account, error)
	UpdateEmail(secret, email, password string) (*model.Account, error)
	UpdatePassword(secret, oldPassword, newPassword string) (*model.Account, error)
}

type session struct {
	config *config.AppwriteConfig
}

func (s *session) UpdateName(secret, name string) (*model.Account, error) {
	if secret == "" {
		return nil, fmt.Errorf("secret cannot be empty")
	}
	accountDB := account.New(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.config.Endpoint), utils.WithProject(s.config.ProjectID)))
	user, err := accountDB.UpdateName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to update username %w", err)
	}
	return model.NewAccount(user), nil
}

func (s *session) UpdateEmail(secret, email, password string) (*model.Account, error) {
	if secret == "" {
		return nil, fmt.Errorf("secret cannot be empty")
	}
	accountDB := account.New(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.config.Endpoint), utils.WithProject(s.config.ProjectID)))
	user, err := accountDB.UpdateEmail(email, password)
	if err != nil {
		return nil, fmt.Errorf("failed to update email %w", err)
	}
	return model.NewAccount(user), nil
}

func (s *session) UpdatePassword(secret, oldPassword, newPassword string) (*model.Account, error) {
	if secret == "" {
		return nil, fmt.Errorf("secret cannot be empty")
	}
	accountDB := account.New(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.config.Endpoint), utils.WithProject(s.config.ProjectID)))
	user, err := accountDB.UpdatePassword(newPassword, accountDB.WithUpdatePasswordOldPassword(oldPassword))
	if err != nil {
		return nil, fmt.Errorf("failed to update password %w", err)
	}
	return model.NewAccount(user), nil
}

func NewSession(config *config.AppwriteConfig) Session {
	return &session{
		config: config,
	}
}

func (s *session) DeleteCurrentSession(secret string) error {
	if secret == "" {
		return fmt.Errorf("secret cannot be empty")
	}
	accountDB := account.New(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.config.Endpoint), utils.WithProject(s.config.ProjectID)))

	_, err := accountDB.DeleteSession("current")
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}
	return nil
}

func (s *session) GetAccount(secret string) (*model.User, error) {
	if secret != "" {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("X-Appwrite-Session", secret).
			SetHeader("X-Appwrite-Project", s.config.ProjectID).
			SetHeader("X-Appwrite-Response-Format", "1.6.0").
			SetResult(&model.User{}).
			Get(s.config.Endpoint + "/account")
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve account information: %w", err)
		}
		data := resp.Result().(*model.User)
		return data, nil
	}
	return nil, fmt.Errorf("secret is empty")
}

func (s *session) UpdatePreferences(secret string, prefs *model.Preferences) (*model.User, error) {
	request := map[string]interface{}{"prefs": *prefs}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Appwrite-Session", secret).
		SetHeader("X-Appwrite-Project", s.config.ProjectID).
		SetHeader("X-Appwrite-Response-Format", "1.6.0").
		SetBody(request).
		SetResult(&model.User{}).
		Patch(s.config.Endpoint + "/account/prefs")
	if err != nil {
		return nil, fmt.Errorf("failed to update preferences: %w", err)
	}
	data := resp.Result().(*model.User)
	return data, nil
}

func (s *session) CreateVerification(secret string, verificationUrl string) (*model.Token, error) {
	if secret == "" {
		return nil, fmt.Errorf("secret cannot be empty")
	}
	accountDB := account.New(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.config.Endpoint), utils.WithProject(s.config.ProjectID)))
	token, err := accountDB.CreateVerification(verificationUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create verification link: %w", err)
	}
	return &model.Token{Token: token}, nil
}

func (s *session) VerifyAccount(secret, userId string) (*model.Token, error) {
	accountDB := account.New(*utils.NewSessionClient(
		secret, utils.WithEndpoint(s.config.Endpoint), utils.WithProject(s.config.ProjectID)))
	token, err := accountDB.UpdateVerification(userId, secret)
	if err != nil {
		return nil, fmt.Errorf("failed to verify account %s: %w", userId, err)
	}
	return &model.Token{Token: token}, nil
	//TODO: increment number of registered user on admin backend
}
