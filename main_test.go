package main

import (
	"errors"
	"testing"
	"time"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/influxdata/influxdb/models"
)

type dummyClient struct {
	pingFn func(timeout time.Duration) (time.Duration, string, error)
}

func (dc dummyClient) Ping(timeout time.Duration) (time.Duration, string, error) {
	if dc.pingFn != nil {
		return dc.pingFn(timeout)
	}
	return time.Second, "pong", nil
}

func (dc dummyClient) Write(bp client.BatchPoints) error {
	return nil
}

func (dc dummyClient) Query(q client.Query) (*client.Response, error) {
	return &client.Response{
		Results: []client.Result{
			{
				Series: []models.Row{
					{
						Name:   "databases",
						Values: [][]interface{}{{"db1", "db2"}},
					},
				},
			},
		},
	}, nil
}

func (dc dummyClient) Close() error {
	return nil
}

func TestPing(t *testing.T) {
	c := dummyClient{}
	p, _ := ping(c)
	if p != "pong" {
		t.Fail()
	}

	c = dummyClient{
		pingFn: func(timeout time.Duration) (time.Duration, string, error) {
			return time.Second, "", errors.New("ping")
		},
	}
	_, err := ping(c)
	if err == nil {
		t.Fail()
	}
}
