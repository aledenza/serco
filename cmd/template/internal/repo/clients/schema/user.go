package clientSchema

type User struct {
	UserId     int    `json:"user_id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}
