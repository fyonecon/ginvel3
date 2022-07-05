package providers

import (
	"context"
	"ginvel/internal/reader/framework_toml"
	"github.com/olivere/elastic"
	"log"
	"time"
)

type InitES struct {}

// ESClient ES助手函数
var ESClient *elastic.Client

// InitElasticSearch 初始化
// 安装请参考：https://blog.csdn.net/weixin_41827162/article/details/118436153
func (initES *InitES) InitElasticSearch() {
	var httpHost string = framework_toml.GetElasticSearchConfig()["HttpHost"]
	var port string = framework_toml.GetElasticSearchConfig()["Port"]

	var address string = httpHost + ":" + port

	defer func() {
		if r := recover(); r != nil {
			log.Println("ElasticSearch初始化出现问题，已经跳过该问题。。。")
		}
	}()

	var err error
	ESClient, err = elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(address),
		elastic.SetHealthcheckInterval(12*time.Second),
		//elastic.SetGzip(true),
		//elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		//elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	)
	if err != nil {
		log.Println("ElasticSearch连接不上（请检查es的连接地址是否正确）,", address, "，", err)
	}
	_, _, err = ESClient.Ping(address).Do(context.Background())
	if err != nil {
		log.Println("ElasticSearch连接不上a，", address, "，", err)
		//os.Exit(200)
	}else {
		log.Println("ElasticSearch已连接 >>> ", address)
	}
	//fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esVersion, err := ESClient.ElasticsearchVersion(address)
	if err != nil {
		log.Println("ElasticSearch连接不上b，", address, "，", err)
	}else {
		log.Println("Elasticsearch version=", esVersion)
	}

	return
}
