package main

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	anaconda.SetConsumerKey(os.Getenv("TwitterConsumerKey"))
	anaconda.SetConsumerSecret(os.Getenv("TwitterConsumerSecret"))
	api := anaconda.NewTwitterApi(os.Getenv("TwitterAccessToken"), os.Getenv("TwitterAccessTokenSecret"))

	// フォロー
	// api.FollowUser("cipepser")
	// api.FollowUserId(492787827, nil)

	// フォロー解除
	// api.UnfollowUserId(492787827)

	// ユーザ検索
	// users, _ := api.GetUsersLookup("cipepser", nil)
	// for _, u := range users {
	// 	fmt.Println(u.Id)
	// }

	// 検索してフォロー
	// v := url.Values{}
	// v.Set("screen_name", "相互")
	// SearchandFollow("相互 マンガ", api, v)

}
