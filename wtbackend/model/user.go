package model

type User struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Languages       []string `json:"languages"`
	PrimaryLanguage string   `json:"primary_language"`
}
