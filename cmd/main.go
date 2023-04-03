package main

import (
	"fmt"
	"os"
	"strconv"
)

var wordList []string

func main() {
	for idx, args := range os.Args {
		if idx > 0 {
			fmt.Println("参数"+strconv.Itoa(idx)+":", args)
			wordList = append(wordList, args)
		}

	}
	if len(wordList) != 4 {
		panic("助记词输入错误")
	}
	fmt.Println(wordList)
}

// $go run main.go 1 3 -X ?
// 参数0: /tmp/go-build116558042/command-line-arguments/_obj/exe/main
// 参数1: 1
// 参数2: 3
// 参数3: -X
// 参数4: ?

// // Generate robot accountlist
// func MakeRobotAccountList() bool {
// 	global.Log.Info("====> Generate robot accountlist begin ......")
// 	global.Log.Info("====> please input three mnemonic words")
// 	var inputWords string
// 	for i := 1; i < 4; i++ {
// 		words, _ := utils.GetPasswdPrompt("Please enter the "+strconv.Itoa(i)+" word :", false, os.Stdin, os.Stdout)
// 		inputWords += " " + strings.ToLower(string(words))
// 	}
// 	mnemonic, err := ioutil.ReadFile(global.Conf.RunParams.MnemonicPath)
// 	if err != nil {
// 		global.Log.Error("MakeRobotAccountList Read mnemonic File error: ", err)
// 		return false
// 	}
// 	mnemonicStr := string(mnemonic) + inputWords
// 	seed := bip39.NewSeed(mnemonicStr, "") // config pasword for mnemonic
// 	wallet, err := hdwallet.NewFromSeed(seed)
// 	if err != nil {
// 		global.Log.Error("MakeRobotAccountList hdwallet.NewFromSeed error: ", err)
// 		return false
// 	}

// 	if global.Conf.RunParams.RobotAccountsCount < 1 {
// 		fmt.Println(model.Red, "MakeRobotAccountList RobotAccountsCount is config error", model.Reset)
// 		return false
// 	}

// 	for i := 1; i <= global.Conf.RunParams.RobotAccountsCount; i++ {
// 		paths := "m/44'/60'/0'/0/" + strconv.Itoa(i-1)
// 		path := hdwallet.MustParseDerivationPath(paths)
// 		account, err := wallet.Derive(path, false)
// 		if err != nil {
// 			global.Log.Error("MakeRobotAccountList wallet.Derive error: ", err)
// 			return false
// 		}

// 		address := account.Address.Hex()
// 		privateKey, _ := wallet.PrivateKeyHex(account)

// 		var accounttemp model.Account
// 		accounttemp.Addr = address
// 		accounttemp.PriKey = privateKey
// 		robotAccountList = append(robotAccountList, &accounttemp)
// 		global.Log.Info("====> robot_[%v] address[%v]", i, accounttemp.Addr)
// 	}

// 	return true
// }
