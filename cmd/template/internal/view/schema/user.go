package schema

type User struct {
	UserId     int    `json:"user_id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

type UserResponse struct {
	Body User
}

type UserRequest struct {
	UserId int `path:"user_id"`
}
