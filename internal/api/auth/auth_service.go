package auth

type Service struct {
	users []User
}

func NewService() Service {
	return Service{
		users: []User{
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
		},
	}
}

func (s Service) Login(username, password string) (User, bool) {
	for _, user := range s.users {
		if user.Username == username && user.Password == password {
			return user, true
		}
	}

	return User{}, false
}
