package filecoin

import (
	"encoding/json"
	"fmt"

	"github.com/Cyvadra/sphinx-plugin/pkg/grpc"
)

var host string

func lotusRpcUrl(host string) string {
	return fmt.Sprintf("http://%v:1234/rpc/v0", host)
}

func SetHost(str string) {
	host = lotusRpcUrl(str)
}

func GetBalance(address string) (string, error) {
	respStr, err := grpc.Request(lotusRpcUrl(host), []string{address}, "Filecoin.WalletBalance")
	if err != nil {
		return "", err
	}
	resl := RespGetBalance{}
	err = json.Unmarshal(respStr, &resl)
	if err != nil {
		return "", err
	}
	return resl.Result, err
}

type RespGetBalance struct {
	Id      int32  `json:"id"`
	JsonRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
}
