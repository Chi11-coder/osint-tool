package host

import "time"

func TimeGenerate() time.Time {
	now := time.Now()
	y, m, d := now.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Location())
}

func SameCacheDate(t1 time.Time) bool {
	today := TimeGenerate()
	t1Midnight := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
	return t1Midnight.Equal(today)
}
