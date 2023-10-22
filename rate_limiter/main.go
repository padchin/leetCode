package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu          sync.Mutex
	requests    []time.Time
	ratePerSec  int
	windowSize  time.Duration
	lastRefresh time.Time
}

func NewRateLimiter(ratePerSec int) *RateLimiter {
	limiter := &RateLimiter{
		ratePerSec: ratePerSec,
		windowSize: time.Second,
		requests:   make([]time.Time, ratePerSec),
	}
	return limiter
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.Sub(rl.lastRefresh) > rl.windowSize {
		// Если окно времени истекло, обнуляем запросы.
		for i := range rl.requests {
			rl.requests[i] = time.Time{}
		}
		rl.lastRefresh = now
	}

	index := rl.ratePerSec - 1
	if !rl.requests[index].IsZero() && now.Sub(rl.requests[index]) < rl.windowSize {
		// Если последний запрос произошел менее чем за секунду назад, отклоняем текущий запрос.
		return false
	}

	// Перемещаем все запросы на одну позицию вниз и записываем текущее время в первую позицию.
	copy(rl.requests[1:], rl.requests)
	rl.requests[0] = now

	return true
}

func main() {
	rateLimiter := NewRateLimiter(5) // Предположим, что ограничение - 5 запросов в секунду.

	for i := 0; i < 20; i++ {
		if rateLimiter.Allow() {
			fmt.Println("Разрешен запрос", i+1)
		} else {
			fmt.Println("Запрос", i+1, "отклонен")
		}
		time.Sleep(100 * time.Millisecond) // Задержка между запросами.
	}
}
