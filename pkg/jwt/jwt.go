package jwt

import (
	"crypto/rsa"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type CustomClaims struct {
	Sub string `json:"sub"`
	jwt.RegisteredClaims
}

type Payload struct {
	Sub string `json:"sub"`
}

func New() *JWT {
	pub, priv := getKeys()
	return &JWT{
		privateKey: priv,
		publicKey:  pub,
	}
}

func (j *JWT) VerifyToken(rawToken string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(rawToken, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return j.publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return claims, nil
}

func (j *JWT) GenerateToken(payload Payload) (string, error) {
	claims := &CustomClaims{
		Sub: payload.Sub,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(j.privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signedToken, nil
}

func getKeys() (*rsa.PublicKey, *rsa.PrivateKey) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("failed to get current working directory: %w", err))
	}

	pubKeyData, err := os.ReadFile(path.Join(cwd, "keys", "public.pem"))
	if err != nil {
		panic(fmt.Errorf("failed to read public key: %w", err))
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyData)
	if err != nil {
		panic(fmt.Errorf("failed to parse public key: %w", err))
	}

	privKeyData, err := os.ReadFile(path.Join(cwd, "keys", "private.pem"))
	if err != nil {
		panic(fmt.Errorf("failed to read private key: %w", err))
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privKeyData)
	if err != nil {
		panic(fmt.Errorf("failed to parse private key: %w", err))
	}

	return pubKey, privKey
}
