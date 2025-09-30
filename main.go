package main

import (
	"fmt"

	"github.com/afauble/FollowTracker/logic"
)

const FOLDERPATH string = "testData"
const FOLLOWERPATH string = "/followers_1.json"
const FOLLOWINGPATH string = "/following.json"

func main() {

	var notFollowingYouBack = make(map[string]string)
	var extraFollower = make(map[string]string)

	// Parse json files
	followerMap := logic.ParseFollowerJson(FOLDERPATH + FOLLOWERPATH)
	followingMap := logic.ParseFollowingJson(FOLDERPATH + FOLLOWINGPATH)

	numberOfFollowers := len(followerMap)
	numberYourFollowing := len(followingMap)

	// Loop through followers to find people you don't following back
	for key, value := range followerMap {
		_, exists := followingMap[key]
		if exists {
			delete(followingMap, key)
		} else {
			extraFollower[key] = value.Href
		}
	}
	// Left over followings are people who don't follow back
	for key, value := range followingMap {
		notFollowingYouBack[key] = value.Href
	}

	fmt.Println("Number of followers: ", numberOfFollowers)
	fmt.Println("Number following: ", numberYourFollowing)
	fmt.Println("People not following you back")
	for key, value := range notFollowingYouBack {
		fmt.Println("\t", key, "\t", value)
	}

	fmt.Println("People you don't follow back")
	for key, value := range extraFollower {
		fmt.Println("\t", key, "\t", value)
	}

}
