package util

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"math/big"
	"regexp"

	"github.com/xdg-go/pbkdf2"
)

func HashString(val string) (string, error) {
	salt, err := generateRandomSalt(128 / 8)
	if err != nil {
		return "", err
	}

	digest := pbkdf2.Key([]byte(val), salt, 1000, 256/8, sha512.New)

	saltStr := base64.StdEncoding.EncodeToString(salt)
	digestStr := base64.StdEncoding.EncodeToString(digest)
	result := "$" + saltStr + "$" + digestStr
	return result, nil
}

func VerifyDigest(val string, digest string) bool {
	sp := regexp.MustCompile(`\$`)
	splitted := sp.Split(digest, -1)

	if len(splitted) != 3 {
		return false
	}

	salt, err := base64.StdEncoding.DecodeString(splitted[1])
	if err != nil {
		return false
	}

	inputHashed := pbkdf2.Key([]byte(val), salt, 1000, 256/8, sha512.New)
	inputHashedStr := base64.StdEncoding.EncodeToString(inputHashed)

	return inputHashedStr == splitted[2]
}

func generateRandomSalt(length int) ([]byte, error) {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		salt, err := rand.Int(rand.Reader, big.NewInt(255))
		if err != nil {
			return nil, err
		}

		result[i] = byte(salt.Int64())
	}
	return result, nil
}
