package model

import (
	"encoding/json"
	"github.com/appwrite/sdk-for-go/models"
	"log"
)

type Prefs struct {
	AvatarID string `json:"avatar_id"`
	Twitter  string `json:"twitter"`
	Facebook string `json:"facebook"`
	Tumblr   string `json:"tumblr"`
}

func NewPrefs(preferences *models.Preferences) *Prefs {
	prefs, err := json.Marshal(preferences)
	if err != nil {
		log.Println(err)
	}
	var prefsData Prefs
	err = json.Unmarshal(prefs, &prefsData)
	if err != nil {
		log.Println(err)
	}
	return &Prefs{
		AvatarID: prefsData.AvatarID,
		Twitter:  prefsData.Twitter,
		Facebook: prefsData.Facebook,
		Tumblr:   prefsData.Tumblr,
	}
}

type Account struct {
	*models.User
	Prefs *Prefs
}

func NewAccount(user *models.User) *Account {
	var preferences Prefs
	err := user.Prefs.Decode(&preferences)
	if err != nil {
		log.Println(err)
	}
	/*prefData, err := json.Marshal(user.Prefs)
	if err != nil {
		log.Warn(err)
	}
	log.Info(string(prefData))
	err = json.Unmarshal(prefData, &preferences)
	if err != nil {
		log.Warn(err)
	}
	log.Info(preferences)*/

	var account Account
	err = user.Decode(&account)
	if err != nil {
		log.Println(err)
	}
	account.Prefs = &preferences

	return &account
}
