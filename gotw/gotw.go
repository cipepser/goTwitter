package gotw

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

// yaml fileからアクセストークンやシークレットを読み取るためのstruct
type Keys struct {
	TwitterConsumerKey       string `yaml:"TwitterConsumerKey"`
	TwitterConsumerSecret    string `yaml:"TwitterConsumerSecret"`
	TwitterAccessToken       string `yaml:"TwitterAccessToken"`
	TwitterAccessTokenSecret string `yaml:"TwitterAccessTokenSecret"`
}

// FollowbySupportAcount
// keywordで検索してでてきたユーザのフォロワーをn人チェックしてフォローする
// チェック内容は以下。
// * 未フォロー
// * descriptionが空欄でない
// * フォロワー数が100以上
// * NWワードがdescriptionかな前に含まれている
func FollowbySupportAcount(keyword string, api *anaconda.TwitterApi, n int, ngs []string) error {
	supportAcounts, _ := api.GetUserSearch(keyword, nil)

	if len(supportAcounts) == 0 {
		return errors.New("support account is not found.")
	}

	rand.Seed(time.Now().UnixNano())
	chosenAccount := rand.Intn(len(supportAcounts))

	c, err := api.GetFollowersUser(supportAcounts[chosenAccount].Id, nil)
	if err != nil {
		return err
	}

	cnt := 0
	for i := 0; i < int(math.Min(float64(n), float64(len(c.Ids)))); i++ {
		u, err := api.GetUsersShowById(c.Ids[i], nil)
		if err != nil {
			return err
		}
		fmt.Print(u.Name, ": ")

		flgNG := false
		for _, ng := range ngs {
			if strings.Contains(u.Description, ng) || strings.Contains(u.Name, ng) {
				flgNG = true
				break
			}
		}
		if flgNG {
			fmt.Println("I do not follow.")
			continue
		}

		if !u.Following && u.Description != "" && u.FollowersCount > 100 {
			cnt++
			fmt.Println("just followed!")
			api.FollowUserId(u.Id, nil)

		} else {
			fmt.Println("I do not follow.")
		}

	}

	fmt.Println("----------------------")
	fmt.Println("result: follow " + strconv.Itoa(cnt) + "/" + strconv.Itoa(n) + " accounts")
	return nil
}

// SearchandFollow
// keywordで探したユーザをフォローしていなければフォローする
// フォロー済みのユーザも出てしまうので微妙。
func SearchandFollow(keyword string, api *anaconda.TwitterApi, v url.Values) error {
	users, err := api.GetUserSearch(keyword, v)
	if err != nil {
		return err
	}

	for _, u := range users {
		// fmt.Println(string(u.Id) + ":" + u.Description)
		// fmt.Println(string(u.Id) + ":" + strconv.FormatBool(u.Following))
		// fmt.Println("-------------------")
		// fmt.Println(strconv.FormatInt(u.Id, 10) + ":" + strconv.FormatBool(u.Following))
		// fmt.Println(u.ScreenName)
		// fmt.Println(u.Name)
		// fmt.Println(u.Description)
		// fmt.Println(u.Entities)

		if !u.Following {
			api.FollowUserId(u.Id, nil)
		}
	}

	return nil
}

func UnfollowNotEachOther(api *anaconda.TwitterApi) error {
	c, err := api.GetFriendsIds(nil)
	if err != nil {
		return err
	}

	u, err := api.GetSelf(nil)
	if err != nil {
		return err
	}

	myId := u.IdStr

	for _, id := range c.Ids {
		v := url.Values{}
		v.Set("source_id", myId)
		v.Set("target_id", strconv.FormatInt(id, 10))

		r, err := api.GetFriendshipsShow(v)
		if err != nil {
			return err
		}

		fmt.Println("-----------")
		fmt.Println(r.Relationship.Target.Screen_name)
		fmt.Println(r.Relationship.Target.Following)

		// リフォローしてくれてない
		if !r.Relationship.Target.Following {
			api.UnfollowUserId(id)
		}
	}

	return nil
}
