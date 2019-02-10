package middle

import (
	"context"
	"encoding/json"
	"github.com/hplauction/auction"
	"net/http"
	"strings"
)

func Authenticator(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS Headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, x-requested-with, origin, X-API-VERSION")
		w.Header().Set("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
		w.Header().Set("Content-Type", "application/json")

		// OK for all pre-flight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		urlParts := strings.Split(r.URL.Path, "/")
		apiName := urlParts[1]

		if r.Method == "POST" || apiName == "login" {
			res, err := auction.LoginHandler{}.Create(r)
			if err != nil {
				json.NewEncoder(w).Encode(err)
				return
			}
			json.NewEncoder(w).Encode(res)
			return
		}
		if !isValidApi(apiName) {
			w.WriteHeader(404)
			response := Response{404, "ERROR", "Not Found", nil}
			json.NewEncoder(w).Encode(response)
			return
		}

		ok, ctx := authorised(r)
		if !ok {
			w.WriteHeader(401)
			response := Response{401, "ERROR", "Not Authorised", nil}
			json.NewEncoder(w).Encode(response)
			return
		}

		if ctx != nil {
			*r = *r.WithContext(ctx)
			inner.ServeHTTP(w, r)
		} else {
			inner.ServeHTTP(w, r)
		}
	})
}

func isValidApi(api string) bool {
	if _, ok := allowedApi[api]; ok {
		return true
	}
	return false
}
func authorised(r *http.Request) (t bool, ctx context.Context) {
	params := make(map[string]interface{})
	ctx = context.WithValue(r.Context(), "countryIsoCode", "IN")
	token := r.Header.Get("Authorization")
	s := strings.Split(token, " ")
	if len(s) == 2 {
		params["accessToken"] = s[1]
	}
	resp := auction.Authorise(params)
	ctx = context.WithValue(ctx, "isAdmin", resp.IsAdmin)
	ctx = context.WithValue(ctx, "id", resp.UserId)
	if resp.Code == 200 {
		return true, ctx
	}
	return false, ctx
}
