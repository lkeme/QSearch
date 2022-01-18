package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	_jwt "github.com/golang-jwt/jwt"
	"strings"
	"time"
)

var (
	ErrTokenExpired           = errors.New("token is expired")
	ErrTokenExpiredMaxRefresh = errors.New("token has passed maximum refresh time")
	ErrTokenNotValidYet       = errors.New("token not active yet")
	ErrTokenMalformed         = errors.New("that's not even a token")
	ErrTokenInvalid           = errors.New("couldn't handle this token")
	ErrHeaderEmpty            = errors.New("token is required to access")
	ErrHeaderMalformed        = errors.New("token format in the request header is incorrect")
)

type JWT struct {
	// encryption key
	SignKey []byte
	// maximum expiration time
	MaxRefreshTime time.Duration
}

// JWTCustomClaims 自定义载荷
type JWTCustomClaims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`

	// StandardClaims 结构体实现了 Claims 接口继承了  Valid() 方法
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	_jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:        []byte(""),
		MaxRefreshTime: time.Duration(1) * time.Minute,
	}
}

func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {
	tokenString, err := jwt.getTokenFromHeader(c)
	// parse string err
	if err != nil {
		return nil, err
	}
	//
	return jwt.parseToken(tokenString)
}

func (jwt *JWT) parseToken(tokenString string) (*JWTCustomClaims, error) {
	token, err := _jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *_jwt.Token) (i interface{}, e error) {
		return jwt.SignKey, nil
	})
	if err != nil {
		if vErr, ok := err.(*_jwt.ValidationError); ok {
			if vErr.Errors == _jwt.ValidationErrorMalformed {
				// Format error
				return nil, ErrTokenMalformed
			} else if vErr.Errors == _jwt.ValidationErrorExpired {
				// Expired
				return nil, ErrTokenExpired
			} else if vErr.Errors == _jwt.ValidationErrorNotValidYet {
				// Not active
				return nil, ErrTokenNotValidYet
			} else {
				// Invalid
				return nil, ErrTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	//c.GetHeader("Authorization")
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}
