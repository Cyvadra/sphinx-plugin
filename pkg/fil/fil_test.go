package fil

import (
	"fmt"
	"testing"
)

func TestGetBalance(t *testing.T) {
	SetHostWithToken("172.16.30.117", "")
	balance, err := GetBalance("f0121")
	if err != nil {
		fmt.Println("err is", err)
		return
	}
	fmt.Println("f0121's balance is", balance)
}
