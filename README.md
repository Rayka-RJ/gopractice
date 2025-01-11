# gopractice

本项目于WSL环境下开发。

使用golang实时生成随机数（float64）存入clickhouse，并在随机数大于阈值时上报clickhouse。

- 可自定义生成随机数：amount， min， max， intervals， threshold
- 可自定义存入数据库：tcpconnection（default=127.0.0.1:9000），dbname，tblname
- 设定prometheus监听端口为8090，程序运行时，[localhost:8090](http://localhost:8090/metrics) 实时生成监听metrics
- 基于prometheus和grafana的可视化端口均为默认值（Prometheus：9090；Grafana：3000）

2025.01.11: 更改端口为docker ip host.docker.internal

一个运行样本如下：
```
Connected to ClickHouse.
Database random_num_records is ready. 
TABLE real_time_data is ready. 
Connect to db: random_num_records 
Insert data: {Timestamp:2025-01-09 21:34:25.056369989 +0800 +08 m=+0.005345866 Value:-260.3972612294658} 
Generated: {Timestamp:2025-01-09 21:34:25.056369989 +0800 +08 m=+0.005345866 Value:-260.3972612294658} 
Insert data: {Timestamp:2025-01-09 21:34:26.876728801 +0800 +08 m=+1.006053313 Value:-235.9260478807132} 
Generated: {Timestamp:2025-01-09 21:34:26.876728801 +0800 +08 m=+1.006053313 Value:-235.9260478807132} 
Generated: {Timestamp:2025-01-09 21:34:27.877898153 +0800 +08 m=+2.007222665 Value:-284.2037864379623} 
......
Insert data: {Timestamp:2025-01-09 21:34:34.882465781 +0800 +08 m=+9.011790292 Value:-269.49272024750724} 
All data has been inserted and alerts processed successfully.

```

一个metrics样本如下(阈值为0)：

```
# HELP real_time_data Randomly generated value (float64)
# TYPE real_time_data gauge
real_time_data{timestamp="2025-01-09T21:33:16.85095782+08:00"} 979.4417755519244
real_time_data{timestamp="2025-01-09T21:33:17.852219095+08:00"} 642.3719597056488
real_time_data{timestamp="2025-01-09T21:33:20.67352618+08:00"} 420.74495311236024
real_time_data{timestamp="2025-01-09T21:33:24.675530517+08:00"} 204.228750105887
real_time_data{timestamp="2025-01-09T21:33:25.675664103+08:00"} 571.0490927396743
real_time_data{timestamp="2025-01-09T21:33:27.677761742+08:00"} 427.3574340464954
real_time_data{timestamp="2025-01-09T21:33:28.678871817+08:00"} 784.8105319439703
real_time_data{timestamp="2025-01-09T21:33:29.679522879+08:00"} 763.0027409710362
```
