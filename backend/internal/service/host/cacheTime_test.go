package host_test

import (
	"testing"
	"time"

	"example.com/security/internal/service/host"
)

func TestIsSameDay_Returned_True_1(t *testing.T) {
	tests := struct {
		name string
		data time.Time
	}{
		"case1: 同一日",
		host.TimeGenerate(),
	}

	t.Run(tests.name, func(t *testing.T) {
		if ok := host.SameCacheDate(tests.data); ok {
			return
		} else {
			t.Errorf("failed test host.IsSameDay returned: %t", ok)
		}

	})
}

func TestIsSameDay_Returned_False_2(t *testing.T) {
	tests := struct {
		name string
		data time.Time
	}{
		"case2: 1日経過",
		host.TimeGenerate().AddDate(0, 0, 1),
	}

	t.Run(tests.name, func(t *testing.T) {
		if ok := host.SameCacheDate(tests.data); ok {
			t.Errorf("failed test host.IsSameDay returned: %t", ok)
		}
	})
}

func TestIsSameDay_Returned_True_3(t *testing.T) {
	tests := struct {
		name string
		data time.Time
	}{
		"case3: 同じ日で異なる時間にキャッシュに保存されたデータを検索",
		time.Now(),
	}

	t.Run(tests.name, func(t *testing.T) {
		sameDay := time.Date(tests.data.Year(), tests.data.Month(), tests.data.Day(), 14, 30, 0, 0, tests.data.Location())

		if !host.SameCacheDate(sameDay) {
			t.Error("same day with different time should return true")
		}
	})
}
