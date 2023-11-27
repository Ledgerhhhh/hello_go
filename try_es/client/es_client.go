package client

import (
	"com.ledger.goproject/myconfig"
	"com.ledger.goproject/try_es/util"
	"github.com/elastic/go-elasticsearch/v7"
	"strconv"
)

func InitEsClient() error {

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://" +
			myconfig.GConfig.ElasticConfig.Host +
			":" +
			strconv.Itoa(myconfig.GConfig.ElasticConfig.Port)},
	})
	if err != nil {
		return err
	}

	util.EsClient = client

	return nil
}
