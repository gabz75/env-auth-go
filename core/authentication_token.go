package core

import (
    "time"
    "fmt"
    "regexp"
    "crypto/rsa"
    "io/ioutil"

    jwt "github.com/dgrijalva/jwt-go"
)

const (
    privKeyPath = "/Users/Gabriel/workspace/go/src/github.com/gabz75/auth-api/config/keys/app.rsa"     // openssl genrsa -out app.rsa keysize
    pubKeyPath  = "/Users/Gabriel/workspace/go/src/github.com/gabz75/auth-api/config/keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
    verifyKey *rsa.PublicKey
    signKey   *rsa.PrivateKey
)

func fatal(err error) {
    if err != nil {
        fmt.Println(err)
    }
}

func init() {
    signBytes, err := ioutil.ReadFile(privKeyPath)
    fatal(err)

    signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
    fatal(err)

    verifyBytes, err := ioutil.ReadFile(pubKeyPath)
    fatal(err)

    verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
    fatal(err)
}

// GenerateToken -
func GenerateToken() string {
    token := jwt.New(jwt.GetSigningMethod("RS256"))
    token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
    tokenString, err := token.SignedString(signKey)
    fatal(err)

    return tokenString
}

// ExtractToken -
func ExtractToken(authorizationHeader string) string {
    regex, _ := regexp.Compile("Bearer (.{233})")
    submatch := regex.FindStringSubmatch(authorizationHeader)
    if len(submatch) == 2 {
        return submatch[1]
    }
    return ""
}
