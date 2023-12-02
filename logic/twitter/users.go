package twitter

import (
	"context"
	"errors"
	"github.com/g8rswimmer/go-twitter/v2"
)

func GetUserFollowersAndListedCount(ctx context.Context, username string) (*UserFollowersAndListedCountResult, error) {
	// todo
	resp, err := GetClient().AuthUserLookup(ctx, twitter.UserLookupOpts{
		UserFields: []twitter.UserField{twitter.UserFieldPublicMetrics},
	})
	if err != nil {
		return nil, err
	}

	userItemSet := resp.Raw.UserDictionaries()
	if len(userItemSet) == 0 {
		return nil, errors.New("user not found")
	}

	result := new(UserFollowersAndListedCountResult)
	for _, userItem := range userItemSet {
		result.FollowersCount = uint(userItem.User.PublicMetrics.Followers)
		result.ListedCount = uint(userItem.User.PublicMetrics.Listed)
		break
	}

	return result, nil
}
