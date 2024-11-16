package service

import (
	todo "ToDoApp"
	"ToDoApp/pkg/repository"
	"crypto/sha1"
	"fmt"
)
const salt = "akjdowiqnfd210392oqkwndq"

// реализация интерфейса Authorization
type AuthService struct {
	repo repository.Authorization
}

//Хэширование пароля и передача на уровень ниже - в базу
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}


func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (*AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}