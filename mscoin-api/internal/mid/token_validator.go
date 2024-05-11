package mid

import (
	"coin-common/common"
	"coin-common/tools"
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// TokenValidator 是用于校验 token 是否存在的中间件
func TokenValidator(secret string) func(next http.HandlerFunc) http.Handler {
	return func(next http.HandlerFunc) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 获取 Authorization 头部
			token := r.Header.Get("Authorization")
			result := common.NewResult()
			result.Fail(4000, "no ucenter")

			if token == "" {
				httpx.WriteJson(w, 200, result)
				return
			}

			userId, err := tools.ParseToken(token, secret)
			if err != nil {
				httpx.WriteJson(w, 200, result)
				return
			}
			context.WithValue(r.Context(), "userId", userId)
			next.ServeHTTP(w, r)
		})
	}
}
