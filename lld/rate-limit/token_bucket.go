package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	refillTime time.Duration
	maxTokens  int
	currTokens int
	mu         sync.Mutex
}

func (tb *TokenBucket) refillTicker() {
	ticker := time.NewTicker(tb.refillTime)
	defer ticker.Stop()
	for range ticker.C {
		tb.mu.Lock()
		if tb.currTokens < tb.maxTokens {
			//TODO : REMOVE
			fmt.Println("Token added")
			tb.currTokens += 1
		}
		tb.mu.Unlock()
	}
}

func (tb *TokenBucket) Usetokens(tokens int) (int, bool) {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	if tb.currTokens > 0 {
		tb.currTokens -= tokens
		return tb.currTokens, true
	} else {
		return tb.currTokens, false
	}
}

func NewTokenBucket(maxTokens int, refillTime time.Duration) *TokenBucket {
	tb := TokenBucket{
		maxTokens:  maxTokens,
		refillTime: refillTime,
		currTokens: maxTokens,
	}
	go tb.refillTicker()
	return &tb
}
