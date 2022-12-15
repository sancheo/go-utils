package ids

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateUserID(t *testing.T) {
	for {
		id := GenerateID()
		time.Sleep(1 * time.Second)
		fmt.Printf("generate id is = %v \n", id)
	}
}
