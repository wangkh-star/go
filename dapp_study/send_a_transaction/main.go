package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

	//获取当前高度
	blockNumber, err := client.BlockNumber(ctx)

	if err != nil {
		log.Fatal("获取当前区块高度失败：%w", err)
	}
	fmt.Println("获取当前高度为：", blockNumber)

	//查询当前账户余额   0x8Dd252DD1C00Cb7723B973a042d026149E3a3d10   0x881A6aF6f7871298Cf9BcF8397DA6D13834eA79d
	selectAddrBlance(ctx, client)

	//发送交易
	//1.获取私钥
	privateKey, err := crypto.HexToECDSA("1574d79ec45aead5f7675fcfc1403ba9af0648c787c91f06139e60c84379e3ed")
	if err != nil {
		log.Fatal("获取密钥失败")
	}
	//获取发送账户的公共地址
	publickey := privateKey.Public()
	publickeyECdsa, ok := publickey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("publickey  is not  ecdsa.PublicKey  ")
	}
	fromAddress := crypto.PubkeyToAddress(*publickeyECdsa)
	//获取随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("获取随机数失败")
	}
	//设置要发送的数量
	value := big.NewInt(1000000000000000000)
	//gas
	gaslimit := uint64(21000)

	gasprice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		log.Fatal("获得平均燃气价格。")
	}

	//将eth发送给谁
	toAddress := common.HexToAddress("0x881A6aF6f7871298Cf9BcF8397DA6D13834eA79d")
	//
	var data []byte

	tx := types.NewTransaction(nonce, toAddress, value, gaslimit, gasprice, data)
	//使用发件人的私钥对事务进行签名
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("签名失败")
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)

	if err != nil {
		log.Fatal("SignTx fail")
	}

	//SendTransaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("SendTransaction fail")
	}
	fmt.Printf("tx sent %s", signedTx.Hash().Hex())

	// 等待交易被挖出
	fmt.Println("等待交易确认...")
	receipt, err := waitForTransactionReceipt(client, signedTx.Hash())
	if err != nil {
		log.Fatal("等待交易时出错: ", err)
	}
	fmt.Printf("交易已在区块 %d 确认\n", receipt.BlockNumber.Uint64())

	// 现在再查询余额，结果一定是更新后的
	fmt.Println("第二次查询账户余额（交易确认后）")

	selectAddrBlance(ctx, client)
}

func selectAddrBlance(ctx context.Context, client *ethclient.Client) {
	account := common.HexToAddress("0x8Dd252DD1C00Cb7723B973a042d026149E3a3d10")
	fmt.Println("查询当前账户为：", account.Hex())
	//获取余额
	balance, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		log.Fatal("获取余额失败")
	}
	ethBalance := new(big.Float)
	ethBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(ethBalance, big.NewFloat(1e18))
	fmt.Println("余额（wei）:", balance.String())
	fmt.Printf("余额(ETH): %.8f\n", ethValue)
}
func waitForTransactionReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			// 如果错误不是“未找到”，则返回错误
			if err.Error() != "not found" {
				return nil, err
			}
			// 交易还未被打包，等待一会儿再重试
			time.Sleep(2 * time.Second) // 每2秒检查一次
			continue
		}
		// 找到了收据，说明交易已被打包
		return receipt, nil
	}
}
