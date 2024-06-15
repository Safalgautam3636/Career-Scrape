package helpers
import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)
var sampleSecretKey []byte

// HELPERS
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateToken(username string) (string, error) {
	godotenv.Load()
	sampleSecretKey = []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(10 * time.Minute).Unix(),
	})

	// SECRET here
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	godotenv.Load()
	sampleSecretKey = []byte(os.Getenv("JWT_KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { return sampleSecretKey, nil })
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}