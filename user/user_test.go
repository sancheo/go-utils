package user

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateUserID(t *testing.T) {
	for {
		id := GenerateUserID()
		time.Sleep(1 * time.Second)
		fmt.Printf("generate user id = %v \n", id)
	}
}
