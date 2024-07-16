package cache

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"snake-game/internal/game"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(address string) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &RedisClient{Client: client}, nil
}

func (rc *RedisClient) GetPlayer(playerID string) (*game.Player, error) {
	data, err := rc.Client.Get(playerID).Bytes()
	if err != nil {
		return nil, err
	}

	var player game.Player
	err = json.Unmarshal(data, &player)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (rc *RedisClient) SetPlayer(playerID string, player *game.Player) error {
	data, err := json.Marshal(player)
	if err != nil {
		return err
	}

	return rc.Client.Set(playerID, data, 0).Err()
}

func (rc *RedisClient) Close() error {
	return rc.Client.Close()
}
