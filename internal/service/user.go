package service

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
)

func (s *Service) ListUser(ctx context.Context) ([]entities.User, error) {

	return s.io.FilterUsers(ctx)
}

func (s *Service) Login(ctx context.Context, user models.User) (*models.JWT, error) {
	userEntity := &entities.User{
		Username: user.Username,
		Password: user.Password,
	}

	if err := s.io.LoginUser(ctx, userEntity); err != nil {
		return nil, err
	}

	token, err := createJWT(s.config.JWT, user.Username)
	if err != nil {
		return nil, err
	}

	jwt := &models.JWT{Token: token}

	return jwt, nil
}

func createJWT(config conf.SecuriyConfiguration, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Duration(config.EXPIRATION) * time.Minute).Unix()
	claims["username"] = username

	tokenStr, err := token.SignedString([]byte(config.SECRET))

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}
