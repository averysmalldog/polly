package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// OutputCounts displays a success/failure counter, much like `ping` would
func OutputCounts(s *int, f *int) {
    var successRate float64
    if *s == 0 {
        successRate = 0
    } else if *f == 0 {
        successRate = 100
    } else {
        successRate = float64(1-(*f / *s))*100
    }
    fmt.Printf("\npolly terminated! Successful GETs: %d, Failed GETs: %d, Success Rate: %v\n",*s,*f,successRate)
}

// InfluxAsyncGet polls the Tesla Gen3 Wall Connector and writes results to
// InfluxDB with an aync, non-blocking client you supply. You must also
// supply the IP of the wall conenctor.
func InfluxAsyncGet(writeAPI *api.WriteAPI, wcIP string) {
	client := *writeAPI
    var data map[string]interface{}
    resp, err := http.Get(fmt.Sprintf("http://%s/api/1/vitals", wcIP))
    if err != nil {
        fmt.Println("error - during GET of hpwc. Do you have the right IP?")
        return
    }
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &data)
	fmt.Printf(".")
	p := influxdb2.NewPoint(
		"hpwc",
		map[string]string{
			"product":  "Gen3 HPWC",
			"vendor":   "Tesla",
			"location": "Garage",
		},
		data,
		time.Now())

	client.WritePoint(p)

}

func main() {
	hpwcIP := os.Getenv("HPWC_IP")
	client := influxdb2.NewClientWithOptions("http://localhost:8086", "my-token", influxdb2.DefaultOptions().SetBatchSize(20))
	writeAPI := client.WriteAPI("admin", "tesla")
	defer client.Close()
	defer writeAPI.Flush()
	for {
		go InfluxAsyncGet(&writeAPI, hpwcIP)
		time.Sleep(time.Millisecond * 1000)
	}
}
