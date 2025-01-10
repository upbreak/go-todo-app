package entity

type User struct {
	UserId  string `json:"userId" db:"USER_ID"`
	UserPwd string `json:"userPwd" db:"USER_PWD"`
}
