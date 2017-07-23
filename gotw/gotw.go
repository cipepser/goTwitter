package gotw

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

// FollowbySupportAcount
// keywordで検索してでてきたユーザのフォロワーをn人チェックしてフォローする
// チェック内容は以下。
// * 未フォロー
// * descriptionが空欄でない
// * フォロワー数が100以上
// * NWワードがdescriptionかな前に含まれている
func FollowbySupportAcount(keyword string, api *anaconda.TwitterApi, n int, ngs []string) {
	supportAcounts, _ := api.GetUserSearch(keyword, nil)

	if len(supportAcounts) == 0 {
		return
	}

	c, _ := api.GetFollowersUser(supportAcounts[0].Id, nil)

	for i := 0; i < int(math.Min(float64(n), float64(len(c.Ids)))); i++ {
		u, _ := api.GetUsersShowById(c.Ids[i], nil)
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

			fmt.Println("just followed!")
			api.FollowUserId(u.Id, nil)

		} else {
			fmt.Println("I do not follow.")
		}

	}

}

// SearchandFollow
// keywordで探したユーザをフォローしていなければフォローする
// フォロー済みのユーザも出てしまうので微妙。
func SearchandFollow(keyword string, api *anaconda.TwitterApi, v url.Values) {
	users, _ := api.GetUserSearch(keyword, v)
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
}

func UnfollowNotEachOther(api *anaconda.TwitterApi) {
	c, _ := api.GetFriendsIds(nil)

	u, _ := api.GetSelf(nil)
	myId := u.IdStr

	for _, id := range c.Ids {
		v := url.Values{}
		v.Set("source_id", myId)
		v.Set("target_id", strconv.FormatInt(id, 10))

		r, _ := api.GetFriendshipsShow(v)

		fmt.Println("-----------")
		fmt.Println(r.Relationship.Target.Screen_name)
		fmt.Println(r.Relationship.Target.Following)

		// リフォローしてくれてない
		if !r.Relationship.Target.Following {
			api.UnfollowUserId(id)
		}

	}
}
