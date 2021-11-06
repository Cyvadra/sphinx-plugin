package main

import (
	"fmt"
	"github.com/Cyvadra/sphinx-plugin/pkg/fil"
)

func main() {
	fil.SetHostWithToken("172.16.30.117", "")
	fmt.Println(fil.GetBalance("t3wkzab6teh4b3wkw2kdhnazynlzjrv2wqwaylbmj6w33lpsb6lry6a4yuwqif2s7nm3mohx6773gk6zsj27kq"))
}
