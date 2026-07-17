package models

type UserRole string

const (
	RoleBuyer  UserRole = "buyer"
	RoleSeller UserRole = "seller"
	RoleBoth   UserRole = "both"
)

type User struct {
	ID           string
	FullName     string
	Email        string
	Phone        string
	City         string
	State        string
	Country      string
	PasswordHash string
	Role         UserRole
}