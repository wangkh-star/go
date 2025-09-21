package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	//创建一个带超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//连接测试网络
	client, err := ethclient.DialContext(ctx, "https://sepolia.infura.io/v3/89363631a8ef455ca1b37f251c8c5698")
	if err != nil {
		log.Fatal("连接失败直接结束: %v", err)
	}
	defer client.Close()
	//连接成功
	fmt.Println("连接测试网成功")

	//查询指定区块数据
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	//区块号
	fmt.Println(block.NumberU64())
	//区块时间戳
	fmt.Println(block.Time())
	//区块难度
	fmt.Println(block.Hash().Hex())
	//交易列表
	fmt.Println(block.Transactions())

}
