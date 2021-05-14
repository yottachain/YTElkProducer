package YTElkProducer

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/yottachain/YTElkProducer/conf"
)

func NewClient(cfg conf.YTESConfig) (Client, error) {
	es, err := elasticsearch.NewClient(cfg.ESConf)
	if err != nil {
		return nil, err
	}
	if cfg.DebugMode {
		info, err := es.Info()
		if err != nil {
			log.Println("Error getting response: ", err)
		}
		if info.IsError() {
			log.Println("Error: ", info.String())
		} else {
			data, err := json.MarshalIndent(info, "", "\t")
			if err != nil {
				log.Println("Marshal Json Error: ", err)
			} else {
				log.Println(string(data))
			}
		}
	}
	return &YTElasticSearchClient{
		es:         es,
		ytESConfig: cfg,
	}, nil
}

type YTElasticSearchClient struct {
	es         *elasticsearch.Client
	ytESConfig conf.YTESConfig
}

func (c *YTElasticSearchClient) indexName() string {
	tl := "2006-01-02"
	indexDate := time.Now().Format(tl)
	return strings.Trim(strings.Trim(c.ytESConfig.IndexPrefix, "-"), "_") + "-" + indexDate
}

func (c *YTElasticSearchClient) AddDoc(doc interface{}) {
	body, err := json.Marshal(doc)
	if err != nil {
		log.Println(err)
		return
	}
	req := esapi.IndexRequest{
		Index: c.indexName(),
		Body:  bytes.NewReader(body),
	}
	res, err := req.Do(context.Background(), c.es)
	if err != nil {
		log.Println(err)
		return
	}
	if c.ytESConfig.DebugMode {
		if res.IsError() {
			log.Println("Error parsing response body: ", err)
		}
	}
}

func (c *YTElasticSearchClient) AddDocAsync(doc interface{}) {
	go c.AddDoc(doc)
}

func (c *YTElasticSearchClient) AddLog(logBody interface{}) {
	logDoc := &LogDocument{
		Timestamp: time.Now(),
		Log:       logBody,
	}
	body, err := json.Marshal(logDoc)
	if err != nil {
		log.Println(err)
		return
	}
	req := esapi.IndexRequest{
		Index: c.indexName(),
		Body:  bytes.NewReader(body),
	}
	res, err := req.Do(context.Background(), c.es)
	if err != nil {
		log.Println(err)
		return
	}
	if c.ytESConfig.DebugMode {
		log.Println(string(body))
		if res.IsError() {
			log.Println("Error parsing response body: ", err)
		}
	}
}

func (c *YTElasticSearchClient) AddLogAsync(logBody interface{}) {
	go c.AddLog(logBody)
}
