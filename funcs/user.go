package funcs

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser(name string) *User {
	return &User{Name: name, Email: "jhodndoe@example.com"}
}
