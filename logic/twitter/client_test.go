package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/g8rswimmer/go-twitter/v2"
	"log"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	user, err := GetClient().UserNameLookup(context.Background(), []string{"test"}, twitter.UserLookupOpts{
		UserFields: []twitter.UserField{twitter.UserFieldPublicMetrics},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("user: %+v\n", user)
}

func TestGetMyInfo(t *testing.T) {
	user, err := GetClient().AuthUserLookup(context.Background(), twitter.UserLookupOpts{
		UserFields: []twitter.UserField{twitter.UserFieldPublicMetrics},
	})
	if err != nil {
		panic(err)
	}

	dictionaries := user.Raw.UserDictionaries()

	enc, err := json.MarshalIndent(dictionaries, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(enc))
}
