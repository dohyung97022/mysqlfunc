package mysqlfunc

import (
	"fmt"
	"sync"
	"testing"
)

func TestMain(t *testing.T) {
	var wg sync.WaitGroup
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
	defer DB.Close()

	for i := 0; i < 100; i++ {
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
