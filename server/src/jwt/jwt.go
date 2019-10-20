package jwt

import (
	//"fmt"
	"time"
	"errors"
	"strings"
	"io/ioutil"
	"crypto/rsa"
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
)

//TODO: 有効期限の検討
const (
	EXPIRED_DAY = time.Duration(24)
	ISSUER      = "auth.ai-ot.com" // NOTE: 変更あるかも
)

var (
	SIGN_METHOD = jwt.SigningMethodRS256
)

type JWTKeys struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func ReadKeys() ( *JWTKeys, error ) {
	data, err := ioutil.ReadFile( "./config/keys/public.key" )
	
	if err != nil {
		return nil, err
	}
	
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM( data )
	if err != nil {
		return nil, err
	}

	data, err = ioutil.ReadFile( "./config/keys/secret.key" )
	if err != nil {
		return nil, err
	}
	
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM( data )
	if err != nil {
		return nil, err
	}

	return &JWTKeys{PrivateKey: privateKey, PublicKey: publicKey}, nil
}

func NewToken(keys *JWTKeys, account string, now time.Time) (string, error) {
	if keys == nil || keys.PrivateKey == nil {
		return "", errors.New("private key is not set")
	}
	if account == "" {
		return "", errors.New("account is empty")
	}

	claims := jwt.StandardClaims{
		Issuer:    ISSUER,
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Hour * EXPIRED_DAY).Unix(),
		Subject:   account,
	}

	token := jwt.NewWithClaims(SIGN_METHOD, claims)
	tokenString, err := token.SignedString(keys.PrivateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken( keys *JWTKeys, tokenString string ) (*jwt.Token, error ) {
	if keys == nil || keys.PrivateKey == nil {
		return nil, errors.New("private key is not set")
	}

	parsedToken, err := jwt.Parse( tokenString, func( token *jwt.Token ) ( interface{}, error ) {
		// check signing method
		if _, err := token.Method.(*jwt.SigningMethodRSA); !err {
			return nil, errors.New("unexpected signing method")
		}
		return keys.PublicKey, nil
	})
	
	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, errors.New("token is not valid")
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	if claims["iss"] != ISSUER {
		return nil, errors.New("bad issuer")
	}

	return parsedToken, nil
}

func CheckToken(req http.Request, keys *JWTKeys) (string, error) {
	tokenString := req.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	//fmt.Println( "token" )
	//fmt.Println( tokenString )
	parsedToken, err := ParseToken(keys, tokenString)
	if err != nil {
		if err.Error() == "private key is not set" {
			return "", err
		} else {
			return "", err
		}
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	loginUser := claims["sub"].(string)

	return loginUser, nil
}
