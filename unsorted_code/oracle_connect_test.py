import cx_Oracle

def getQuery(oracleConn):
    oracleCursor = oracleConn.cursor()
    oracleCursor.execute("SELECT USERNAME FROM all_users")
    orclUserNames = [x[0] for x in oracleCursor.fetchall()]
    oracleCursor.close()
    oracleConn.close()
    return orclUserNames

if __name__== "__main__":
    dsn_tns = '(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=jamalvm01.ddns.net)(PORT=1539))(CONNECT_DATA=(SID=orcl)))'
    test_local = True
    oracleConn = cx_Oracle.connect('webdbtool', 'f15cim2w3e4r', dsn_tns)
    print(getQuery(oracleConn))

    if test_local:
        dsn_tns = '(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=jamalvm01.ddns.net)(PORT=1539))(CONNECT_DATA=(SID=orcl)))'
        oracleConn = cx_Oracle.connect('webdbtool', 'f15cim2w3e4r', dsn_tns)
        print(getQuery(oracleConn))
