package main

import (
	"crypto/sha256"
	"math/rand"
)

var letterRunes = []int32("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]int32, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func hashUrl(url string) string {
	h := sha256.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)
	return string(bs)
}
