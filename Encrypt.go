package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
)

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(b string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(b)
}

func encryptData(data string) (string, error) {

	Secret, err := getEnvVar("S3CRET")
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

func decryptData(eData string) (string, error) {

	Secret, err := getEnvVar("S3CRET")
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
