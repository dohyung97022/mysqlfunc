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

	// SELECT * FROM channels_views
	// WHERE query = 'horror games'
	// AND avr_views BETWEEN 2100 AND 3000
	// AND subs BETWEEN 0 AND 3000

	search := "horror games"
	v, err := checkUpdateTime(search)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	fmt.Printf("v : %v\n", v)
	subsMinMax := []string{"0"}
	avrMinMax := []string{"2500"}
	v, err = fetchChannels(search, subsMinMax, avrMinMax)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	fmt.Printf("v : %v\n", v)

	// time.Now().AddDate(0, 0, -1)
	// fmt.Printf("v : %s\n", time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
}
func checkUpdateTime(search string) (map[int]map[string]interface{}, error) {
	whereArray := []Where{
		Where{a: "query", is: "=", b: search}}
	return GetDataOfWhere("search", []string{"last_update"}, whereArray)
}
func fetchChannels(search string, subsMinMax []string, avrMinMax []string) (map[int]map[string]interface{}, error) {
	whereArray := []Where{
		Where{a: "query", is: "=", b: search},
		Where{a: "avr_views", is: ">", b: 2000},
		Where{a: "subs", is: "=", b: 0}}
	return GetDataOfWhere("channels_views", []string{"*"}, whereArray)
}
