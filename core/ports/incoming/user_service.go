package incoming

type UserService interface {
	CreateUser(name, email string) (string, error)
}
