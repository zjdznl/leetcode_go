package leetcode

import (
	"code.byted.org/gopkg/logs"
	"testing"
)

func TestConstructor(t *testing.T) {
	defer logs.Flush()
	tweetCounts := Constructor()
	tweetCounts.RecordTweet("tweet3", 0);
	tweetCounts.RecordTweet("tweet3", 60);
	tweetCounts.RecordTweet("tweet3", 10);                             // All tweets correspond to "tweet3" with recorded times at 0, 10 and 60.
	tweetCounts.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59); // return [2]. The frequency is per minute (60 seconds), so there is one interval of time: 1) [0, 60> - > 2 tweets.
	tweetCounts.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60); // return [2, 1]. The frequency is per minute (60 seconds), so there are two intervals of time: 1) [0, 60> - > 2 tweets, and 2) [60,61> - > 1 tweet.
	tweetCounts.RecordTweet("tweet3", 120);                            // All tweets correspond to "tweet3" with recorded times at 0, 10, 60 and 120.
	tweetCounts.GetTweetCountsPerFrequency("hour", "tweet3", 0, 210);  // return [4]. The frequency is per hour (3600 seconds), so there is one interval of time: 1) [0, 211> - > 4 tweets.
}
