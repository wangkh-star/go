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

	// 查询账户信息   0x881A6aF6f7871298Cf9BcF8397DA6D13834eA79d /0x8Dd252DD1C00Cb7723B973a042d026149E3a3d10
	account := common.HexToAddress("0x881A6aF6f7871298Cf9BcF8397DA6D13834eA79d")
	if err := getAccountBalance(ctx, client, account); err != nil {
		log.Fatalf("查询账户信息失败: %v", err)
	}

	//获取私钥   0x881A6aF6f7871298Cf9BcF8397DA6D13834eA79d  784283fd1cb8d8bcf9741e6a2a373e4b31525074240395a3073e83e1c7e5cc6c
	//         	0x8Dd252DD1C00Cb7723B973a042d026149E3a3d10  1574d79ec45aead5f7675fcfc1403ba9af0648c787c91f06139e60c84379e3ed
	privateKey, err := crypto.HexToECDSA("784283fd1cb8d8bcf9741e6a2a373e4b31525074240395a3073e83e1c7e5cc6c")
	if err != nil {
		log.Fatal("获取私钥失败", err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//发送的帐户的公共地址 - 这个我们可以从私钥派生。
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//读取我们应该用于帐户交易的随机数。
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	//设置我们将要转移的 ETH 数量。 但是我们必须将 ETH 以太转换为 wei，因为这是以太坊区块链所使用的。 以太网支持最多 18 个小数位，因此 1 个 ETH 为 1 加 18 个零。
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	//ETH 转账的燃气应设上限为“21000”单位。
	gasLimit := uint64(21000) // in units
	//燃气价格总是根据市场需求和用户愿意支付的价格而波动的，因此对燃气价格进行硬编码有时并不理想。 go-ethereum 客户端提供 SuggestGasPrice 函数，用于根据'x'个先前块来获得平均燃气价格。
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//将 ETH 发送给谁。
	toAddress := common.HexToAddress("0x8Dd252DD1C00Cb7723B973a042d026149E3a3d10")

	//现在我们最终可以通过导入 go-ethereum core/types 包并调用 NewTransaction 来生成我们的未签名以太坊事务，这个函数需要接收 nonce，地址，值，燃气上限值，燃气价格和可选发的数据。
	//发送 ETH 的数据字段为“nil”。 在与智能合约进行交互时，我们将使用数据字段，仅仅转账以太币是不需要数据字段的。
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	//下一步是使用发件人的私钥对事务进行签名。 为此，我们调用 SignTx 方法，该方法接受一个未签名的事务和我们之前构造的私钥。 SignTx 方法需要 EIP155 签名者，这个也需要我们先从客户端拿到链 ID。
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//现在通过在 client 实例调用 SendTransaction 来将已签名的事务广播到整个网络。

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
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
