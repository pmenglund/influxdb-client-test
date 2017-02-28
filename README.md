# influxdb-client-test

A sample repo that shows how to use a drop-in replacement for the InfluxDB Go [client](https://godoc.org/github.com/influxdata/influxdb/client/v2#Client).

Since the InfluxDB client defines an interface for the `Client` it is easy to drop-in a mock to use in your unit tests.
