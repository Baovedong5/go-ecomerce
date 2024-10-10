package service

import "github.com/phuong/go-ecomerce/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInforUserService() string {
	return us.userRepo.GetInforUser()
}
