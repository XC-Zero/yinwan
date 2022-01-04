package client

import (
	"crypto/tls"
	cfg "github.com/XC-Zero/yinwan/pkg/config"
	"github.com/elastic/go-elasticsearch/v7"
	"net"
	"net/http"
	"time"
)

// InitElasticsearch ...
func InitElasticsearch(config cfg.ESConfig) {
	var cfg = elasticsearch.Config{
		Addresses: []string{
			config.Host,
		},
		Username: config.User,
		Password: config.Password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Duration(config.ResponseHeaderTimeoutSeconds) * time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic("Init ES failed ")
	}

	_, err = esClient.Ping()
	if err != nil {
		panic(err)
	}

	ESClient = esClient
}

// InitElasticsearchWithoutTLS ...
func InitElasticsearchWithoutTLS(config cfg.ESConfig) {
	var cfg = elasticsearch.Config{
		Addresses: []string{
			config.Host,
		},
		Username: config.User,
		Password: config.Password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Duration(config.ResponseHeaderTimeoutSeconds) * time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		},
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	ESClient = esClient
}
