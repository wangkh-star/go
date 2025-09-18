package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 创建带超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 连接Infura
	client, err := ethclient.DialContext(ctx, "https://sepolia.infura.io/v3/89363631a8ef455ca1b37f251c8c5698")
	if err != nil {
		log.Fatalf("连接Infura失败: %v", err)
	}
	defer client.Close()

	fmt.Println("✅ 成功连接到以太坊测试网")

	// 验证连接
	if err := verifyConnection(ctx, client); err != nil {
		log.Fatalf("连接验证失败: %v", err)
	}

	// 查询账户信息
	account := common.HexToAddress("0x881A6aF6f7871298Cf9BcF8397DA6D13834eA79d")
	if err := getAccountBalance(ctx, client, account); err != nil {
		log.Fatalf("查询账户信息失败: %v", err)
	}
}

func verifyConnection(ctx context.Context, client *ethclient.Client) error {
	// 获取区块高度来验证连接
	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("获取区块高度失败: %w", err)
	}
	fmt.Printf("📦 当前区块高度: %d\n", blockNumber)
	return nil
}

func getAccountBalance(ctx context.Context, client *ethclient.Client, account common.Address) error {
	fmt.Printf("👤 查询账户: %s\n", account.Hex())

	// 获取余额
	balance, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		return fmt.Errorf("查询余额失败: %w", err)
	}

	// 将wei转换为ETH
	ethBalance := new(big.Float)
	ethBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(ethBalance, big.NewFloat(1e18))

	fmt.Printf("💰 余额(wei): %s\n", balance.String())
	fmt.Printf("💰 余额(ETH): %.8f\n", ethValue)

	return nil
}
