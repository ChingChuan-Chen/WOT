# WOT - Web Oracle Tool

##  設計目標
Web Oracle Tool (下面簡稱WOT)目標為以Oracle Database為核心，構建一系列的web services及web views，並提供自動部署、監控、主動即時的Alarm系統、Admin頁面以及重要元件HA機制，所形成的一個完整解決方案。

## 架構圖

![](images/架構圖.svg)

## 系統需求 (許願中)

###  Web Services
  a. 待定

###  Web View
  - [ ] 支援LDAP登入，並以username控制可以查詢的DB (如目前的JASDA)
  - SQL編輯器，支援function, table, column提示
     - [ ] syntax highlight
     - [ ] theme selection (bright or dark)
     - [ ] 右鍵view table, query data, edit data
     - [ ] function提示
     - [ ] table提示
     - [ ] column提示
     - [ ] 查詢execusion plan
  - procedure, function編輯器，並支援儲存
     - [ ] view procedure/function
     - [ ] edit procedure/function
  - select SQL查詢，並顯示資料於grid
     - [ ] view
     - [ ] lock (see [lock sql](https://docs.oracle.com/cd/E17952_01/mysql-5.6-en/lock-tables.html))
     - [ ] perform insert / update on grid
  - [ ] 執行transaction SQL，並提供commit按鈕
  - [ ] 執行procedure
  - 查看table schema
    - basic info (owner, name, tablespace)
      - [ ] view
    - columns info
      - [ ] view
      - [ ] edit
    - key info (see [PRIMARY KEY](https://www.w3schools.com/sql/sql_primarykey.asp), [FOREIGN KEY](https://www.w3schools.com/sql/sql_foreignkey.asp))
      - [ ] view
      - [ ] edit
    - index info
      - [ ] view
      - [ ] edit
    - checks info (see [CHECK](https://www.w3schools.com/sql/sql_check.asp))
      - [ ] view
      - [ ] edit
    - table privilege
      - [ ] view
      - [ ] edit
    - partition info
      - [ ] view
      - [ ] edit

###  Admin頁面
  - [ ] 可以設定group, connection部分
  - [ ] Web View權限管理 (卡user view/edit權限)
  - [ ] 查詢locked table
  - [ ] 砍特定session (for long run, table lock等)

###  系統服務監控
  - [ ] Use Grafana, InfluxDB and Telegraf to monitoring server and web app

###  主動即時的Alarm系統
  - [ ] Use Grafana alerting
  - [ ] 自動砍table lock (可以在Admin頁面設定時間)

###  元件HA機制
  - [ ] Middleware Group
  - [ ] web server
  - [ ] Configuration of connection setting and group setting in Oracle
  - Database with high availability for server monitoring, service monitoring and logging
    - [X] solution survey - Grafana + InfluxDB + Telegraf
    - [ ] build-up

###  自動部署
  - [ ] 使用shell撰寫，自動安裝並啟動服務 for 1, 2, 3
  - [ ] 自動部署6-c，並initialize tables

### 測試Oracle部署
  - [X] 安裝: [My Blogger](http://chingchuan-chen.github.io/posts/201607/2016-07-24-deployment-of-oracle-database.html)
  - [ ] 倒入測試資料 [IMDb Datasets](https://www.imdb.com/interfaces/)

## References
  1. Monitoring System
    1. [Grafana + InfluxDB + Telegraf](https://runnerlee.com/2017/08/18/influxdb-telegraf-grafana-monitor)
    1. [Grafana + InfluxDB + Telegraf](https://github.com/anryko/grafana-influx-dashboard)
    1. [ProxySQL Monitoring Solution](http://seanlook.com/2017/07/16/mysql-proxysql-monitor/)
    1. [MySQL Monitoring Solution](https://hackernoon.com/mysql-monitoring-with-telegraf-influxdb-grafana-4489e6df0220)
    1. [InfluxDB + Python monitor Web App](https://stackoverflow.com/questions/37909251/send-python-web-app-metrics-to-influxdb)
    1. [influxdb-relay for InfluxDB HA](https://github.com/influxdata/influxdb-relay)
  2. HA solution for Web Service and Web App
    1. [使用nginx+keepalived實現RESTful API的負載平衡和高可用性](https://ieevee.com/tech/2015/07/02/nginx-keepalived.html)
    1. [nginx + keepalived實現website高可用性](https://segmentfault.com/a/1190000002881132)
  3. Middleware for keeping alive
    1. [Use nginx to pass hostname of the upstream](https://serverfault.com/questions/598202/make-nginx-to-pass-hostname-of-the-upstream-when-reverseproxying)
    1. [nginx: keep alive](https://skyao.gitbooks.io/learning-nginx/content/documentation/keep_alive.html)

## Production環境架構
  1. Oracle cluster: 數台
  1. InfluxDB: 2 servers
  1. Grafana: at least 2 servers
  1. middleware: at least 3 servers per group
  1. web view server: at least 3 servers
  1. admin web server: at least 3 servers

## Test Server配置 (9 servers)：
  1. testing Oracle: 1 servers
  1. InfluxDB: 2 servers
  1. keepalived + nginx + middleware: 3 servers
  1. keepalived + nginx + web view server: 3 servers
  1. keepalived + nginx + admin web server/Grafana: 2 servers (co-exist on web view server)
