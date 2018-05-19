Sys.setenv(TZ = "Asia/Taipei", ORA_SDTZ = "Asia/Taipei")
library(ROracle)
library(nycflights13)

getOraConn <- function(){
  host <- "192.168.1.113"
  port <- 1539
  sid <- "orcl"
  connectString <- sprintf("(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=%s)(PORT=%i))(CONNECT_DATA=(SID=%s)))",
                           host, port, sid)
  dbConnect(dbDriver("Oracle"), username = "webdbtool",
            password = "f15cim2w3e4r", dbname = connectString)
}

oraConn <- getOraConn()
dbWriteTable(oraConn, "AIRLINES", airlines)
dbWriteTable(oraConn, "FLIGHTS", flights)
dbWriteTable(oraConn, "PLANES", planes)
dbWriteTable(oraConn, "AIRPORTS", airports)
dbWriteTable(oraConn, "WEATHER", weather)
dbDisconnect(oraConn)
