package tools

import (
	"net"
	"net/http"
	"strings"
)

func GetRemoteClientIp(r *http.Request) string {
	// 首先检查 X-Forwarded-For 头部
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0]) // 返回第一个 IP
	}

	// 如果没有 X-Forwarded-For，检查 X-Real-IP
	xri := r.Header.Get("X-Real-IP")
	if xri != "" {
		return strings.TrimSpace(xri)
	}

	// 如果没有代理头部，返回 RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr) // 分割地址和端口
	if err != nil {
		return r.RemoteAddr // 如果 SplitHostPort 失败，返回原始地址
	}
	return host
}
