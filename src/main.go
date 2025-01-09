package main

import (
	"database/sql"
	"fmt"
	"sync"
	"net/http"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	default_db := DbSetting{
		Tcpconnection: "tcp://127.0.0.1:9000?username=default&password=", 
		Dbname: "random_num_records", 
		Tblname: "real_time_data",
	}
	
	// 0. Optional (Only run it for the first time)
	default_db.initializeDB()
	
	prometheus.MustRegister(rdmGauge)

	gene := Generator{
		DataAmount: 10,
		Min: -1000,
		Max: 1000,
		Interval: 1*time.Second,
		Threshold: 0,
	}

	// 1. Connect to databse
	conn, err := sql.Open("clickhouse", default_db.Tcpconnection)
	if err != nil {
		fmt.Printf("Failed to connect to: %s", err)
	} 
	fmt.Printf("Connect to db: %s \n", default_db.Dbname)

	// 2. Create Channel and WaitGroup

	buf := make(chan TimeData)
	alert := make(chan TimeData)
	var wg  sync.WaitGroup

	// 3. Start data generating & insert
	wg.Add(1)
	go gene.dataCollect(buf, alert, &wg)
	wg.Add(1)
	go default_db.insertData(conn, buf, &wg)
	wg.Add(1)
	go getAlertData(alert, &wg)

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		http.ListenAndServe(":8090", nil)
	}()

	wg.Wait()
	fmt.Println("All data has been inserted and alerts processed successfully.")
}


