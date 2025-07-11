package storage

import (
	"context"
	"urlshortener/config"

	"github.com/redis/go-redis/v9"
)

type URLStorage interface {
	SaveURL(shortPath, longURL string)
	GetURL(shortPath string) (string, bool)
}

type MapStorage struct {
	urlMap map[string]string
}

func NewMapStorage() *MapStorage {
	return &MapStorage{
		urlMap: make(map[string]string),
	}
}

func (m *MapStorage) SaveURL(shortPath, longURL string) {
	m.urlMap[shortPath] = longURL
}

func (m *MapStorage) GetURL(shortPath string) (string, bool) {
	val, ok := m.urlMap[shortPath]
	return val, ok
}

type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStorage(cfg *config.Config) *RedisStorage {
	if cfg.RedisAddr == "" {
		return nil
	}

	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	return &RedisStorage{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *RedisStorage) SaveURL(shortPath, longURL string) {
	if r.client != nil {
		r.client.Set(r.ctx, shortPath, longURL, 0)
	}
}

func (r *RedisStorage) GetURL(shortPath string) (string, bool) {
	if r.client != nil {
		val, err := r.client.Get(r.ctx, shortPath).Result()
		if err == nil {
			return val, true
		}
	}
	return "", false
}
