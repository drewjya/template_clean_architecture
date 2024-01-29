package service

import "template_clean_architecture/app/module/auth/repository"

type authService struct {
	Repo repository.AuthRepository
}

type AuthService interface {
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		Repo: repo,
	}
}
