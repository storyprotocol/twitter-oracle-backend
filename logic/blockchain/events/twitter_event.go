package events

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"sync/atomic"
	"twitter-oracle-backend/internal/log"
	"twitter-oracle-backend/logic/blockchain/actions"
	"twitter-oracle-backend/logic/blockchain/eth/contracts/twitter"
	twitterClient "twitter-oracle-backend/logic/twitter"
)

type TwitterEvent struct {
	ev *Events

	Addr       common.Address
	BlockStart int64

	isStart atomic.Bool
	isClose atomic.Bool
	ctx     context.Context
	cancel  context.CancelFunc

	sol *twitter.Twitter
}

func NewTwitterEvent(ev *Events, addr string) (e *TwitterEvent, err error) {
	e = &TwitterEvent{
		ev:         ev,
		Addr:       common.HexToAddress(addr),
		BlockStart: -1,
	}

	e.ctx, e.cancel = context.WithCancel(context.Background())
	e.sol, err = twitter.NewTwitter(e.Addr, e.ev.ethClient.Client())
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (e *TwitterEvent) Start(rt ratelimit.Limiter) {
	if e.isStart.CompareAndSwap(false, true) {
		twitterChan := make(chan *twitter.TwitterTwitterFollowerNumRequest, 100)
		go func() {
			defer close(twitterChan)

			e.ev.filterRange(e.ctx, rt, e.BlockStart, func(watchStart, watchEnd uint64) error {
				fmt.Printf("watchStart: %d\n", watchStart)

				twitterEvents, err := e.sol.FilterTwitterFollowerNumRequest(&bind.FilterOpts{Start: watchStart, End: &watchEnd}, nil, nil)
				if err != nil {
					return err
				}
				defer twitterEvents.Close()

				for twitterEvents.Next() {
					if twitterEvents.Error() != nil {
						return twitterEvents.Error()
					}
					twitterChan <- twitterEvents.Event
				}

				return nil
			})
		}()

		go func() {
			e.processTransfers(twitterChan)
			e.isClose.Store(true)
		}()
	}
}

func (e *TwitterEvent) processTransfers(transferChan chan *twitter.TwitterTwitterFollowerNumRequest) {
	for transfer := range transferChan {
		err := e.processTwitter(transfer)
		if err != nil {
			log.GetLogger().Warn("[TwitterTwitterFollowerNumRequest]", zap.Any("event", transfer), zap.Error(err))
		}
	}
}

func (e *TwitterEvent) processTwitter(ev *twitter.TwitterTwitterFollowerNumRequest) error {
	result, err := twitterClient.GetUserFollowersAndListedCount(context.Background(), ev.Username)
	if err != nil {
		return err
	}

	log.GetLogger().Info("get twitter user public metrics succeed.", zap.Uint("followersCount", result.FollowersCount), zap.Uint("ListedCount", result.ListedCount))

	// push oracle data
	txHash, err := actions.PushTwitterUserFollowersAndListedCount(ev.CallbackAddr, ev.RequestId, result.FollowersCount, result.ListedCount)
	if err != nil {
		return err
	}

	log.GetLogger().Info("push oracle data succeed.", zap.String("txHash", txHash))
	return nil
}
