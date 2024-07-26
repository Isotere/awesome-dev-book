package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Circuit func(context.Context) (string, error)

// Breaker
// - failureThreshold количество отказов для размыкания
func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	consecutiveFailures := 0
	lastAttemtp := time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		m.RLock()

		d := consecutiveFailures - int(failureThreshold)

		if d >= 0 {
			shouldRetryAt := lastAttemtp.Add(time.Second * 2 << d)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()

				return "", errors.New("service unreachable")
			}
		}

		m.RUnlock()

		response, err := circuit(ctx)

		m.Lock()
		defer m.Unlock()

		lastAttemtp = time.Now()

		if err != nil {
			consecutiveFailures++

			return response, err
		}

		consecutiveFailures = 0

		return response, nil
	}
}
