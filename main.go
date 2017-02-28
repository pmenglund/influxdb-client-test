package main

import (
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

func ping(c client.Client) (string, error) {
	_, pong, err := c.Ping(time.Second)
	if err != nil {
		return "", err
	}
	return pong, nil
}

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		log.Fatal(err)
	}

	version, err := ping(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)
}
