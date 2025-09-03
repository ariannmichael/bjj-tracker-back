package domain_user

type UserRepository interface {
	Create(user *User) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(user *User) (*User, error)
	FindByID(id string) (*User, error)
	FindAll() ([]User, error)
}
