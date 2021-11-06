package filecoin

import (
	"encoding/json"
	"fmt"

	"github.com/Cyvadra/sphinx-plugin/pkg/grpc"
	log "github.com/EntropyPool/entropy-logger"
)

var host string

func lotusRpcUrl(host string) string {
	return fmt.Sprintf("http://%v:1234/rpc/v0", host)
}

func SetHost(str string) {
	host = lotusRpcUrl(str)
}

func GetBalance(address string) (string, error) {
	respStr, err := grpc.Request(host, []string{address}, "Filecoin.WalletBalance")
	if err != nil {
		log.Errorf(log.Fields{}, "lotusapi request error, %v", err)
		return "", err
	}
	resl := RespGetBalance{}
	err = json.Unmarshal(respStr, &resl)
	if err != nil {
		log.Errorf(log.Fields{}, "cannot unmarshal chain resp, %v", err)
		return "", err
	}
	return resl.Result, err
}

type RespGetBalance struct {
	Id      int32  `json:"id"`
	JsonRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
}
