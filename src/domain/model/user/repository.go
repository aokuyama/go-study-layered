package user

type UserRepository interface {
	LoadByID(id UserID) (*User, error)
}
