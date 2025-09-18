package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接到Sepolia测试网（需要替换为你的Alchemy或Infura URL）
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

	// 替换为你的测试网私钥（不要使用主网私钥！）
	privateKey, err := crypto.HexToECDSA("784283fd1cb8d8bcf9741e6a2a373e4b31525074240395a3073e83e1c7e5cc6c")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // 0 ETH
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x8Dd252DD1C00Cb7723B973a042d026149E3a3d10")                            // 替换为接收者地址
	tokenAddress := common.HexToAddress("0x4d4300ad8d213a4f34171e4be59f5d2f2c994f99c851eab2963e3a535717cd63") // 替换为ERC20代币合约地址

	// 构造transfer函数调用数据
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000个代币

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// 估算Gas
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress, // 注意：这里应该是代币合约地址
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 创建并签名交易
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("交易已发送: %s\n", signedTx.Hash().Hex())
}
