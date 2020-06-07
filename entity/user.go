package entity

// User holds user meta
type User struct {
	ID       uint64
	Username string
	Name     string
	Email    string
	Role     Role
}

// Role defines user's role
type Role string

// role types
const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)
