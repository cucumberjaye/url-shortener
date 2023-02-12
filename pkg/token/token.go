package token

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"github.com/cucumberjaye/url-shortener/configs"
)

func GenerateNewToken(id string) (string, error) {
	key := md5.Sum([]byte(configs.SigningKey))

	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	enc := aesgcm.Seal(nil, key[:12], []byte(id), nil)

	return hex.EncodeToString(enc), nil
}

func CheckToken(token string) (string, error) {
	data, err := hex.DecodeString(token)
	if err != nil {
		return "", err
	}

	key := md5.Sum([]byte(configs.SigningKey))

	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	dec, err := aesgcm.Open(nil, key[:12], data, nil)
	if err != nil {
		return "", err
	}

	return string(dec), nil
}
