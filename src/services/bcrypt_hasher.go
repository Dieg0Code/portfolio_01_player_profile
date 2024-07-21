package services

type PasswordHasher interface {
	HashPassword(password string) (string, error)
}
