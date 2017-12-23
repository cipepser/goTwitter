package main

import (
	"bufio"
	"io"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/ChimeraCoder/anaconda"

	"./gotw"
)

func main() {
	f, err := os.Open("./secret.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	k := gotw.Keys{}
	r := bufio.NewReader(f)
	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(l, &k)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}

	anaconda.SetConsumerKey(k.TwitterConsumerKey)
	anaconda.SetConsumerSecret(k.TwitterConsumerSecret)
	api := anaconda.NewTwitterApi(k.TwitterAccessToken, k.TwitterAccessTokenSecret)

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
