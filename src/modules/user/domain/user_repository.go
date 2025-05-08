package domain_user

type UserRepository interface {
	Create(user *User) (*User, error)
	FindByEmail(email string) (*User, error)
}
