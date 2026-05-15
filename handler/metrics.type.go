package handler

import "sync/atomic"

type ApiConfig struct {
	fileserverHits atomic.Int32
}
