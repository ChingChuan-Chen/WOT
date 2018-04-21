Sys.setenv(PATH = paste0(Sys.getenv("PATH"), ";C:\\Oracle\\instantclient_12_2"))
library(ROracle)
library(data.table)
library(pipeR)
Sys.setenv(TZ = "Asia/Taipei", ORA_SDTZ = "Asia/Taipei")

# local
host <- "192.168.1.113"
port <- 1539
sid <- "orcl"
connectString <- sprintf("(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=%s)(PORT=%i))(CONNECT_DATA=(SID=%s)))",
                         host, port, sid)
oraConn <- dbConnect(dbDriver("Oracle"), username = "webdbtool",
                     password = "f15cim2w3e4r", dbname = connectString)
dbListTables(oraConn)
dbDisconnect(oraConn)

# remote
host <- "jamalvm01.ddns.net"
port <- 15039
sid <- "orcl"
connectString <- sprintf("(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=%s)(PORT=%i))(CONNECT_DATA=(SID=%s)))",
                         host, port, sid)
oraConn <- dbConnect(dbDriver("Oracle"), username = "webdbtool",
                     password = "f15cim2w3e4r", dbname = connectString)
dbListTables(oraConn)
dbDisconnect(oraConn)


