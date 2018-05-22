Sys.setenv(TZ = "Asia/Taipei", ORA_SDTZ = "Asia/Taipei")
library(ROracle)
library(nycflights13)
library(pipeR)
library(data.table)

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
dbWriteTable(oraConn, "AIRLINES", airlines %>>% setDT %>>% setnames(toupper(names(.))))
dbWriteTable(oraConn, "FLIGHTS", flights %>>% setDT %>>% setnames(toupper(names(.))))
dbWriteTable(oraConn, "PLANES", planes %>>% setDT %>>% setnames(toupper(names(.))))
dbWriteTable(oraConn, "AIRPORTS", airports %>>% setDT %>>% setnames(toupper(names(.))))
dbWriteTable(oraConn, "WEATHER", weather %>>% setDT %>>% setnames(toupper(names(.))))
dbDisconnect(oraConn)
