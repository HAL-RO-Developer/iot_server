package model

import (
	"time"

	"crypto/sha256"
	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

/* トークン情報 */
type TokenInfo struct {
	Id    int
	Token string
}

var secretKey = "75c92a074c341e9964329c0550c2673730ed8479c885c43122c90a2843177d5ef21cb50cfadcccb20aeb730487c11e09ee4dbbb02387242ef264e74cbee97213"

/*
	トークンの作成
*/
func CreateTokenString(name string) (string, error) {
	/*
		アルゴリズムの指定
	*/
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"user": name,
		"exp":  time.Now().Add(time.Hour * 48).Unix(),
	}

	/*
	  トークンに対して署名の付与
	*/
	return token.SignedString([]byte(secretKey))
}

/*
	トークンの検証
*/
func AuthorityCheck(c *gin.Context) (string, bool) {
	token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(secretKey)
		return b, nil
	})

	if err == nil {
		claims := token.Claims.(jwt.MapClaims)
		return claims["user"].(string), true
	} else {
		return "", false
	}
}

func ToHash(pass string) string {
	converted := sha256.Sum256([]byte(pass))
	return hex.EncodeToString(converted[:])
}
