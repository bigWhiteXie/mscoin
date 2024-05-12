package market

import (
	"encoding/json"
	"fmt"
	"io"
	"job-center/internal/config"
	"job-center/internal/dao"
	"job-center/internal/database"
	"job-center/internal/domain"
	"job-center/internal/model"
	"job-center/internal/svc"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// kline的定时任务领域
type Kline struct {
	wg          *sync.WaitGroup
	c           *config.Config
	klineDomain *domain.KlineDomain
	queueDomain *domain.QueueDomain
	reqPath     string
}

func NewKline(svr *svc.ServiceContext) *Kline {
	klineDao := dao.NewKlineDao(database.ConnectMongo(svr.Config).Db)
	klineDomain := domain.NewKlineDomain(klineDao)
	queueDomain := domain.NewQueueDomain(svr)
	return &Kline{
		wg:          &sync.WaitGroup{},
		c:           svr.Config,
		klineDomain: klineDomain,
		queueDomain: queueDomain,
		reqPath:     "/api/v5/market/candles",
	}
}

func (k *Kline) Do(period string) {
	log.Println("============启动k线数据拉取==============")
	size := len(k.c.Symbols)
	k.wg.Add(size)
	for _, symbol := range k.c.Symbols {
		log.Printf("===============拉取%s数据===============\n", symbol)
		go k.fetchKlineAndPush(symbol, strings.Replace(symbol, "-", "/", 1), period)
	}
	k.wg.Wait()
	log.Println("===============k线数据拉取结束===============")
}

func (k *Kline) fetchKlineAndPush(instId string, symbol string, period string) {
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading HTTP response: %v", err)
	}
	klineRes := &model.OkxKlineRes{}
	if err := json.Unmarshal(body, klineRes); err != nil {
		log.Printf(err.Error())
		return
	}
	if err = k.klineDomain.Save(klineRes.Data, symbol, period); err != nil {
		log.Printf("save error: " + err.Error())
		return
	}
	fmt.Println("HTTP Response:", string(body))
	for _, data := range klineRes.Data {
		k.queueDomain.PushKline(data, symbol, period)
	}

}
