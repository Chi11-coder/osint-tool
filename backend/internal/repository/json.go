//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=repository
package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type JsonRepository[T any] interface {
	Load(ctx context.Context, v string) (*T, error)
	Save(ctx context.Context, v string, d *T) error
}

type JsonFile[T any] struct {
	FileName string
	Reports  []T
	mu       sync.RWMutex
}

func NewJsonRepository[T any](filename string) JsonRepository[T] {
	return &JsonFile[T]{FileName: filename}
}

func (j *JsonFile[T]) Load(ctx context.Context, v string) (*T, error) {
	j.mu.RLock()
	defer j.mu.RUnlock()

	f, err := j.openJsonFile()
	if err != nil {
		return nil, fmt.Errorf("failed to open: %w", err)
	}
	if f == nil {
		return nil, nil
	}
	defer f.Close()

	var d = map[string]T{}

	if err := json.NewDecoder(f).Decode(&d); err != nil {
		return nil, fmt.Errorf("failed to decode: %w", err)
	}

	if val, ok := d[v]; ok {
		return &val, nil
	}

	return nil, nil
}

// ファイル書き込み
func (j *JsonFile[T]) Save(ctx context.Context, v string, d *T) (retErr error) {
	j.mu.Lock()
	defer j.mu.Unlock()

	// T型の空マップを作成
	currentData := make(map[string]T)

	f, err := j.openJsonFile()
	if err != nil {
		return fmt.Errorf("failed to open: %w", err)
	}

	if f != nil {
		if err := json.NewDecoder(f).Decode(&currentData); err != nil {
			if closeErr := f.Close(); closeErr != nil {
				return fmt.Errorf("decode error: %w (close failed: %v)", err, closeErr)
			}
			return fmt.Errorf("decode cache: %w", err)
		}
		if closeErr := f.Close(); closeErr != nil {
			return fmt.Errorf("close cache: %w", closeErr)
		}
	}

	// API実行結果がnilでないならcurrentDataにデータを追加
	if d != nil {
		currentData[v] = *d
	}

	outFile, err := os.Create(j.FileName)
	if err != nil {
		return fmt.Errorf("file path: %w", err)
	}

	defer func() {
		if closeErr := outFile.Close(); closeErr != nil && retErr == nil {
			retErr = fmt.Errorf("failed to close cache: %w", closeErr)
		}
	}()

	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(currentData); err != nil {
		return fmt.Errorf("file encode: %w", err)
	}

	return nil
}

// Jsonファイル読み取り
func (j *JsonFile[T]) openJsonFile() (*os.File, error) {
	// ファイルが存在しない場合はnilを返却
	if _, err := os.Stat(j.FileName); os.IsNotExist(err) {
		return nil, nil
	}
	// ファイル読み取り
	f, err := os.Open(j.FileName)
	if err != nil {
		return nil, fmt.Errorf("file opened: %w", err)
	}

	// ファイルサイズが空ならnilを返却
	info, err := f.Stat()
	if err != nil {
		if closeErr := f.Close(); closeErr != nil {
			return nil, fmt.Errorf("stat cache file: %w (close failed: %v)", err, closeErr)
		}
		return nil, err
	}
	// ファイルが存在するがサイズが0ならキャッシュが存在しないためファイルを閉じnilを返却
	if info.Size() == 0 {
		if closeErr := f.Close(); closeErr != nil {
			return nil, fmt.Errorf("close empty cache file: %w", closeErr)
		}
		return nil, nil
	}

	return f, nil
}
