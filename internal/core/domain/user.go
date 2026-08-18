package domain

type User struct {
	ID           string
	Version      int
	Name         string
	FullName     string
	phone_number *string
}
