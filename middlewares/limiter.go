package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

var (
	limiters = make(map[string]*rate.Limiter)
	lastSeen = make(map[string]time.Time) // Track last request time per IP
	mutex    = &sync.Mutex{}
)

const cleanupInterval = 5 * time.Minute // Remove inactive limiters every 5 minutes

// Get or create a rate limiter for a user (IP-based)
func getLimiter(ip string) *rate.Limiter {
	mutex.Lock()
	defer mutex.Unlock()

	lastSeen[ip] = time.Now()

	if limiter, exists := limiters[ip]; exists {
		return limiter
	}

	// Create a new rate limiter: 10 requests per minute (1 token every 6s)
	limiter := rate.NewLimiter(rate.Every(time.Minute/10), 10)
	limiters[ip] = limiter
	return limiter
}

func CleanupLimiters() {
	for {
		time.Sleep(cleanupInterval)

		mutex.Lock()
		now := time.Now()

		for ip, lastTime := range lastSeen {
			if now.Sub(lastTime) > cleanupInterval {
				delete(limiters, ip)
				delete(lastSeen, ip)
			}
		}

		mutex.Unlock()
	}
}

// Middleware to check the rate limit.
func Limiter(ctx *gin.Context) {
	ip := ctx.ClientIP()
	limiter := getLimiter(ip)

	// Check if the request is allowed
	if !limiter.Allow() {
		ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
		return
	}

	ctx.Next()
}
