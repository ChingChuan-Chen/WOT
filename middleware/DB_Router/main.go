package main

import (
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
	"github.com/weizhe0422/WOT_v2/connect"
)

func listObjects(t *testing.T, querySQL string) []string {
	conn, err := connect.GetConnection("")
	if err != nil {
		t.Errorf("error connectiong: %s", err)
		t.FailNow()
		return nil
	}
	defer conn.Close()

	//"SELECT USERNAME FROM all_users"
	//"SELECT owner, object_name, object_id FROM all_objects WHERE ROWNUM < 20"
	qry := querySQL

	log.Printf(`executing "%s"`, qry)
	rows, err := conn.Query(qry)

	if err != nil {
		t.Logf(`error with %q: %s`, qry, err)
		t.FailNow()
		return nil
	}

	var (
		userName string
		results  []string
	)

	for rows.Next() {
		if err = rows.Scan(&userName); err != nil {
			t.Errorf("error fetching: %s", err)
			break
		}
		log.Println("rows: ", userName)
		results = append(results, userName)
	}
	return results
}

func main() {
	flag.Parse()
	t := new(testing.T)
	c := make(chan os.Signal)
	var wg sync.WaitGroup

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

			CTxt.JSON(200, gin.H{
				"usrname": listObjects(t, querySQL),
			})
			wg.Done()
		}()
		signal.Notify(c, syscall.SIGUSR1)
		wg.Wait()
	})

	router.Run(":8080")
	//log.Fatal(gateway.ListenAndServe(":8080", router))

}
