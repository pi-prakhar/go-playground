package main

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindow struct {
	windowLength time.Duration
	maxTokens    int
	tokens       *[]time.Time
	mu           sync.Mutex
}

func NewSlidingWindow(windowLength time.Duration, maxTokens int) *SlidingWindow {
	var tokens []time.Time
	sw := SlidingWindow{
		windowLength: windowLength,
		maxTokens:    maxTokens,
		tokens:       &tokens,
	}
	return &sw
}

func (sw *SlidingWindow) UseToken() (int, bool) {
	now := time.Now()
	lastAcceptedTime := now.Add(-sw.windowLength)

	var newIndex int
	fmt.Println("last : ", *sw.tokens)

	// get accepted tokens only
	for index, token := range *sw.tokens {
		if !token.After(lastAcceptedTime) {
			newIndex = index
		}
	}

	// update the tokens array
	(*sw.tokens) = (*sw.tokens)[:newIndex]

	// check to see if we can use a new token
	if len(*sw.tokens) <= sw.maxTokens {
		(*sw.tokens) = append(*sw.tokens, now)
		return len(*sw.tokens), true
	} else {
		return len(*sw.tokens), false
	}
}
