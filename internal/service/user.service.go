package service

import (
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/internal/repo"
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/pkg/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetInfoUser()
// }

// Implement with interface
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserExist
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
