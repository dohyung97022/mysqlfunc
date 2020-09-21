package mysqlfunc

import (
	"fmt"
	"sync"
	"testing"
)

func gorutineWorker() {
	var wg sync.WaitGroup
	for i := 0; i < 99; i++ {
		fmt.Println("Main: Starting worker", i)
		wg.Add(1)
		go worker(&wg, i)
	}
	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Main: Completed")
}
func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("Worker %v: Started\n", id)
	queryStr := "INSERT INTO test () VALUES (); SET @last_id = LAST_INSERT_ID(); INSERT INTO test2 (id) VALUES (@last_id);"
	err = ExecQuery(queryStr)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	fmt.Printf("Worker %v: Finished\n", id)
}

func TestMain(t *testing.T) {

	//always delete password before save!!!!
	id := "dohyung97022"
	ps := ""
	endpoint := "adiy-db.cxdzwqqcqoib.us-east-1.rds.amazonaws.com"
	port := 3306
	schema := "adiy"

	err := Init(id, ps, endpoint, port, schema)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	// defer DB.Close()
	// // queryStr := "SELECT chan_id,channel,chan_url FROM channels"
	// // select data1 data2 data3 FROM TEST1
	var whereArray []Where
	whereArray = append(whereArray, Where{a: "chan_id", is: "=", b: 1})

	v, err := GetDataOfWhere("channels", []string{"chan_id", "channel", "last_update"},
		// []string{"last_update < " + "'" + time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05") + "'"})
		whereArray)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	fmt.Printf("v : %v\n", v)

	// time.Now().AddDate(0, 0, -1)
	// fmt.Printf("v : %s\n", time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
}
