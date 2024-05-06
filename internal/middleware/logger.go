package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (r *customResponseWriter) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func RequestLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := &customResponseWriter{ResponseWriter: w}

		t1 := time.Now()

		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("[error-io.ReadAll()] \n%v\n", err)
		}

		defer func() {
			err = r.Body.Close()
			if err != nil {
				log.Printf("[error-body.Close()] \n%v\n", err)
			}

			fields := make(map[string]interface{})

			if len(b) > 0 {
				body := make(map[string]interface{})
				err = json.Unmarshal(b, &body)
				if err != nil {
					log.Printf("[error-json.Unmarshal()] \n%v\n", err)
				}

				fields["@request"] = body
			}

			token, _ := GetTokenFromHeader(r)
			claims := ParseWithoutVerified(token)
			if token != "" && claims != nil {
				fields["@auth"] = map[string]interface{}{
					"user_id": claims.UserID,
				}
			}

			logfield, err := json.Marshal(fields)
			if err != nil {
				log.Printf("[error-json.Marshal()] \n%v\n", err)
			}

			log.Printf("%s from %s - %d in %s \n", fmt.Sprintf("%s %s %s", r.Method, r.RequestURI, r.Proto), r.RemoteAddr, ww.statusCode, time.Since(t1).Abs().String())
			log.Println(string(logfield))
		}()

		r.Body = io.NopCloser(bytes.NewBuffer(b))

		h.ServeHTTP(ww, r)
	})
}
