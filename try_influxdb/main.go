package main

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"os"
	"time"
)

func main() {
	token := os.Getenv("-nqZ7ovL8FdB7P8lHlyIfjem6cVeAm6l63LT2vyb5w46eYgTa8s_spSWeHjgC9ugWSNSIYx4APb3Shg2Wb-7-g==")
	url := "http://60.204.241.30:8086"
	client := influxdb2.NewClient(url, token)
	//write1(client)
	query1(client, "ledgerhhh")
}

func query1(client influxdb2.Client, org string) {
	queryAPI := client.QueryAPI(org)
	query := `from(bucket: "hhh")
            |> range(start: -10000m)
            |> filter(fn: (r) => r._measurement == "measurement1")`

	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
}

func write1(client influxdb2.Client) {
	org := "ledgerhhh"
	bucket := "test-go"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	for value := 0; value < 5; value++ {
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"field1": value,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}
	}
}
