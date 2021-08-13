package userStorage

type IUserStorage interface {
	CreateUser()
	RemoveUser()
	UpdateUser()
	DeleteUser()
}
