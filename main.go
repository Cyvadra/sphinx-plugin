package main

import (
	"fmt"
	"github.com/Cyvadra/sphinx-plugin/pkg/filecoin"
)

func main() {
	filecoin.SetHost("172.16.30.117")
	fmt.Println(filecoin.WalletBalance("f0121"))
}
