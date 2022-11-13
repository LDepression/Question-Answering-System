package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"math"
	"time"
	"wenba/internal/model/common"
	"wenba/internal/model/request"
	"wenba/internal/pkg/utils"
)

// 比方说需要200个投票才能给问题续一天
const (
	ScorePer = 24 * 3600 / 200
	OneWeek  = 24 * 3600 * 7
)

var ErrExpired = errors.New("投票已经超时了")

//投票功能的实现
//我们假定只让投票7天,7天之后就不能进行投票了
//1.先去判断投票的限制,是否超时之类的
//2.更新回答的分数
//3.记录该用户为该回答的投票记录

//本项目使用简化版的投票分数
/*
	投票的几种情况:
direction=1时,有两种情况:
	1.之前没有投反对票                        ->更新分数和投票记录 差值的绝对值:1 +432
	2.之前投了反对票,现在改投赞成票             ->更新分数和投票记录 差值的绝对值:2 +432*2
direction=0时,有两种情况
	1.之前投过赞成票,现在要取消赞成票								差值的绝对值:1 +432
	2.之前投过反对票,现在要取消投票								差值的绝对值:1 -432
direction=-1时,有两种情况:
	1.之前没有投过票,现在投反对票								差值的绝对值:1 -432
	2.之前投赞成票,现在改投反对票								差值的绝对值:2 -432*2
*/

//CreateAnswer 完成对问题redis数据库的修改
func CreateAnswer(answerID int64) error {
	ctx := context.Background()
	pipe := rdb.TxPipeline()
	pipe.ZAdd(ctx, getRedisKey(KeyAnswerTimeZset), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: answerID,
	})
	pipe.ZAdd(ctx, getRedisKey(KeyAnswerScoreZset), &redis.Z{
		Score:  0,
		Member: answerID,
	})
	_, err := pipe.Exec(ctx)
	return err
}

func VoteForAnswer(userID int64, vote request.ReqVote) error {
	ctx := context.Background()
	//先来判断时间是否过期
	postTime := rdb.ZScore(ctx, getRedisKey(KeyAnswerTimeZset), utils.IDToSting(vote.AnswerID)).Val()
	//此时表示超时了
	if int(time.Now().Unix())-int(postTime) > OneWeek {
		return ErrExpired
	}
	var dir float64
	ov := rdb.ZScore(ctx, getRedisKey(KeyAnswerVotedZsetPF+utils.IDToSting(vote.AnswerID)), utils.IDToSting(userID)).Val()
	if float64(vote.Direction) > dir {
		dir = 1
	} else {
		dir = -1
	}
	//现在来进行更新分数
	//先去根据用户去取得用户的投票记录
	pipe := rdb.TxPipeline()
	diff := math.Abs(ov-float64(vote.Direction)) * dir
	pipe.ZIncrBy(ctx, getRedisKey(KeyAnswerScoreZset), diff*ScorePer, utils.IDToSting(vote.AnswerID))
	//记录用户为该帖子投票的数据
	if vote.Direction == 0 {
		pipe.ZRem(ctx, getRedisKey(KeyAnswerVotedZsetPF+utils.IDToSting(vote.AnswerID)), userID)
	} else {
		pipe.ZAdd(ctx, getRedisKey(KeyAnswerVotedZsetPF)+utils.IDToSting(vote.AnswerID), &redis.Z{
			Score:  float64(vote.Direction), //记录所投的票是赞成票还是反对票
			Member: userID,
		})
	}
	_, err := pipe.Exec(ctx)
	return err
}

func GetAnswerIDsInOrder(p request.ReqAnswerList) (data []string, err error) {
	ctx := context.Background()
	key := getRedisKey(KeyAnswerTimeZset)
	if p.Order == common.OrderScore {
		key = getRedisKey(KeyAnswerScoreZset)
	}
	//确定查询的起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	data, err = rdb.ZRevRange(ctx, key, start, end).Result()
	return
}

//GetVoteNum 获取投票的数量
func GetVoteNum(answerIDs []string) ([]int64, error) {
	ctx := context.Background()
	pipe := rdb.TxPipeline()
	for _, answerID := range answerIDs {
		key := getRedisKey(KeyAnswerVotedZsetPF) + answerID
		pipe.ZCount(ctx, key, "1", "1")
	}
	cmders, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}
	data := make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return data, nil
}
