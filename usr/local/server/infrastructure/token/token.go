package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/config"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/errors"
)

type Token struct {
	expire time.Duration
	secret string
	issuer string
}

func New(config *config.Config) (*Token, error) {
	conf := config.Token
	if expire, err := time.ParseDuration(conf.Expire); err != nil {
		return nil, errors.New(err, "parse auth expire failure")
	} else {
		return &Token{issuer: conf.Issuer, secret: conf.Secret, expire: expire}, nil
	}
}
func (s *Token) Signed(info Info) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		info, jwt.RegisteredClaims{Issuer: s.issuer, ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expire))},
	}).SignedString([]byte(s.secret))
}
func (s *Token) Verify(token string) (Info, error) {
	claims := &claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) { return []byte(s.secret), nil })
	return claims.Info, err
}

type Info struct {
	Username string
}
type claims struct {
	Info
	jwt.RegisteredClaims
}
