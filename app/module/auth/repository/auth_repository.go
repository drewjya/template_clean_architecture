package repository

type authRepository struct {
}

type AuthRepository interface {
}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}
