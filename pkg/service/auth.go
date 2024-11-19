package service

import (
	todo "ToDoApp"
	"ToDoApp/pkg/repository"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "akjdowiqnfd210392oqkwndq"
	tokenTTL   = 12 * time.Hour
	signingKey = "kJSHU*&A^*&*WHBIDB(*!(()))"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json: "user_id"`
}

// реализация интерфейса Authorization
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// Хэширование пароля и передача на уровень ниже - в базу
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUser(username, password string) (todo.User, error) {
	return s.repo.GetUser(username, password)
}

// func (s *AuthService) GetUser(user todo.User) (int, error) {

// }

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
