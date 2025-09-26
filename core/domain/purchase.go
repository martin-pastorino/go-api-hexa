package domain

type Purchase struct {
	ID       string
	CreateAt string
	Items    []Product
	UserID   string
	CartID   string
}
