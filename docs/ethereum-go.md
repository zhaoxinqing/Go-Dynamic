
#

使用以太坊创建去中心化应用（即"dapps"）

Ethereum-以斯锐姆

Goerli-戈尔利

abigen - A-B-I-gen

# `Seaport` 协议

Seaport 核心智能合约是开源的，本质上是去中心化的，没有合约所有者、具有可升级性以及其他特性。现在，OpenSea 是一个市场平台（类似于 Facebook）。该平台提供让你买卖 NFT 的服务。
OpenSea 公司 100% 控制平台，因此对代码的任何更改（如接受 APE）都是由 OpenSea 一人“当家作主”（中心化）。

# 以太坊

## 创建有效以太坊交易所需的授权数据的集合 `TransactOpts`

```go
type TransactOpts struct {
 From   common.Address // 发送交易的以太坊账户
 Nonce  *big.Int       // 用于交易执行（nil = 使用挂起状态）
 Signer SignerFn       // 用于签署交易的方法（必填）

 Value     *big.Int // 随交易转账的资金（nil = 0 = 没有资金）
 GasPrice  *big.Int // 用于交易执行的 gas 价格（nil = gas price oracle）
 GasFeeCap *big.Int // 用于 1559 交易执行的 Gas 上限 (nil = gas price oracle)
 GasTipCap *big.Int // 用于 1559 交易执行的 Gas 优先费用上限 (nil = gas price oracle)
 GasLimit  uint64   // 为交易执行设置的 Gas 限制 (0 = estimate)

 Context context.Context // 支持取消和超时的网络上下文 (nil = no timeout)

 NoSend bool // 执行所有交易步骤但不发送交易
}
```
