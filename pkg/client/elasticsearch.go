package client

//import (
//	"crypto/tls"
//	cfg "github.com/XC-Zero/yinwan/pkg/config"
//	"github.com/olivere/elastic/v7"
//	"net"
//	"net/http"
//	"time"
//)
//
//
//// InitElasticsearch ...
//func InitElasticsearch(config cfg.ESConfig) (*elastic.Client, error) {
//	elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
//	var c = elastic.Config{
//		Addresses: []string{
//			config.Host,
//		},
//		Username: config.User,
//		Password: config.Password,
//		Transport: &http.Transport{
//			MaxIdleConnsPerHost:   10,
//			ResponseHeaderTimeout: time.Duration(config.ResponseHeaderTimeoutSeconds) * time.Second,
//			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
//			TLSClientConfig: &tls.Config{
//				MinVersion:         tls.VersionTLS11,
//				InsecureSkipVerify: true,
//			},
//		},
//	}
//
//	esClient, err := elastic.NewClient(c)
//	if err != nil {
//		return nil, err
//	}
//
//	_, err = esClient.Ping()
//	if err != nil {
//		return nil, err
//	}
//
//	return esClient, nil
//}
//
//// InitElasticsearchWithoutTLS ...
//func InitElasticsearchWithoutTLS(config cfg.ESConfig) (*elastic.Client, error) {
//	var cfg = elastic.Config{
//		Addresses: []string{
//			config.Host,
//		},
//		Username: config.User,
//		Password: config.Password,
//		Transport: &http.Transport{
//			MaxIdleConnsPerHost:   10,
//			ResponseHeaderTimeout: time.Duration(config.ResponseHeaderTimeoutSeconds) * time.Second,
//			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
//		},
//	}
//
//	esClient, err := elastic.NewClient(cfg)
//	if err != nil {
//		return nil, err
//	}
//	return esClient, nil
//}
//
//func CreateIndex(indexName string) {
//	ESClient.Index.
//}
