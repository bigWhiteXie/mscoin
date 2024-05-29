package config

import (
	"coin-common/queue"
	"coin-common/tools"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	Okx     *Okx
	Symbols []string
	Mongo   Mongo
	Kafka   queue.KafkaConfig
}

//type KafkaConfig struct {
//	Addr     string   `json:"addr,optional"`
//	WriteCap int      `json:"writeCap,optional"`
//	ReadCap  int      `json:"readCap,optional"`
//	Groups   []string `json:"Groups,optional"`
//}

type Mongo struct {
	Url      string
	Username string
	Password string
	DataBase string
}
type Okx struct {
	Key        string
	Secret     string
	Passphrase string
	Host       string
	Proxy      string
}

func (o *Okx) SetHeaderAndProxy(req *http.Request, timestamp string, message string) *http.Client {
	proxyURL, err := url.Parse(o.Proxy) // 代理地址
	if err != nil {
		log.Fatalf("Failed to parse proxy URL: %v", err)
	}

	// 配置 http.Transport 并设置代理
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL), // 使用代理
	}
	// 准备 HMAC SHA256 签名
	signature := tools.GenerateHmacSha256(message, o.Secret) // 生成 HMAC SHA256 签名，替换为你的秘密密钥
	// 设置请求头
	req.Header.Set("OK-ACCESS-KEY", o.Key)
	req.Header.Set("OK-ACCESS-SIGN", signature) // 使用 Base64 编码的 HMAC SHA256 签名
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", o.Passphrase)

	// 发送 HTTP 请求
	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}
	return client
}
