package middleware

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
	"github.com/ngobrut/eniqlo-store-api/pkg/constant"
	"github.com/sirupsen/logrus"
)

func Recover(log *logrus.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					buf := make([]byte, 2084)
					n := runtime.Stack(buf, false)
					buf = buf[:n]

					log.Errorf("[err] %v\n", err)
					log.Errorf("[stack-trace] \n%s\n", buf)

					w.Header().Set("Content-Type", "application/json; charset=utf-8")
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(response.JsonResponse{
						Message: "Error",
						Error: &response.ErrorResponse{
							Code:    http.StatusInternalServerError,
							Message: constant.HTTPStatusText(http.StatusInternalServerError),
						},
					})
				}
			}()

			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
