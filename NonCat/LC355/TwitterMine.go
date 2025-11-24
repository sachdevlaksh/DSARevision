package LC355

import (
	"slices"
)
//
//type Twitter struct {
//	UserData map[int]Data
//}
//
//type Data struct {
//	TweetIDs        []int
//	FolloweesIDsMap map[int]bool
//}
//func Constructor() Twitter {
//	dummyBolMap := make(map[int]bool)
//	dummySlice := make([]int, 0)
//	dummyData := Data{dummySlice,dummyBolMap}
//	dummyDataMap := make(map[int]Data)
//	dummyDataMap[0] = dummyData
//
//	h := Twitter{UserData: dummyDataMap}
//	return  h
//}
//
//func (this *Twitter) PostTweet(userId int, tweetId int) {
//	currentData := *this
//	userIdData := currentData.UserData[userId]
//	currentUserIdLastestTweets := append(userIdData.TweetIDs, tweetId)
//	userIdData.TweetIDs = currentUserIdLastestTweets
//	data := Data{TweetIDs: currentUserIdLastestTweets, FolloweesIDsMap: userIdData.FolloweesIDsMap}
//	mapData := make(map[int]Data)
//	mapData[userId] = data
//	currentData.UserData = mapData
//	this = &currentData
//}
//
//func (this *Twitter) GetNewsFeed(userId int) []int {
//	result := make([]int, 0)
//	currentData := *this
//	userIdData := currentData.UserData[userId]
//	var keys []int
//	for key, boo := range userIdData.FolloweesIDsMap {
//		if boo {
//			keys = append(keys, key)
//		}
//	}
//	for _, fid := range keys {
//		result = append(result, currentData.UserData[fid].TweetIDs...)
//	}
//
//	slices.Sort(result)
//	return result
//
//}
//
//func (this *Twitter) Follow(followerId int, followeeId int) {
//	currentData := *this
//	userIdData := currentData.UserData[followerId]
//	someMap := make(map[int]bool)
//	for k,v := range userIdData.FolloweesIDsMap{
//		someMap[k] = v
//	}
//	someMap[followeeId] = true
//	userIdData.FolloweesIDsMap = someMap
//
//	userIdData = Data{TweetIDs: userIdData.TweetIDs, FolloweesIDsMap: someMap}
//	currentData.UserData[followerId] = userIdData
//	this = &currentData
//}
//
//func (this *Twitter) Unfollow(followerId int, followeeId int) {
//	currentData := *this
//	userIdData := currentData.UserData[followerId]
//	userIdData.FolloweesIDsMap[followeeId] = false
//	currentData.UserData[followerId] = userIdData
//	this = &currentData
//}
//
//
///**
// * Your Twitter object will be instantiated and called as such:
// * obj := Constructor();
// * obj.PostTweet(userId,tweetId);
// * param_2 := obj.GetNewsFeed(userId);
// * obj.Follow(followerId,followeeId);
// * obj.Unfollow(followerId,followeeId);
// */


