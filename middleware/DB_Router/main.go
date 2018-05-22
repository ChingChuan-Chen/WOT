package main

import (
	"database/sql"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"

	//"github.com/apex/gateway"
	"github.com/weizhe0422/WOT/middleware/DB_Router/connect"
)

var converters = map[int]string{
	1:  "VARCHAR2",
	2:  "NUMBER",
	8:  "LONG",
	11: "ROWID",
	12: "DATE",
}

func listObjects(t *testing.T, querySQL string, conn *sql.DB) ([]string, []string) {
	var (
		userName   string
		rowResults []string
		colResults []string
	)

	log.Println("column info: ")
	columns, err := connect.GetColumns(conn, querySQL)
	for _, col := range columns {
		log.Println("Name: ", col.Name, "info:", converters[col.Type], "Length", col.Length)
		colResults = append(colResults, col.Name)
	}

	if err != nil {
		//return fmt.Sprint("get column converters")
	}
	//log.Printf("columns: %#v", columns)

	//"SELECT USERNAME FROM all_users"
	//"SELECT owner, object_name, object_id FROM all_objects WHERE ROWNUM < 20"
	qry := querySQL

	log.Printf(`executing "%s"`, qry)
	rows, err := conn.Query(qry)

	if err != nil {
		t.Logf(`error with %q: %s`, qry, err)
		t.FailNow()
		return nil, nil
	}

	for rows.Next() {
		if err = rows.Scan(&userName); err != nil {
			t.Errorf("error fetching: %s", err)
			break
		}
		log.Println("rows: ", userName)
		rowResults = append(rowResults, userName)
	}
	return colResults, rowResults
}

func main() {
	flag.Parse()
	t := new(testing.T)
	c := make(chan os.Signal)
	var wg sync.WaitGroup

	conn, err := connect.GetConnection("")
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

		wg.Add(1)
		go func() {
			//log.Printf("waiting for signal...")
			//sig := <-c
			//log.Printf("got signal %s", sig)
			log.Println("querySQL:", querySQL)

			mapResult := make(map[string][]string)
			var (
				colInfo  []string
				dataInfo []string
			)
			colInfo, dataInfo = listObjects(t, querySQL, conn)

			for _, colName := range colInfo {
				mapResult[colName] = dataInfo
			}
			CTxt.JSON(200, mapResult)

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
