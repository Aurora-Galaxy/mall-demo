package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 生成token签名的密钥
var jwtSecret = []byte("LLL")

type Claims struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

type EmailClaims struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operationType"`
	jwt.StandardClaims
}

func GenerateToken(id uint, userName string, authority int) (string, error) {
	timeNow := time.Now()
	//token过期时间
	expireTime := timeNow.Add(24 * time.Hour)
	claims := Claims{
		ID:        id,
		Username:  userName,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-mall", //指定token的颁发者
		},
	}
	//创建一个新的jwt-token对象，使用hs256签名算法，claims为token主体
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//根据给出的jwtSecret生成token并返回
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 验证token
func ParseToken(token string) (*Claims, error) {
	//token：待解析的 JWT 令牌字符串。
	//&Claims{}：一个 Claims 类型的空实例，用于在解析过程中将令牌中的声明解码为结构体。
	//一个回调函数：这个函数接受一个 *jwt.Token 对象作为参数，返回用于验证签名的密钥。在这里，回调函数直接返回了 jwtSecret 作为用于验证签名的密钥。
	//tokenClaims 将包含解析后的令牌信息
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	// tokenClaims 非空代表解析成功
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims) //类型断言，判断解析后是否为定义的*Claims类型
		//判断token的有效性
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func GenerateEmailToken(id uint, email string, operationType uint, password string) (string, error) {
	timeNow := time.Now()
	//token过期时间
	expireTime := timeNow.Add(24 * time.Hour)
	claims := EmailClaims{
		ID:            id,
		Email:         email,
		OperationType: operationType,
		Password:      password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-mall-email", //指定token的颁发者
		},
	}
	//创建一个新的jwt-token对象，使用hs256签名算法，claims为token主体
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//根据给出的jwtSecret生成token并返回
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseEmailToken  验证token
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	// tokenClaims 非空代表解析成功
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*EmailClaims) //类型断言，判断解析后是否为定义的*Claims类型
		//判断token的有效性
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
