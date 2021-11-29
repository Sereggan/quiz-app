package service

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/Sereggan/quiz-app/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	salt            = "hjqrhjqw124617ajfhajs"
	signingKey      = "qrkjk#4#%35FSFJlja#4353KSFjH"
	accessTokenTTL  = 2 * time.Hour
	refreshTokenTTL = 48 * time.Hour
	autoLogOutTime  = 0
)

type JwtCustomClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthService struct {
	userRepository repository.User
	tokenClient    repository.TokenClient
}

func (a *AuthService) CreateUser(context context.Context, user *model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return a.userRepository.Create(user)
}

func (a *AuthService) GenerateToken(context context.Context, username, password string) (string, error) {
	user, err := a.userRepository.Find(username, generatePasswordHash(password))
	if err != nil {
		logrus.Errorf("Failed to find user, username: %s, err: %s", username, err.Error())
		return "", err
	}

	accessToken, err := createToken(user.Id, accessTokenTTL)
	refreshToken, err := createToken(user.Id, refreshTokenTTL)

	tokens, err := json.Marshal(Tokens{
		accessToken,
		refreshToken,
	})

	err = a.tokenClient.Set(fmt.Sprintf("token-%d", user.Id), tokens, autoLogOutTime)

	if err != nil {
		logrus.Errorf("Failed to save Token for user, username: %s, err: %s", username, err.Error())
		return "", err
	}

	return string(tokens), nil
}

func (a *AuthService) ParseToken(context context.Context, accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logrus.Errorf("Failed to parse token: %s", token)
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		logrus.Errorf("Failed to validate token: %s", token)

		return claims.UserId, nil
	}
	return 0, errors.New("invalid token claims")
}

func (a *AuthService) LogOut(context context.Context, usedId int) error {
	logrus.Printf("Logging out as user id: %d", usedId)
	return a.tokenClient.Delete(strconv.Itoa(usedId))
}

func NewAuthService(repo repository.User, client repository.TokenClient) *AuthService {
	return &AuthService{
		userRepository: repo,
		tokenClient:    client,
	}
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func createToken(userID int, ttl time.Duration) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaims{
		UserId: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}).SignedString([]byte(signingKey))
}
