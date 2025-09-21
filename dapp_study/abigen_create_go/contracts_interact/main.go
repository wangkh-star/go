package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractAddr = "0x15b8628600aCD8359437F67d08E294d64F773477"
)

func main() {
	// 连接到以太坊测试网
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/89363631a8ef455ca1b37f251c8c5698")
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer client.Close()

	fmt.Println("成功连接到以太坊测试网")

	// 加载私钥
	privateKey, err := crypto.HexToECDSA("784283fd1cb8d8bcf9741e6a2a373e4b31525074240395a3073e83e1c7e5cc6c")
	if err != nil {
		log.Fatal("私钥加载失败:", err)
	}

	// 获取公钥地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("公钥类型错误")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("操作账户: %s\n", fromAddress.Hex())

	// 检查账户余额
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal("获取余额失败:", err)
	}
	fmt.Printf("账户余额: %s ETH\n", weiToEther(balance))

	// 1. 首先进行投票操作
	fmt.Println("\n1. 执行投票操作...")
	err = vote(client, privateKey, fromAddress, "0x881A6aF6f7871298Cf9BcF8397DA6D13834eA79d")
	if err != nil {
		log.Fatal("投票失败:", err)
	}

	// 2. 查询投票状态
	fmt.Println("\n2. 查询投票状态...")
	hasVoted, err := getHasVotes(client, fromAddress, fromAddress)
	if err != nil {
		log.Fatal("查询投票状态失败:", err)
	}
	fmt.Printf("地址 %s 是否已投票: %t\n", fromAddress.Hex(), hasVoted)

	// 3. 查询投票数量
	fmt.Println("\n3. 查询投票数量...")
	voteCount, err := getVotes(client, fromAddress)
	if err != nil {
		log.Fatal("查询投票数量失败:", err)
	}
	fmt.Printf("地址 %s 获得的票数: %d\n", fromAddress.Hex(), voteCount)

	// 4. 查询目标地址的投票数量
	targetAddress := common.HexToAddress("0x881A6aF6f7871298Cf9BcF8397DA6D13834eA79d")
	targetVoteCount, err := getVotes(client, targetAddress)
	if err != nil {
		log.Fatal("查询目标地址投票数量失败:", err)
	}
	fmt.Printf("目标地址 %s 获得的票数: %d\n", targetAddress.Hex(), targetVoteCount)
}

// 投票函数
func vote(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address, voteTo string) error {
	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("获取nonce失败: %v", err)
	}

	// 获取 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("获取gas价格失败: %v", err)
	}

	// 准备合约 ABI
	contractABI, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"address","name":"voteTo","type":"address"}],"name":"vote","outputs":[],"stateMutability":"nonpayable","type":"function"}]`))
	if err != nil {
		return fmt.Errorf("解析ABI失败: %v", err)
	}

	// 准备调用数据
	voteToAddress := common.HexToAddress(voteTo)
	input, err := contractABI.Pack("vote", voteToAddress)
	if err != nil {
		return fmt.Errorf("打包数据失败: %v", err)
	}

	// 创建并签名交易
	chainID := big.NewInt(11155111) // Sepolia 链ID
	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddr), big.NewInt(0), 200000, gasPrice, input)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return fmt.Errorf("签名交易失败: %v", err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("发送交易失败: %v", err)
	}

	fmt.Printf("投票交易已发送: %s\n", signedTx.Hash().Hex())

	// 等待交易确认
	receipt, err := waitForReceipt(client, signedTx.Hash())
	if err != nil {
		return fmt.Errorf("等待交易确认失败: %v", err)
	}

	if receipt.Status == 1 {
		fmt.Println("投票成功!")
	} else {
		return fmt.Errorf("投票失败，交易被回滚")
	}

	return nil
}

// 查询是否已投票
func getHasVotes(client *ethclient.Client, fromAddress common.Address, queryAddress common.Address) (bool, error) {
	contractABI, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"address","name":"voteTo","type":"address"}],"name":"getHasVotes","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		return false, err
	}

	input, err := contractABI.Pack("getHasVotes", queryAddress)
	if err != nil {
		return false, err
	}

	to := common.HexToAddress(contractAddr)
	callMsg := ethereum.CallMsg{
		From: fromAddress,
		To:   &to,
		Data: input,
	}

	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return false, err
	}

	var hasVoted bool
	err = contractABI.UnpackIntoInterface(&hasVoted, "getHasVotes", result)
	return hasVoted, err
}

// 查询投票数量
func getVotes(client *ethclient.Client, queryAddress common.Address) (*big.Int, error) {
	contractABI, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"address","name":"voteTo","type":"address"}],"name":"getVotes","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		return nil, err
	}

	input, err := contractABI.Pack("getVotes", queryAddress)
	if err != nil {
		return nil, err
	}

	to := common.HexToAddress(contractAddr)
	callMsg := ethereum.CallMsg{
		To:   &to,
		Data: input,
	}

	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, err
	}

	var voteCount *big.Int
	err = contractABI.UnpackIntoInterface(&voteCount, "getVotes", result)
	return voteCount, err
}

// 等待交易确认
func waitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("等待交易确认超时")
		default:
			receipt, err := client.TransactionReceipt(ctx, txHash)
			if err == nil {
				return receipt, nil
			}
			if err != ethereum.NotFound {
				return nil, err
			}
			time.Sleep(2 * time.Second)
		}
	}
}

// 将 wei 转换为 ETH
func weiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
}
