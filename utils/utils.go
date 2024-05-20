package utils

import (
	"errors"
	"fmt"
	"gin-temp-app/database"
	"gin-temp-app/types"
	"log"
	// "net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) (*types.User, error) {
	secretKey := os.Getenv("SecretKey")
	tokenString, err := c.Cookie("auth_token")
	fmt.Println(tokenString, "auth")
	if(err != nil){
		return nil,errors.New("Unauthenticted User.")
	}
	fmt.Println(tokenString)
    log.Println(len(tokenString))
    
	if tokenString == "" {
		log.Println("Empty token string")
		return nil, errors.New("missing authorization header")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
    log.Println(token,"token")
    
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("token expired or invalid")
	}
	log.Println(token,err)

	claims := token.Claims.(jwt.MapClaims)
	log.Println(claims)
	email := claims["email"].(string)
	log.Println(email,"email")
	user := database.Mgr.GetSingleRecordByUsername(email,"user")
	return user, nil
}

func CreateToken(c *gin.Context, email string) (string,error) {
	secretKey := os.Getenv("SecretKey")
	claims := jwt.MapClaims{
        "email": email, // Example claim
        "exp":   fmt.Sprintf("%s",time.Now().Add(time.Hour * 24 * 7).Unix()) , // Expiration time (adjust as needed)
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    signedToken,err := token.SignedString([]byte(secretKey))
    log.Println(signedToken,"token")
	// http.SetCookie(c.Writer, &http.Cookie{
    //     Name:     "auth_token",
    //     Value:    signedToken,
    //     Expires:  time.Now().Add(time.Hour * 24 * 7), 
    //     HttpOnly: false, 
    //     Secure:   false, 
    // })
    c.SetCookie("auth_token",signedToken, 10000, "/","localhost",false,false)
	return signedToken,err
}