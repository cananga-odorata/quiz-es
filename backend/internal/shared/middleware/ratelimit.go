package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// IPRateLimiter holds rate limiters for each IP address
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// NewIPRateLimiter creates a new IPRateLimiter
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	// Clean up old IP entries periodically
	go i.cleanupLoop()

	return i
}

// GetLimiter returns the rate limiter for the provided IP address
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.ips[ip]
	if !exists {
		limiter = rate.NewLimiter(i.r, i.b)
		i.ips[ip] = limiter
	}

	return limiter
}

// cleanupLoop removes limits for IPs that haven't been seen in a while to prevent memory leaks
func (i *IPRateLimiter) cleanupLoop() {
	for {
		time.Sleep(5 * time.Minute)
		i.mu.Lock()
		// In a real implementation, we would track last access time and remove stale entries.
		// For simplicity/demo, we just clear the map occasionally or implement LRU.
		// Resetting map to prevent infinite growth in this demo
		if len(i.ips) > 10000 {
			i.ips = make(map[string]*rate.Limiter)
		}
		i.mu.Unlock()
	}
}

// RateLimitMiddleware creates a middleware for rate limiting
func RateLimitMiddleware(limit float64, burst int) func(next http.Handler) http.Handler {
	limiter := NewIPRateLimiter(rate.Limit(limit), burst)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract IP without port (e.g. "127.0.0.1:54321" -> "127.0.0.1")
			ip := r.RemoteAddr
			if host, _, err := net.SplitHostPort(ip); err == nil {
				ip = host
			}

			if !limiter.GetLimiter(ip).Allow() {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte(`{"error": "Too Many Requests", "message": "Rate limit exceeded"}`))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
