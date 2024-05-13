package sub

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/m/go/model"
)

var (
	contractAddress = common.HexToAddress("0x48cA7b156fC9EE34A9eE65897821A94268055ee8")
)

// Ethereum 클라이언트 초기화
func initClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(os.Getenv("INFURA_KEY"))
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 트랜잭션 서명을 위한 Transactor 초기화
func initTransactor() (*bind.TransactOpts, error) {
	privateKeyECDSA, err := crypto.HexToECDSA(os.Getenv("WALLET_PRIV_KEY"))
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, big.NewInt(1)) // Chain ID 1 for Ethereum Mainnet
	if err != nil {
		return nil, err
	}
	return auth, nil
}

// SimpleDex
func initSimpleDex(client *ethclient.Client) (*Contract, error) {
	simpleDex, err := NewContract(contractAddress, client)
	if err != nil {
		return nil, err
	}
	return simpleDex, nil
}

// 주문 실행
func executeOrderOnContract(simpleDex *Contract, auth *bind.TransactOpts, buyOrder *model.Order, sellOrder model.Order) error {

	fmt.Println("buy", buyOrder.ID)
	fmt.Println("sell", sellOrder.ID)

	_, err := simpleDex.ExecuteOrder(auth, big.NewInt(int64(buyOrder.ID)), big.NewInt(int64(sellOrder.ID)))
	if err != nil {
		return err
	}
	return nil
}

// 매수 주문 실행 로직
func ExecuteBuyOrders(order *model.Order, sellOrders []model.Order) {
	auth, simpleDex := InitContract()

	for _, sellOrder := range sellOrders {
		if order.Amount <= 0 {
			break
		}

		amountToBuy := min(order.Amount, sellOrder.Amount)

		err := executeOrderOnContract(simpleDex, auth, order, sellOrder)
		if err != nil {
			log.Printf("Failed to execute order: %v", err)
			continue
		}

		order.Amount -= amountToBuy

		log.Printf("Order executed: buyOrderID=%d, sellOrderID=%d, amount=%d", order.ID, sellOrder.ID, amountToBuy)
	}

	if order.Amount > 0 {
		log.Printf("Partially filled or not filled at all: remainingAmount=%d", order.Amount)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
