package cms

import (
	"errors"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

var PostInfoCache []PostInfo
var PostCache []Post

func refreshCaches() {
	log.Println("Refreshing Caches...")
	log.Println("Getting posts...")
	PostInfoCache = getAllPostInfo()
	PostCache = getAllPosts()

	log.Println("Loaded", len(PostCache), "posts")
	log.Println("Getting posts OK")

	log.Println("Sorting post cards...")
	sortPosts() // sort posts from most recent to least recent
	log.Println("Sorting post cards OK")
	log.Println("Refreshing Caches OK")
}

// / interval is in hours
func PeriodicallyRefresh(interval int) {
	go func() {
		for {
			refreshCaches()
			time.Sleep(time.Duration(time.Hour * time.Duration(interval)))
		}
	}()
}

func GetPost(name string) (Post, error) {
	for _, post := range PostCache {
		if post.Location == name {
			return post, nil
		}
	}

	return Post{}, errors.New("unluko, no post")
}

func sortPosts() {
	sort.Slice(PostInfoCache, func(i, j int) bool {
		dateI := strings.Split(PostInfoCache[i].Info.Date, "-")
		dateJ := strings.Split(PostInfoCache[j].Info.Date, "-")

		yearI, _ := strconv.Atoi(dateI[0])
		yearJ, _ := strconv.Atoi(dateJ[0])
		monthI, _ := strconv.Atoi(dateI[1])
		monthJ, _ := strconv.Atoi(dateJ[1])
		dayI, _ := strconv.Atoi(dateI[2])
		dayJ, _ := strconv.Atoi(dateJ[2])

		if yearI == yearJ {
			if monthI == monthJ {
				return dayI > dayJ
			} else {
				return monthI > monthJ
			}
		} else {
			return yearI > yearJ
		}
	})
}
