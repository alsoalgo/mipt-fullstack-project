package auth

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	bcrypt "golang.org/x/crypto/bcrypt"

	httpmodels "travelgo/internal/models/http"
	tokenrepo "travelgo/internal/repository/token"
	userrepo "travelgo/internal/repository/user"
)

var jwtKey, _ = os.LookupEnv("JWT_KEY")

const (
	ErrUserAlreadyExists = "user already exists"
	ErrUserNotFound      = "user not found"
	ErrInvalidPassword   = "invalid password"
)

type Service struct {
	token *tokenrepo.Repository
	repo  *userrepo.Repository

	secret string
}

func New(repo *userrepo.Repository, token *tokenrepo.Repository, secret string) *Service {
	return &Service{
		repo:   repo,
		token:  token,
		secret: secret,
	}
}

func (s *Service) Login(ctx context.Context, req *httpmodels.LoginRequest) (string, error) {
	ctx = context.WithValue(ctx, "opName", "Login")

	exists, err := s.repo.ExistsUser(ctx, req.Email)
	if err != nil {
		return "", errors.New("ExistsUser: " + err.Error())
	}

	if !exists {
		return "", errors.New(ErrUserNotFound)
	}

	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("GetUserByEmail: " + err.Error())
	}

	errHash := compareHashPassword(req.Password, user.PasswordHash)
	if !errHash {
		return "", errors.New(ErrInvalidPassword)
	}

	userId := strconv.FormatInt(user.ID, 10)
	claims := &httpmodels.Claims{
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			Subject: userId,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", errors.New("SignedString: " + err.Error())
	}

	_, err = s.token.AddToken(ctx, user.ID, tokenString)
	if err != nil {
		return "", errors.New("AddToken: " + err.Error())
	}

	return tokenString, nil
}

func (s *Service) Register(ctx context.Context, req *httpmodels.RegisterRequest) (int64, error) {
	ctx = context.WithValue(ctx, "opName", "Register")

	exists, err := s.repo.ExistsUser(ctx, req.Email)
	if err != nil {
		return 0, errors.New("ExistsUser: " + err.Error())
	}

	if exists {
		return 0, errors.New(ErrUserAlreadyExists)
	}

	hash, err := generateHashPassword(req.Password)
	if err != nil {
		return 0, err
	}

	return s.repo.CreateUser(ctx, req.Email, hash)
}

func (s *Service) CheckToken(ctx context.Context, req *httpmodels.CheckTokenRequest) (bool, error) {
	ctx = context.WithValue(ctx, "opName", "CheckToken")

	exists, err := s.token.IsTokenExpired(ctx, req.Token)
	if err != nil {
		return false, errors.New("IsTokenExpired: " + err.Error())
	}

	return exists, nil
}

func compareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ParseToken(req *http.Request) (int64, error) {
	cookies := req.Cookies()
	tokenString := ""
	for _, cookie := range cookies {
		if cookie.Name == "token" {
			tokenString = cookie.Value
		}
	}

	token, err := jwt.ParseWithClaims(tokenString, &httpmodels.Claims{}, func(token *jwt.Token) (any, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return 0, errors.New("ParseWithClaims: " + err.Error())
	}

	claims, ok := token.Claims.(*httpmodels.Claims)
	if !ok {
		return 0, errors.New("cast to *httpmodels.Claims")
	}

	userId, err := strconv.ParseInt(claims.StandardClaims.Subject, 10, 64)
	if err != nil {
		return 0, errors.New("PaParseIntrseWithClaims: " + err.Error())
	}

	return userId, nil
}
