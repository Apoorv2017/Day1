package Middleware
import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("super-secret-auth-key")
var api_key1 = "1234"
var api_key2 = "5678"

func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenStr, err := token.SignedString(SECRET)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return tokenStr, nil
}
func ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header["Token"]==nil{
			c.String(http.StatusUnauthorized, "Message: Enter the Token.")
			c.Abort()
			return 
		}else{
			_, err := jwt.Parse(c.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method")
				}else{
					return SECRET, nil
				}
				})
			if err!=nil{
				c.String(http.StatusUnauthorized, "message: Enter the correct Token.")
				c.Abort()
				return
			}else{
				c.Next()
			}
		}
	}
}
func GetJwtClient(c *gin.Context) {
	if c.Request.Header["Access"] != nil {
		if c.Request.Header["Access"][0] != api_key1 {
			c.String(http.StatusUnauthorized, "message: Wrong Access Key.")
			c.Abort()
			return
		} else {
			token, err := CreateJWT()
			if err != nil {
				c.String(http.StatusUnauthorized, "Token not created")
				c.Abort()
				return
			}
			c.JSON(http.StatusOK,  gin.H{"token":token})
			}
	}else{
		c.String(http.StatusUnauthorized, "message:Enter the Access Key.")
		c.Abort()
	}
}
func GetJwtRetailer(c *gin.Context) {
	if c.Request.Header["Access"] != nil {
		if c.Request.Header["Access"][0] != api_key2 {
			c.String(http.StatusUnauthorized, "message: Wrong Access Key.")
			c.Abort()
			return
		} else {
			token, err := CreateJWT()
			if err != nil {
				c.String(http.StatusUnauthorized, "Token not created")
				c.Abort()
				return
			}
			c.JSON(http.StatusOK,  gin.H{"token":token})
		}
	}else{
		c.String(http.StatusUnauthorized, "message:Enter the Access Key.")
		c.Abort()
	}
}