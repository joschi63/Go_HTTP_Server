package handler

import (
	"sync/atomic"

	"github.com/joschi64/Go_HTTP_Server/internal/database"
)

type ApiConfig struct {
	DB             *database.Queries
	fileserverHits atomic.Int32
}
