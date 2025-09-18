package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接Infura（建议使用环境变量配置URL）
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/89363631a8ef455ca1b37f251c8c5698")
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer client.Close()

	fmt.Println("✅ 成功连接到以太坊测试网")

	// 创建带取消的上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 设置信号监听，优雅退出
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(ctx, headers)
	if err != nil {
		log.Fatalf("订阅失败: %v", err)
	}
	defer sub.Unsubscribe()

	fmt.Println("🚀 开始监听新区块...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("订阅错误: %v", err)
			return

		case header := <-headers:
			go processBlock(client, header) // 并发处理区块

		case <-sigCh:
			fmt.Println("\n🛑 接收到终止信号，正在退出...")
			cancel()
			time.Sleep(1 * time.Second) // 等待清理
			return

		case <-ctx.Done():
			fmt.Println("上下文已取消，退出程序")
			return
		}
	}
}

func processBlock(client *ethclient.Client, header *types.Header) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	block, err := client.BlockByNumber(ctx, header.Number)
	if err != nil {
		log.Printf("获取区块失败: %v", err)
		return
	}

	// 输出区块信息
	fmt.Printf("\n=== 新区块 ===\n")
	fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
	fmt.Printf("区块高度: %d\n", block.Number().Uint64())
	fmt.Printf("区块时间: %s\n", time.Unix(int64(block.Time()), 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("交易数量: %d\n", len(block.Transactions()))
	fmt.Printf("处理耗时: %v\n", time.Since(start))

	// 可选：处理交易详情
	if len(block.Transactions()) > 0 {
		fmt.Println("第一笔交易哈希:", block.Transactions()[0].Hash().Hex())
	}
}
