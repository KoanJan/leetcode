package main

type Twitter struct {
	Users      map[int]bool
	FollowInfo map[int]map[int]bool
	Tweets     []Tweet
}

type Tweet struct {
	Id     int
	UserId int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	twitter := Twitter{}
	twitter.Users = make(map[int]bool)
	twitter.FollowInfo = make(map[int]map[int]bool)
	twitter.Tweets = []Tweet{}
	return twitter
}

/** Compose a new tweet. */
func (this *Twitter) Posttweet(userId int, tweetId int) {
	if !this.Users[userId] {
		this.Users[userId] = true
	}
	this.Tweets = append(this.Tweets, Tweet{Id: tweetId, UserId: userId})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) Getnewsfeed(userId int) []int {
	res := []int{}
	for i := len(this.Tweets) - 1; i >= 0; i-- {
		if this.Tweets[i].UserId == userId {
			res = append(res, this.Tweets[i].Id)
		} else if _, ok := this.FollowInfo[userId]; ok {
			if this.FollowInfo[userId][this.Tweets[i].UserId] {
				res = append(res, this.Tweets[i].Id)
			}
		}
		if len(res) == 10 {
			break
		}
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.FollowInfo[followerId]; ok {
		this.FollowInfo[followerId][followeeId] = true
	} else {
		this.FollowInfo[followerId] = map[int]bool{followeeId: true}
	}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.FollowInfo[followerId]; ok {
		delete(this.FollowInfo[followerId], followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Posttweet(userId,tweetId);
 * param_2 := obj.Getnewsfeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
