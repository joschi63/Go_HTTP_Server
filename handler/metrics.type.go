package handler

import (
	"sync/atomic"

	"github.com/joschi64/Go_HTTP_Server/internal/database"
)

type ApiConfig struct {
	SECRET         string
	PolkaKey       string
	DB             *database.Queries
	PLATFORM       string
	fileserverHits atomic.Int32
}
