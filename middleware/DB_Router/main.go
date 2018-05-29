package main

import (
	"context"
	"database/sql"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/weizhe0422/WOT/middleware/DB_Router/connect"
	"gopkg.in/rana/ora.v4"
)

var converters = map[int]string{
	1:  "VARCHAR2",
	2:  "NUMBER",
	8:  "LONG",
	11: "ROWID",
	12: "DATE",
}

type colInfoStruct struct {
	colName string
	colType string
	colLeng int
}

func listObjects(ctx context.Context, t *testing.T, querySQL string, conn *sql.DB) ([]map[string]interface{}, []map[string]string) {
	log.Println("column info: ")

	columns, err := connect.GetColumns(conn, querySQL)
	if err != nil {
		return nil, nil
	}
	colSlice := make([]colInfoStruct, len(columns))

	for idx, col := range columns {
		//log.Println("Name: ", col.Name, "info:", converters[col.Type], "Length", col.Length)
		colSlice[idx] = colInfoStruct{colName: col.Name, colType: converters[col.Type], colLeng: col.Length}
	}

	colMap := make([]map[string]string, 0)
	for _, value := range colSlice {
		tmpMap := make(map[string]string)
		tmpMap["ColName"] = value.colName
		tmpMap["ColType"] = value.colType
		tmpMap["ColLength"] = strconv.Itoa(value.colLeng)

		colMap = append(colMap, tmpMap)
	}

	//log.Printf("columns: %#v", columns)

	qry := querySQL

	rows, err := conn.QueryContext(ctx, qry)

	if err != nil {
		t.Logf(`error with %q: %s`, qry, err)
		t.FailNow()
		return nil, nil
	}
	scanFrom := make([]interface{}, len(columns))
	scanTo := make([]interface{}, len(columns))
	for i := range scanFrom {
		scanFrom[i] = &scanTo[i]
	}

	assocArray := make([]map[string]interface{}, 0)

	defer rows.Close()
	for rows.Next() {
		returnMap := make(map[string]interface{})

		if err = rows.Scan(scanFrom...); err != nil {
			t.Errorf("error fetching: %s", err)
			break
		}
		for i := range scanTo {
			returnMap[columns[i].Name] = scanTo[i]
		}
		assocArray = append(assocArray, returnMap)
	}
	return assocArray, colMap
}

func main() {
	flag.Parse()
	t := new(testing.T)
	c := make(chan os.Signal)
	var wg sync.WaitGroup

	conn, _, err := connect.GetConnection("")
	if err != nil {
		t.Errorf("error connectiong: %s", err)
		t.FailNow()
	}
	defer conn.Close()

	router := gin.Default()

	router.GET("hello/:usr", func(CTxt *gin.Context) {
		name := CTxt.Param("usr")
		CTxt.String(http.StatusOK, "Hello %s", name)
	})

	router.POST("/query", func(CTxt *gin.Context) {
		querySQL := CTxt.PostForm("querySql")
		fetchSize, _ := strconv.Atoi(CTxt.PostForm("fetchSize"))
		AppInfo := CTxt.PostForm("AppInfo")

		wg.Add(1)
		go func() {
			//log.Printf("waiting for signal...")
			//sig := <-c
			//log.Printf("got signal %s", sig)

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			ctx = ora.WithStmtCfg(ctx, ora.Cfg().StmtCfg.SetPrefetchRowCount(uint32(fetchSize)))
			defer cancel()

			log.Println("querySQL:", querySQL)

			var (
				msgMap []map[string]interface{}
				colMap []map[string]string
			)
			msgMap, colMap = listObjects(ctx, t, querySQL, conn)

			returnMsg := make(map[string]interface{}, 0)
			returnMsg["AppInfo"] = AppInfo
			returnMsg["ColInfo"] = colMap
			returnMsg["RawData"] = msgMap

			if msgMap == nil {
				CTxt.JSON(http.StatusNoContent, nil)
			} else {
				CTxt.JSON(http.StatusOK, returnMsg)
			}

			/*CTxt.JSON(200, gin.H{
				"usrname": ,
			})*/
			wg.Done()
		}()
		signal.Notify(c, syscall.SIGUSR1)
		wg.Wait()
	})

	router.Run(":8080")
	//log.Fatal(gateway.ListenAndServe(":8080", router))

}
