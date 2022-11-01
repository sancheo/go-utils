package user

import (
	"fmt"
	"testing"
	"time"
)

func TestGetUserID(t *testing.T) {
	for {
		id := GetUserID()
		time.Sleep(1 * time.Second)
		fmt.Printf("generate user id = %v \n", id)
	}
}
