package mysqlfunc

import (
	"fmt"
	"strings"
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
	defer DB.Close()

	// whereArray := []Where{
	// 	Where{a: "last_update", is: "<", b: time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05")},
	// 	Where{a: "chan_id", is: "<", b: 3},
	// 	Where{a: "channel", is: "=", b: "/channel/UCwFl9Y49sWChrddQTD9QhRA"}}
	// v, err := GetDataOfWhere("channels", []string{"chan_id", "channel", "last_update"}, whereArray)
	// if err != nil {
	// 	fmt.Printf("error : %v\n", err)
	// }
	// fmt.Printf("v : %v\n", v)
	search := "horror games"
	subsMinMax := []string{"0"}
	avrMinMax := []string{"2500"}
	v, err := fetchADIY(search, subsMinMax, avrMinMax)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	fmt.Printf("v : %v\n", v)

	// time.Now().AddDate(0, 0, -1)
	// fmt.Printf("v : %s\n", time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
}

func fetchADIY(search string, subsMinMax []string, avrMinMax []string) (map[int]map[string]interface{}, error) {
	var q strings.Builder
	q.WriteString("SELECT c.chan_id, channel, subs, chan_url, c.last_update, chan_img, avr_views, about FROM search s JOIN search_channels sc ON s.srch_id = sc.srch_id JOIN channels c ON sc.chan_id = c.chan_id WHERE s.query = '")
	q.WriteString(search)
	q.WriteString("'")
	if len(subsMinMax) > 0 {
		q.WriteString(" AND c.subs")
		if len(subsMinMax) == 2 {
			q.WriteString(" BETWEEN ")
			q.WriteString(subsMinMax[0])
			q.WriteString(" AND ")
			q.WriteString(subsMinMax[1])
		} else {
			q.WriteString(" >= ")
			q.WriteString(subsMinMax[0])
		}
	}
	if len(avrMinMax) > 0 {
		q.WriteString(" AND c.avr_views")
		if len(avrMinMax) == 2 {
			q.WriteString(" BETWEEN ")
			q.WriteString(avrMinMax[0])
			q.WriteString(" AND ")
			q.WriteString(avrMinMax[1])
		} else {
			q.WriteString(" >= ")
			q.WriteString(avrMinMax[0])
		}
	}
	return GetQuery(q.String())
}
