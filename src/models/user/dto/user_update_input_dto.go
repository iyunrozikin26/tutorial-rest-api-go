package user

// untuk membuat cetakan object input user
type UpdateUserInput struct {
	Name  string `json:"name" valiidate:"required"`
	Email string `json:"email" valiidate:"email"`
}
