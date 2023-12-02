package twitter

import (
	"context"
	"fmt"
	"testing"
)

func TestGetUserFollowersAndListedCount(t *testing.T) {
	result, err := GetUserFollowersAndListedCount(context.Background(), "test")
	if err != nil {
		return
	}

	fmt.Printf("result: %+v\n", result)
}
