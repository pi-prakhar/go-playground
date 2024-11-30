package main

import (
	"fmt"
	"sync"
	"time"
)

type SlidingWindow struct {
	windowLength time.Duration
	maxTokens    int
	tokens       []time.Time
	mu           sync.Mutex
}

func NewSlidingWindow(windowLength time.Duration, maxTokens int) *SlidingWindow {
	var tokens []time.Time
	sw := SlidingWindow{
		windowLength: windowLength,
		maxTokens:    maxTokens,
		tokens:       tokens,
	}
	return &sw
}

func (sw *SlidingWindow) UseToken() (int, bool) {
	now := time.Now()
	lastAcceptedTime := now.Add(-sw.windowLength)

	var newIndex int = 0

	// get accepted tokens only
	for index, token := range sw.tokens {
		if token.After(lastAcceptedTime) {
			newIndex = index
			break
		}
	}

	// update the tokens array
	sw.tokens = sw.tokens[newIndex:]
	//printTokens(sw.tokens)

	// check to see if we can use a new token
	if len(sw.tokens) < sw.maxTokens {
		sw.tokens = append(sw.tokens, now)
		return sw.maxTokens - len(sw.tokens), true
	} else {
		return sw.maxTokens - len(sw.tokens), false
	}
}

func printTokens(tokens []time.Time) {
	for _, token := range tokens {
		fmt.Print(token.Format("05.000"), " ")
	}
}
