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

func listObjects(t *testing.T, querySQL string, conn *sql.DB) []map[string]interface{} {
	log.Println("column info: ")

	columns, err := connect.GetColumns(conn, querySQL)
	if err != nil {
		return nil
	}

	for _, col := range columns {
		log.Println("Name: ", col.Name, "info:", converters[col.Type], "Length", col.Length)
	}

	//log.Printf("columns: %#v", columns)

	qry := querySQL

	log.Printf(`executing "%s"`, qry)
	rows, err := conn.Query(qry)

	if err != nil {
		t.Logf(`error with %q: %s`, qry, err)
		t.FailNow()
		return nil
	}
	scanFrom := make([]interface{}, len(columns))
	scanTo := make([]interface{}, len(columns))
	for i, _ := range scanFrom {
		scanFrom[i] = &scanTo[i]
	}

	returnMap := make(map[string]interface{})
	assocArray := make([]map[string]interface{}, 0)
	for rows.Next() {
		if err = rows.Scan(scanFrom...); err != nil {
			t.Errorf("error fetching: %s", err)
			break
		}
		for i, _ := range scanTo {
			returnMap[columns[i].Name] = scanTo[i]
		}
		assocArray = append(assocArray, returnMap)
	}
	return assocArray
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

			var (
				msgMap []map[string]interface{}
			)
			msgMap = listObjects(t, querySQL, conn)

			if msgMap == nil {
				CTxt.JSON(http.StatusNoContent, nil)
			} else {
				CTxt.JSON(http.StatusOK, msgMap)
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
