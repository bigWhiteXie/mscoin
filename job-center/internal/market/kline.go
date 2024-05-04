package market

import (
	"fmt"
	"io/ioutil"
	"job-center/internal/config"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// kline的定时任务领域
type Kline struct {
	wg      *sync.WaitGroup
	c       *config.Config
	reqPath string
}

func NewKline(c *config.Config) *Kline {
	return &Kline{
		wg:      &sync.WaitGroup{},
		c:       c,
		reqPath: "/api/v5/market/candles",
	}
}

func (k *Kline) Do(period string) {
	log.Println("============启动k线数据拉取==============")
	size := len(k.c.Symbols)
	k.wg.Add(size)
	for _, symbol := range k.c.Symbols {
		log.Printf("===============拉取%s数据===============\n", symbol)
		go k.syncToMongo(symbol, strings.Replace(symbol, "-", "/", 1), period)
	}
	k.wg.Wait()
	log.Println("===============k线数据拉取结束===============")
}

func (k *Kline) syncToMongo(instId string, symbol string, period string) {
	defer k.wg.Done()
	// 配置请求参数
	requestPath := fmt.Sprintf(k.reqPath+"?instId=%s&bar=%s", instId, period) // 请求路径
	httpMethod := "GET"                                                       // HTTP 方法
	host := k.c.Okx.Host
	timestamp := time.Now().UTC().Format(time.RFC3339) // UTC 时间戳
	message := timestamp + httpMethod + requestPath    //需要签名的内容
	// 创建 HTTP 请求
	req, err := http.NewRequest(httpMethod, host+requestPath, nil) // 替换为目标 URL
	req.Close = true
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}
	client := k.c.Okx.SetHeaderAndProxy(req, timestamp, message) //设置请求头部
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
		return
	}
	defer resp.Body.Close()

	// 读取和处理响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading HTTP response: %v", err)
	}

	fmt.Println("HTTP Response:", string(body))
}