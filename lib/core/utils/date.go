package utils

import "time"

func GetExpiresAt(durationMs int64) int64 {
	var expiresAt int64 = -1

	if durationMs > 0 {
		expiresAt = time.Now().UnixMilli() + durationMs
	}

	return expiresAt
}
