package actions

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestPushTwitterUserFollowersAndListedCount(t *testing.T) {
	requestId := [32]byte{144, 125, 81, 106, 158, 94, 205, 160, 141, 151, 184, 141, 122, 142, 161, 33, 194, 25, 139, 118, 113, 221, 147, 82, 187, 0, 198, 149, 194, 59, 184, 9}
	tx, err := PushTwitterUserFollowersAndListedCount(common.HexToAddress("0x25Cfb200eB0e3E7cDe5eBe86bBc1069F1F51cC75"), requestId, 100, 11)
	if err != nil {
		panic(err)
	}

	fmt.Printf("tx: %s\n", tx)
}
