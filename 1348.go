package leetcode

var freqToS = map[string]int{
	"minute": 60,
	"hour":   60 * 60,
	"day":    60 * 60 * 24,
}

type TweetCounts struct {
	Records map[string][]int
}

func Constructor() TweetCounts {
	return TweetCounts{
		Records: make(map[string][]int, 0),
	}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	this.Records[tweetName] = append(this.Records[tweetName], time)
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	step := freqToS[freq]
	allTime := this.Records[tweetName]
	slots := (endTime-startTime)/step + 1
	res := make([]int, slots)
	for _, t := range allTime {
		if t < startTime || t > endTime {
			continue
		}
		res[(t-startTime)/step] += 1
	}

	//logs.Info("%v", res)
	return res
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */
