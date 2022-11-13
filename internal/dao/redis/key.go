package redis

const (
	KeyPrefix            = "wenba:"
	KeyAnswerTimeZset    = "answer:time"  //按照时间排序来进行的投票
	KeyAnswerScoreZset   = "answer:score" //按照分数来进行的投票
	KeyAnswerVotedZsetPF = "answer:vote"  //这是记录用户的投票类型,后面要另外加上Answer_id作为key,存的值就是用户的投票类型
)

func getRedisKey(s string) string {
	return KeyPrefix + s
}
