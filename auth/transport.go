package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	errMethodNotAllowed = errors.New("auth: method not allowed")
	errForbidden        = errors.New("auth: forbidden")
	errBadRequest       = errors.New("auth: bad request body")
	errUnauthorized     = errors.New("auth: Unauthorized")
)

type errorResponse struct {
	Error string `json:"error"`
}

// MakeMiddleware creates new auth middleware
func MakeMiddleware(s Service) func(http.Handler) http.Handler {
	fmt.Println("start exec MakeMiddleware ....")
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// get token from header
			tokenID := r.Header.Get("x-Access-Token")

			fmt.Println("auth.transport token : ", tokenID)
			if len(tokenID) == 0 {
				h.ServeHTTP(w, r)
				return
			}

			fmt.Println("s.GetToken")

			tk, err := s.GetToken(tokenID)
			if err != nil {
				// h.ServeHTTP(w, r)
				if !strings.Contains(r.URL.String(), "signin") {
					errorEncoder(w, r, err)
					return
				}
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, keyToken{}, tk)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

func mustLogin(s Service) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientID := GetCompanyID(r.Context())
			if clientID < 0 {
				errorEncoder(w, r, errForbidden)
				fmt.Println("error mustLogin auth.transport.go")
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func jsonDecoder(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		log.Printf("[Auth] API request: %+v\n", v)
		return errBadRequest
	}
	return nil
}

func jsonEncoder(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Token,x-access-token")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(status)
	log.Printf("[Auth] API response %d: %+v\n", status, v)
	return json.NewEncoder(w).Encode(v)
}

func requestDecoder(r *http.Request, v interface{}) error {
	if r.Method != http.MethodPost {
		return errMethodNotAllowed
	}
	return jsonDecoder(r, v)
}

func responseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) {
	jsonEncoder(w, http.StatusOK, v)
}

func errorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	encoder := jsonEncoder
	status := http.StatusInternalServerError
	switch err {
	case errMethodNotAllowed:
		status = http.StatusMethodNotAllowed
	case errForbidden:
		status = http.StatusForbidden
	case errBadRequest:
		status = http.StatusBadRequest
	case errUnauthorized:
		status = http.StatusUnauthorized
	case ErrTokenExpired:
		status = http.StatusUnauthorized
	case ErrTokenNotFound:
		status = http.StatusUnauthorized
	}
	if r.Method == http.MethodOptions {
		encoder(w, http.StatusNoContent, nil)
	} else {
		encoder(w, status, &errorResponse{err.Error()})
	}
}
