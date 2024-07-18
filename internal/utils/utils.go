package  utils


import(
"time"
"os"
"github.com/dgrijalva/jwt-go"
)


type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID int) (string,error) {
	claims:= &Claims {
		UserID:userID,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt:time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte("SECRET_KEY"))
}

func ParseToken() {
	
}

