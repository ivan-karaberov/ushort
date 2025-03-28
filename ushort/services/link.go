package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
	"ushort/config"
	"ushort/storage"
)

type Link struct {
	Id       string `json:"-"`
	Url      string `json:"url"`
	Password string `json:"password,omitempty"`
}

func SetInRedis(ctx context.Context, cfg config.Config, link *Link) bool {
	rdb, err := storage.RedisClient(ctx, cfg)
	if err != nil {
		log.Printf("failed connect to Redis")
		return false
	}

	linkJSON, err := json.Marshal(link)
	if err != nil {
		log.Printf("failed encode to JSON")
		return false
	}

	err = rdb.Set(ctx, link.Id, linkJSON, 0).Err()
	return err == nil
}

func GetFromRedis(ctx context.Context, cfg config.Config, id string) *Link {
	rdb, err := storage.RedisClient(ctx, cfg)
	if err != nil {
		log.Printf("failed connect to Redis")
		return nil
	}

	linkJSON, err := rdb.Get(ctx, id).Result()
	if err != nil {
		return nil
	}

	var link Link

	err = json.Unmarshal([]byte(linkJSON), &link)
	if err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		return nil
	}

	return &link
}

func GenerateRandomID(length int) string {
	const chr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.New(rand.NewSource(time.Now().UnixNano()))
	id := make([]byte, length)
	for i := range id {
		id[i] = chr[rand.Intn(len(chr))]
	}
	return string(id)
}

func SaveLink(ctx context.Context, cfg config.Config, url string, password string) (string, error) {
	id := GenerateRandomID(8)

	counter := 0 // делаем 3 попытки записи
	for counter < 3 {
		if GetFromRedis(ctx, cfg, id) != nil {
			log.Printf("ID %s already exists in Redis. Generating a new ID.", id)
			id = GenerateRandomID(8)
			counter += 1
		} else {
			link := Link{Id: id, Url: url, Password: password}
			if SetInRedis(ctx, cfg, &link) {
				return id, nil
			}
			return "", fmt.Errorf("failed to set link in Redis")
		}
	}

	return "", fmt.Errorf("failed to save link after 3 attempts")
}
