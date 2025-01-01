package funcs

type User struct {
    UserName       string `json:"name"`
    Email          string `json:"email"`
    Cvv            string `json:"cvv"`
    Address        string `json:"address"`
    OwnerName      string `json:"owner_name"`
    PatientAddress string `json:"patient_address"`
}

func GetUser(name string) *User {
    user := User{}
    user.UserName = name
    user.Email = "jhodndo@example.com"
    user.Address = "home"
    user.Cvv = "012"
    user.PatientAddress = "home"
    user.OwnerName = "jhodndo"
    return &user
}
