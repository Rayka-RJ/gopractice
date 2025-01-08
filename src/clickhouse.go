package main

import (
	"database/sql"
	"fmt"
	"sync"
)

type DbSetting struct {
	Tcpconnection string
	Dbname string
	Tblname string
}

func (d *DbSetting) initializeDB() {
	// 1. Connect to Clickhouse
	conn, err := sql.Open("clickhouse", d.Tcpconnection)
	if err != nil {
		fmt.Printf("[DB] Failed to connect to Clickhouse: %s \n", err)
		return
	} 
	fmt.Println("Connected to ClickHouse.")
	defer conn.Close()
	
	// 2. Prepare DB & Table
	query1 := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", d.Dbname)
	if _, err = conn.Exec(query1); err != nil {
		fmt.Printf("[DB] Failed to create database: %s \n", err)
		return
	} 
	fmt.Printf("Database %s is ready. \n", d.Dbname)

	query2 := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (
		timestamp DateTime,
		value Float64
	) ENGINE = MergeTree()
	PARTITION BY toYYYYMM(Col4) 
	ORDER BY timestamp
	SETTINGS index_granularity = 8192; 
	`, d.Dbname, d.Tblname)

	if _, err = conn.Exec(query2); err != nil {
		fmt.Printf("[DB] Failed to create table: %s", err)		
		return
	}
	fmt.Printf("TABLE %s is ready. \n", d.Tblname)

}

func(d *DbSetting) insertData(conn *sql.DB, buf chan TimeData, wg *sync.WaitGroup) {
	// 0. consumer
	defer wg.Done()

	// 1. start a transaction
	scope, err := conn.Begin()
	if err != nil {
		fmt.Printf("failed to begin transaction: %s ", err)
		return
	}
	defer scope.Rollback()

	// 2. Prepare for batch insert
	query := fmt.Sprintf("INSERT INTO %s.%s (timestamp, value)", d.Dbname, d.Tblname)
	batch, err := scope.Prepare(query)
	if err != nil {
		fmt.Printf("failed to prepare batch insert: %s ", err)
		return
	}
	defer batch.Close()

	// 3. Insert data from buffer
	for data := range buf {
		_, err := batch.Exec(data.Timestamp, data.Value)
		if err != nil {
			fmt.Printf("Failed to insert data: %v \n", err)
			continue
		}
		fmt.Printf("Insert data: %+v \n", data)
	}

	// 4. commit the transaction
	if err := scope.Commit(); err != nil {
		fmt.Printf("failed to commit transaction: %s", err)
	}
}

