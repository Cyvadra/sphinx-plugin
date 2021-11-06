package main

import (
	"fmt"
	"github.com/Cyvadra/sphinx-plugin/pkg/fil"
)

func main() {
	fil.SetHostWithToken("172.16.30.117")
	fmt.Println(fil.GetBalance("f0121"))
}
