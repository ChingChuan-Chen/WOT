# WOT - Web Oracle Tool

##  設計目標
Web Oracle Tool (下面簡稱WOT)目標為以Oracle Database為核心，構建一系列的web services及web views，並提供自動部署、監控、主動即時的Alarm系統、Admin頁面以及重要元件HA機制，所形成的一個完整解決方案。

## 架構圖

![](images/架構圖.svg)

## 系統需求 (許願中)

1.  Web Services
  a. 待定

2.  Web View
  a. [ ] 支援LDAP登入，並以username控制可以查詢的DB (如目前的JASDA)
  b. SQL編輯器，支援function, table, column提示
     i. [ ] syntax highlight
     ii. [ ] theme selection (bright or dark)
     iii. [ ] 右鍵view table, query data, edit data
     iv. [ ] function提示
     v. [ ] table提示
     vi. [ ] column提示
  c. procedure, function編輯器，並支援儲存
     i. [ ] view procedure/function
     ii. [ ] edit procedure/function
  d. select SQL查詢，並顯示資料於grid
     i. [ ] view
     ii. [ ] lock (see [lock sql](https://docs.oracle.com/cd/E17952_01/mysql-5.6-en/lock-tables.html))
     iii. [ ] perform insert / update on grid
  e. [ ] 執行transaction SQL，並提供commit按鈕
  f. [ ] 執行procedure
  g. 查看table schema
    i. basic info (owner, name, tablespace)
      (1) [ ] view
    ii. columns info
      (1) [ ] view
      (2) [ ] edit
    iii. key info (see [PRIMARY KEY](https://www.w3schools.com/sql/sql_primarykey.asp), [FOREIGN KEY](https://www.w3schools.com/sql/sql_foreignkey.asp))
    iv. index info
      (1) [ ] view
      (2) [ ] edit
    v. checks info (see [CHECK](https://www.w3schools.com/sql/sql_check.asp))
      (1) [ ] view
      (2) [ ] edit
    vi. table privilege
      (1) [ ] view
      (2) [ ] edit

3.  Admin頁面
  a. [ ] 可以設定group, connection部分
  b. [ ] Web View權限管理 (卡user view/edit權限)
  c. [ ] 查詢locked table
  d. [ ] 砍特定session (for long run, table lock等)

4.  系統服務監控
  a. [ ] 暫定搭配zabbix，定時確認service存活狀態，然後寄信並自動重啟，設定可export，也可一次性部署到到各台

5.  主動即時的Alarm系統
  a. [ ] 暫定用grafana，可以拉KPI圖表並設定alarm mail，其設定可以export跟import

6.  元件HA機制
  a. [ ] middleware採group方式搭配nginx and GDNS
  b. [ ] 使用web server多台搭配nginx for 2. and 3.
  c. [ ] prograsql with high availability for server logging, service logging, connection setting and group setting

7.  自動部署
  a. [ ] 使用shell撰寫，自動安裝並啟動服務 for 1, 2, 3
  b. [ ] 自動部署6-c，並initialize tables
