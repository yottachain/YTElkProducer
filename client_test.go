package YTElkProducer

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/yottachain/YTElkProducer/conf"
	"testing"
	"time"
)

type TestStruct struct {
	Aaa  string
	Bbb  int
	Date time.Time
}

func TestAddLog(t *testing.T) {
	elkConf := elasticsearch.Config{
		Addresses: []string{"elk服务器url"},
		Username:  "用户名",
		Password:  "口令",
	}
	ytESConfig := conf.YTESConfig{
		ESConf:      elkConf,
		DebugMode:   true,
		IndexPrefix: "main-net-dn",
		IndexType:   "log",
	}
	client := NewClient(ytESConfig)
	t1 := &TestStruct{
		Aaa:  "aaa",
		Bbb:  0,
		Date: time.Now(),
	}
	t2 := &TestStruct{
		Aaa:  "dd",
		Bbb:  0,
		Date: time.Now().AddDate(0, 0, 1),
	}
	t3 := &TestStruct{
		Aaa:  "aajkjkl",
		Bbb:  3345,
		Date: time.Now().Add(5 * time.Duration(time.Second)),
	}
	client.AddLogAsync(t1)
	client.AddLogAsync(t2)
	client.AddLogAsync(t3)
	time.Sleep(10 * time.Duration(time.Second))
}
