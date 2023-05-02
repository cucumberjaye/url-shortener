package token

import (
	"fmt"

	"github.com/google/uuid"
)

func ExampleGenerateNewToken() {
	id := uuid.New().String()

	token, err := GenerateNewToken(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(token)
}

func ExampleCheckToken() {
	id := uuid.New().String()

	token, err := GenerateNewToken(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	decID, err := CheckToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("start id: %s, decrypted id: %s\n", id, decID)
}
