package sub

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func InitContract() (*bind.TransactOpts, *Contract) {
	client, err := initClient()
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	auth, err := initTransactor()
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	simpleDex, err := initSimpleDex(client)
	if err != nil {
		log.Fatalf("Failed to initialize SimpleDex contract: %v", err)
	}
	return auth, simpleDex
}
