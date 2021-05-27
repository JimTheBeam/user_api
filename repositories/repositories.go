package repositories

// UserRepo is a store for users
type UserRepo interface {
	CreateUser()
	GetAllUsers()
	GetUserById()
	UpdateUser()
	DeleteUser()
}
