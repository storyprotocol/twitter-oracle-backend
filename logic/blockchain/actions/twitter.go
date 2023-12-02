package actions

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"twitter-oracle-backend/logic/blockchain/eth"
	"twitter-oracle-backend/logic/blockchain/eth/contracts/twitter"
	"twitter-oracle-backend/logic/config"
)

func PushTwitterUserFollowersAndListedCount(target common.Address, requestId [32]byte, followersCount, listedCount uint) (string, error) {
	twitterContract, err := twitter.NewTwitter(target, eth.DefaultClient.Client())
	if err != nil {
		return "", err
	}

	opts, err := eth.DefaultClient.GetDefaultTransactionOptions(context.Background(), config.Get().Wallet.PrivateKey)
	transaction, err := twitterContract.HandleCallback(opts, requestId, new(big.Int).SetUint64(uint64(followersCount)), new(big.Int).SetUint64(uint64(listedCount)))
	if err != nil {
		return "", err
	}

	return transaction.Hash().String(), nil
}
