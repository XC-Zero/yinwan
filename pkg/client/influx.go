package client

import (
	"context"
	"fmt"

	cfg "github.com/XC-Zero/yinwan/pkg/config"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func InitInflux(config cfg.InfluxConfig) {

	url := fmt.Sprintf("http://%s:%s", config.Host, config.Port)

	client := influxdb2.NewClient(url, config.Token)

	_, err := client.Ready(context.TODO())
	if err != nil {
		panic(err)
	}
	InfluxDBClient = &client
}
