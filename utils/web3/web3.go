package web3

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	genesis "github.com/takez0o/honestwork-api/utils/abi"
)

// move contract addr to config
func FetchUserState(address string) int {
	nft_address := "0x10c0C42A0883699225363341743F719CC87a404e"

	client, err := ethclient.Dial("https://decoded.wtf")
	if err != nil {
		log.Fatal(err)
	}

	nft_address_hex := common.HexToAddress(nft_address)
	instance, err := genesis.NewGenesis(nft_address_hex, client)
	if err != nil {
		log.Fatal(err)
	}

	user_address_hex := common.HexToAddress(address)
	state, err := instance.GetUserState(nil, user_address_hex)
	if err != nil {
		log.Fatal(err)
	}

	return int(state.Int64())
}