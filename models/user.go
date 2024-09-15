package models

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Pincode int    `json:"pincode" bson:"pincode"`
}

type User struct {
	Name    string   `json:"name" bson:"user_name"`
	Age     int64    `json:"age" bson:"user_age"`
	Address Address  `json:"address" bson:"user_address"`
	Email   string   `json:"email" bson:"email"`
	Friends []string `json:"friends,omitempty" bson:"friends"`
}
