# gopractice

本项目于WSL环境下开发。

使用golang实时生成随机数（float64）存入clickhouse，并在随机数大于阈值时上报clickhouse。

- 可自定义生成随机数：amount， min， max， intervals， threshold
- 可自定义存入数据库：tcpconnection（default=127.0.0.1:9000），dbname，tblname
- 设定prometheus监听端口为8090，程序运行时，[localhost:8090](http://localhost:8090/metrics) 实时生成监听metrics
- 基于prometheus和grafana的可视化端口均为默认值（Prometheus：9090；Grafana：3000）

**Attention**: 该程序为实时数据监测，interval的设置必须与prometheus的fetch frequency保持一致。

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
metric设有两个label，value记录了实时生成的值，alert显示了该值是否超出阈值（0为safe，1为alert）。

一个metrics样本如下(阈值为0)：

```
# HELP real_time_data Randomly generated value (float64)
# TYPE real_time_data gauge
real_time_data{Value="Alert"} 0
real_time_data{Value="Value"} -173.3504388717426
```
