package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/cucumberjaye/url-shortener/configs"
	"strconv"
)

type AuthService struct {
	currentId int
	lastId    int
}

func New() *AuthService {
	return &AuthService{
		currentId: 0,
		lastId:    0,
	}
}

func (s *AuthService) GenerateNewToken() (string, error) {
	key := md5.Sum([]byte(configs.SigningKey))

	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return "", err
	}

	strId := fmt.Sprintf("%d", s.lastId)
	s.currentId = s.lastId
	s.lastId++

	enc := aesgcm.Seal(nil, key[:12], []byte(strId), nil)

	return hex.EncodeToString(enc), nil
}

func (s *AuthService) CheckToken(token string) (int, error) {
	data, err := hex.DecodeString(token)
	if err != nil {
		return 0, err
	}

	key := md5.Sum([]byte(configs.SigningKey))

	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return 0, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return 0, err
	}

	dec, err := aesgcm.Open(nil, key[:12], data, nil)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(string(dec))
}

func (s *AuthService) SetCurrentId(id int) {
	s.currentId = id
}

func (s *AuthService) GetCurrentId() int {
	return s.currentId
}
