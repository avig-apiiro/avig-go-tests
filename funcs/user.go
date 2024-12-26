package funcs

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Cvv     string `json:"cvv"`
	Address string `json:"address"`
}

func GetUser(name string) *User {
	user := User{}
	user.Name = name
	user.Email = "jhodndo@example.com"
	user.Address = "home"
	user.Cvv = "012"
	return &user
}
