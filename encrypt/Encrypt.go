package encrypt

import (
	envFunc "AccountCreationService/envFuncs"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
)

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(b string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(b)
}

func EncryptData(data string) (string, error) {

	Secret, err := envFunc.GetEnvVar("S3CRET")
	if err != nil {
		log.Fatal("Error: %s", err)
	}
	block, err := aes.NewCipher([]byte(Secret))
	if err != nil {
		return "", err
	}

	plainText := []byte(data)

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return Encode(append(iv, cipherText...)), nil
}

func DecryptData(eData string) (string, error) {

	Secret, err := envFunc.GetEnvVar("S3CRET")
	if err != nil {
		log.Fatal("Error: %s", err)
	}

	DeData, err := Decode(eData)

	if err != nil {
		return "", err
	}

	if len(DeData) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := DeData[:aes.BlockSize]
	DeData = DeData[aes.BlockSize:]

	block, err := aes.NewCipher([]byte(Secret))
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(DeData))
	cfb.XORKeyStream(plainText, DeData)

	return string(plainText), nil

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//func VerifyEmail(password) bool {
//
//}
