package main

var users = []User{
	User{
		ID:       1,
		Username: "Hannah",
		Password: "1",
	},
	User{
		ID:       2,
		Username: "Olivia",
		Password: "2",
	},
	User{
		ID:       3,
		Username: "Amelia",
		Password: "3",
	},
}

type AuthService struct{}

func NewAuthService() AuthService {
	return AuthService{}
}

func (this AuthService) Login(username, password string) (User, bool) {
	for _, u := range users {
		if u.Username == username && u.Password == password {
			return u, true
		}
	}
	return User{}, false
}
