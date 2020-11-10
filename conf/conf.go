package conf

import "github.com/elastic/go-elasticsearch/v8"

type YTESConfig struct {
	ESConf      elasticsearch.Config
	IndexPrefix string
	IndexType   string
	DebugMode   bool
}
