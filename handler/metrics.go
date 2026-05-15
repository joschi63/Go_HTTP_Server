package handler

import (
	"fmt"
	"net/http"
)

func (a *ApiConfig) HandleServerHits(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf(`<html>
  		<body>
    		<h1>Welcome, Chirpy Admin</h1>
   			<p>Chirpy has been visited %d times!</p>
  		</body>
	</html>`, a.fileserverHits.Load())))
}

func (a *ApiConfig) HandleHitsReset(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("Serverhits have been reseted"))
	a.fileserverHits.Store(0)
}

func (a *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
