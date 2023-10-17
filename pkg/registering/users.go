package registering

type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Username    string `json:"user_name"`
	Email       string `json:"email"`
}
