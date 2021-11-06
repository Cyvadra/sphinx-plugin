package main

import (
	"fmt"
	"github.com/Cyvadra/sphinx-service/pkg/grpc"
	"github.com/NpoolPlatform/sphinx-service/message/npool"
	"github.com/myxtype/filecoin-client"
)

var host string

func init() {
	fmt.Println("hello")
}

func lotusRpcUrl(host string) string {
	return fmt.Sprintf("http://%v:1234/rpc/v0", host)
}

func SetHost(str string) {
	host = lotusRpcUrl(str)
}

func WalletBalance(address string) (int64, error) {
	bh, err := grpc.Request(lotusRpcUrl(host), []string{address}, "Filecoin.WalletBalance")
	if err != nil {
		log.Errorf(log.Fields{}, "lotusapi request error: %v", err)
		return -1, err
	}
	head := types.TipSet{}
	err = json.Unmarshal(bh, &head)
	if err != nil {
		log.Errorf(log.Fields{}, "cannot unmarshal chain head resp, err is %v", err)
		return -2, err
	}

	return int64(head.Height()), nil
}

func WalletBalance(address string)
