package auth

var users = []User{
	{
		ID:       1,
		Username: "Hannah",
		Password: "1",
	},
	{
		ID:       2,
		Username: "Olivia",
		Password: "2",
	},
	{
		ID:       3,
		Username: "Amelia",
		Password: "3",
	},
}

type AuthService struct{}

func NewAuthService() AuthService {
	return AuthService{}
}

func (AuthService) Login(username, password string) (User, bool) {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return user, true
		}
	}

	return User{}, false
}
