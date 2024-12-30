package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	StatusCode  int    `json:"status"`
	Message     string `json:"message"`
	CreatedUser User   `json:"created_user"`
}
