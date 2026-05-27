package user

func FindUser(users []User, email string) *User {
	for i := range users {
		if users[i].Email == email {
			return &users[i]
		}
	}
	return nil
}

func AddUser(User []User, newUser User) []User {
	User = append(User, newUser)
	return User
}

func RemoveUser(a User, b []User) []User {
	for i, user := range b {
		if a.Name == user.Name {
			b = append(b[:i], b[i+1:]...)
		}
	}
	return b
}
