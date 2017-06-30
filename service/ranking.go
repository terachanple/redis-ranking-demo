package service

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/terachanple/redis-ranking-demo/entity"
)

func GetCurrentDailyRanking() (entity.Rankings, error) {
	key := fmt.Sprintf("daily:%s", time.Now().Format("20060102"))

	result, err := redisClient.ZRevRangeWithScores(key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	rankings := make(entity.Rankings, 0, len(result))
	for _, v := range result {
		ranking := entity.Ranking{
			ID:    v.Member.(string),
			Score: int(v.Score),
		}
		rankings = append(rankings, ranking)
	}

	return rankings, nil
}

func IncrementDailyCountByID(id string) (*entity.Ranking, error) {
	key := fmt.Sprintf("daily:%s", time.Now().Format("20060102"))
	score, err := redisClient.ZIncrBy(key, 1, id).Result()
	if err != nil {
		return nil, err
	}

	if _, err := redisClient.Expire("key", 7*24*time.Hour).Result(); err != nil {
		return nil, err
	}

	return &entity.Ranking{ID: id, Score: int(score)}, nil
}

func GenrateWeeklyRanking() error {
	day := 7
	now := time.Now()
	dateFormat := "20060102"

	distKey := fmt.Sprintf("weekly:%s", now.Format(dateFormat))
	weights := make([]float64, 0, day)
	targetKeys := make([]string, 0, day)

	for i := 0; i < day; i++ {
		t := now.Add(time.Duration(i) * -24 * time.Hour)
		targetKeys = append(targetKeys, fmt.Sprintf("daily:%s", t.Format(dateFormat)))
		weights = append(weights, 1)
	}

	if _, err := redisClient.ZUnionStore(distKey, redis.ZStore{Weights: weights}, targetKeys...).Result(); err != nil {
		return err
	}

	if _, err := redisClient.Expire("key", 7*24*time.Hour).Result(); err != nil {
		return err
	}

	return nil
}
