package monitor

import (
	"sync"
)

type healthChecker struct {
	id          string
	description string
	check       HealthCheck
	sync.Mutex
	lastRes map[string]string
	lastErr error
}

func newHealthChecker(id, desc string, c HealthCheck) *healthChecker {
	return &healthChecker{
		id:          id,
		description: desc,
		check:       c,
		lastRes:     make(map[string]string),
	}
}

func (h *healthChecker) Id() string {
	return h.id
}

func (h *healthChecker) Description() string {
	return h.description
}

func (h *healthChecker) Run() (map[string]string, error) {
	h.Lock()
	defer h.Unlock()
	res, err := h.check()
	h.lastRes = res
	h.lastErr = err
	return res, err
}

func (h *healthChecker) Status() (map[string]string, error) {
	h.Lock()
	defer h.Unlock()
	return h.lastRes, h.lastErr
}
