import cx_Oracle

# local
oracleSystemLoginInfo = u'webdbtool/f15cim2w3e4r@192.168.1.113:1539/orcl.jamal.org'
oracleConn = cx_Oracle.connect(oracleSystemLoginInfo)
oracleCursor = oracleConn.cursor()
oracleCursor.execute("SELECT USERNAME FROM all_users")
orclUserNames = [x[0] for x in oracleCursor.fetchall()]
oracleCursor.close()
oracleConn.close()

# remote
import cx_Oracle
oracleSystemLoginInfo = u'webdbtool/f15cim2w3e4r@jamalvm01.ddns.net:15039/orcl.jamal.org'
oracleConn = cx_Oracle.connect(oracleSystemLoginInfo)
oracleCursor = oracleConn.cursor()
oracleCursor.execute("SELECT USERNAME FROM all_users")
orclUserNames = [x[0] for x in oracleCursor.fetchall()]
oracleCursor.close()
oracleConn.close()

