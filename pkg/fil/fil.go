package filecoin

import (
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
	respBytes, err := grpc.Request(host, []string{address}, "Filecoin.WalletBalance")
	if err != nil {
		log.Errorf(log.Fields{}, "lotusapi request error, %v", err)
		return "", err
	}
	return string(respBytes), err
}
