package services

import (
	"math/rand"
	"time"
)

type Link struct {
	Id       string `json:"-"`
	Url      string `json:"url"`
	Password string `json:"password,omitempty"`
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
