package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/upbreak/go-todo-app/clock"
	"github.com/upbreak/go-todo-app/config"
	"net/http"
	"strings"
	"time"
)

type JWTUtils struct {
	Cfg   *config.JwtConfig
	Clock clock.Clocker
}

type JWTRole string

const (
	Admin JWTRole = "admin"
	User  JWTRole = "user"
)

// claims 정의
type JWTClaims struct {
	UserId string
	Role   JWTRole
	Token  string
}

// JWTUtils 구조체 생성
func JwtNew(c clock.Clocker) (*JWTUtils, error) {
	jwt := &JWTUtils{}

	jwtConfig, err := config.GetJwtConfig()
	if err != nil {
		return nil, err
	}

	jwt.Cfg = jwtConfig
	jwt.Clock = c

	return jwt, nil
}

// 토큰 생성
func (j *JWTUtils) GenerateToken(userId string) (string, error) {
	// 비밀 키
	secretKey := j.Cfg.SecretKey

	// JWT 클레임 설정
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    j.Clock.Now().Add(10 * time.Microsecond), // 만료 시간: 1시간
		"role":   "admin",
	}

	// 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 비밀 키로 서명
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("jwtUtils.go/GenerateToken() err: %w", err)
	}

	return tokenString, nil
}

// 토큰 유효성 검사
func (j *JWTUtils) validateJWT(r *http.Request) (*JWTClaims, error) {
	// Authorization 헤더 읽기
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("Authorization header is empty")
	}

	// "Bearer " 접두사 제거
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return nil, errors.New("Invalid Authorization token format")
	}

	// 비밀 키
	secretKey := j.Cfg.SecretKey

	// 토큰 파싱 및 검증
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 서명 방법 확인
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	// 클레임 확인
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	jwtClaims := &JWTClaims{
		UserId: claims["userId"].(string),
		Role:   JWTRole(claims["role"].(string)),
	}

	return jwtClaims, nil
}
