package ratelimit

import (
	"net/http"

	"github.com/go-redis/redis_rate/v8"
	"github.com/labstack/echo/v4"
)

type RateLimiter struct {
	Limiter *redis_rate.Limiter
	Rate    *redis_rate.Limit
}

func (limiter *RateLimiter) Limit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := limiter.Limiter.Allow(c.RealIP(), limiter.Rate)
		if err != nil {
			return err
		}
		if res.Allowed {
			return next(c)
		}
		return echo.NewHTTPError(http.StatusTooManyRequests)
	}
}
