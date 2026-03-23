package entities

type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUserDTO(id int, username, email string) *UserDTO {
	return &UserDTO{
		ID:       id,
		Username: username,
		Email:    email,
	}
}

func (u *UserDTO) GetID() int {
	return u.ID
}

func (u *UserDTO) GetUsername() string {
	return u.Username
}
