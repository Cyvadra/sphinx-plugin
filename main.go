package main

import (
	"fmt"
	"github.com/Cyvadra/sphinx-plugin/pkg/filecoin"
)

func main() {
	filecoin.SetHost("10.155.8.32")
	fmt.Println(filecoin.WalletBalance("f0121"))
}
