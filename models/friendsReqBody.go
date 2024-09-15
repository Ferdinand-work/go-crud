package models

type FriendsReqBody struct {
	Friends []string `json:"friends,omitempty"`
	Friend  string   `json:"friend,omitempty"`
}
