package funcs

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Cvv     string `json:"cvv"`
	Address string `json:"address"`
}

func GetUser(name string) *User {
	return &User{Name: name, Email: "jhodndoe@example.com", Address: "home", Cvv: "012"}
}
