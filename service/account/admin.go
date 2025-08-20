package account

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"github.com/appwrite/sdk-for-go/account"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/users"
)

type Admin interface {
	SignIn(email, password string) (*model.Session, error)
	GetUser(userId string) (*model.User, error)
	SignUp(username, email, password string) (*model.Account, error)
	PasswordReset(email string, recoveryUrl string) (*model.Token, error)
	UpdateVerification(secret, userId string) (*model.Token, error)
}

type admin struct {
	client *client.Client
}

func NewAdmin(client *client.Client) Admin {
	return &admin{
		client: client,
	}
}

func (s *admin) SignIn(email, password string) (*model.Session, error) {
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}
	if password == "" {
		return nil, fmt.Errorf("password cannot be empty")
	}
	accountDB := account.New(*s.client)
	session, err := accountDB.CreateEmailPasswordSession(email, password)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}
	return &model.Session{Session: session}, nil
}

func (s *admin) GetUser(userId string) (*model.User, error) {
	if userId == "" {
		return nil, fmt.Errorf("userId cannot be empty")
	}
	usersDB := users.New(*s.client)
	user, err := usersDB.Get(userId)
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
	usersDB := users.New(*s.client)
	user, err := usersDB.Create(
		id.Unique(),
		usersDB.WithCreateEmail(email),
		usersDB.WithCreatePassword(password),
		usersDB.WithCreateName(username))
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
	accountDB := account.New(*s.client)
	token, err := accountDB.CreateRecovery(email, recoveryUrl)
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
	accountDB := account.New(*s.client)
	token, err := accountDB.UpdateVerification(userId, secret)
	if err != nil {
		return nil, fmt.Errorf("failed to update verification: %w", err)
	}
	return &model.Token{Token: token}, nil
}
