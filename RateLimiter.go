package main

//
//import (
//	"fmt"
//	"time"
//
//	"golang.org/x/net/context"
//)
//
//type SlidingWindowLogRatelimiter struct {
//	con *redis.Client
//
//	REQUESTS        string
//	WINDOW_TIME     string
//	METADATA_SUFFIX string
//	TIMESTAMPS      string
//	INF             int64
//}
//
//func NewSlidingWindowLogRatelimiter() *SlidingWindowLogRatelimiter {
//	return &SlidingWindowLogRatelimiter{
//		con: redis.NewClient(&redis.Options{
//			Addr: "127.0.0.1:6379",
//			DB:   0,
//		}),
//		REQUESTS:        "requests",
//		WINDOW_TIME:     "window_time",
//		METADATA_SUFFIX: "_metadata",
//		TIMESTAMPS:      "_timestamps",
//		INF:             9999999999,
//	}
//}
//
//// GetCurrentTimestampInSec returns the current timestamp in seconds
//func (r *SlidingWindowLogRatelimiter) GetCurrentTimestampInSec() int64 {
//	return time.Now().Unix()
//}
//
//// AddUser adds a new user's rate of requests to be allowed
//func (r *SlidingWindowLogRatelimiter) AddUser(userId string, requests int64, windowTimeInSec int64) {
//	r.con.HSet(context.Background(), userId+r.METADATA_SUFFIX, r.REQUESTS, requests, r.WINDOW_TIME, windowTimeInSec)
//}
//
//// GetRateForUser retrieves the user metadata storing the number of requests per window time
//func (r *SlidingWindowLogRatelimiter) GetRateForUser(userId string) (int64, int64, error) {
//	val, err := r.con.HGetAll(context.Background(), userId+r.METADATA_SUFFIX).Result()
//	if err != nil || len(val) == 0 {
//		return 0, 0, fmt.Errorf("Un-registered user: %s", userId)
//	}
//	requests := val[r.REQUESTS]
//	windowTime := val[r.WINDOW_TIME]
//	return requests, windowTime, nil
//}
//
//// RemoveUser removes a user's metadata and timestamps
//func (r *SlidingWindowLogRatelimiter) RemoveUser(userId string) {
//	r.con.Del(context.Background(), userId+r.METADATA_SUFFIX, userId+r.TIMESTAMPS)
//}
//
//// AddTimeStampAtomicallyAndReturnSize atomically adds an element to the timestamps and returns the total number of requests
//func (r *SlidingWindowLogRatelimiter) AddTimeStampAtomicallyAndReturnSize(userId string, timestamp int64) (int64, error) {
//	pipe := r.con.TxPipeline()
//	pipe.ZAdd(context.Background(), userId+r.TIMESTAMPS, &redis.Z{Score: float64(timestamp), Member: timestamp})
//	pipe.ZCount(context.Background(), userId+r.TIMESTAMPS, "0", fmt.Sprintf("%d", r.INF))
//	_, err := pipe.Exec(context.Background())
//	if err != nil {
//		return 0, err
//	}
//	count, err := pipe.ZCount(context.Background(), userId+r.TIMESTAMPS, "0", fmt.Sprintf("%d", r.INF)).Result()
//	return count, err
//}
//
//// ShouldAllowServiceCall decides to allow a service call or not
//func (r *SlidingWindowLogRatelimiter) ShouldAllowServiceCall(userId string) (bool, error) {
//	maxRequests, unitTime, err := r.GetRateForUser(userId)
//	if err != nil {
//		return false, err
//	}
//	currentTimestamp := r.GetCurrentTimestampInSec()
//	oldestPossibleEntry := currentTimestamp - unitTime
//	r.con.ZRemRangeByScore(context.Background(), userId+r.TIMESTAMPS, "0", fmt.Sprintf("%d", oldestPossibleEntry))
//	currentRequestCount, err := r.AddTimeStampAtomicallyAndReturnSize(userId, currentTimestamp)
//	if err != nil {
//		return false, err
//	}
//	fmt.Println(currentRequestCount, maxRequests)
//	if currentRequestCount > maxRequests {
//		return false, nil
//	}
//	return true, nil
//}
