package ping

import (
	"fmt"
	"net/http"
	"time"
)

func CheckHTTP(url string) string {
	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	if err != nil {
		return fmt.Sprintf("❌ %s - DOWN (%s)", url, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return fmt.Sprintf("✅ %s - UP (%.2fms)", url, float64(duration.Microseconds())/1000.0)
	}
	return fmt.Sprintf("⚠️ %s - UNHEALTHY (HTTP %d)", url, resp.StatusCode)
}
