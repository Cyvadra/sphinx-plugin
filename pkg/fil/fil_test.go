package fil

import (
	"fmt"
	"testing"
)

func TestGetBalance(t *testing.T) {
	SetHostWithToken("172.16.30.117", "")
	balance, err := GetBalance("t3wkzab6teh4b3wkw2kdhnazynlzjrv2wqwaylbmj6w33lpsb6lry6a4yuwqif2s7nm3mohx6773gk6zsj27kq")
	if err != nil {
		fmt.Println("err is", err)
		return
	}
	fmt.Println("test account balance:", balance)
}
