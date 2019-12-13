package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	openprio_pt_position_data "openprio_log/openprio"
	"os"
	"time"

	// "github.com/elastic/go-elasticsearch/v8"
	// "github.com/elastic/go-elasticsearch/v8/esapi"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/golang/protobuf/proto"
)

func connect() mqtt.Client {
	opts := createClientOptions()
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions() *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()

	host := os.Getenv("MQTT_HOST")
	if host == "" {
		log.Fatal("MQTT_HOST is not set.")
	}
	opts.AddBroker(fmt.Sprintf("ssl://%s", host))

	deviceID := os.Getenv("MQTT_DEVICE_ID")
	if deviceID == "" {
		log.Fatal("MQTT_DEVICE_ID is not set.")
	}
	opts.SetUsername(deviceID)
	opts.SetClientID(deviceID)

	password := os.Getenv("MQTT_PASSWORD")
	if password == "" {
		log.Fatal("MQTT_PASSWORD is not set.")
	}
	opts.SetPassword(password)

	return opts
}

func listen(topic string, ch chan openprio_pt_position_data.LocationMessage) {
	client := connect()
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		newTest := openprio_pt_position_data.LocationMessage{}
		proto.Unmarshal(msg.Payload(), &newTest)

		fmt.Printf("* [%s] %v\n", msg.Topic(), newTest)
		ch <- newTest

	})
}

func main() {
	input := make(chan openprio_pt_position_data.LocationMessage)
	listen("#", input)
	msgs := []openprio_pt_position_data.LocationMessage{}
        //last_bulk_import := time.time()
	for {
		data := <-input
		msgs = append(msgs, data)
		if len(msgs) > 50 {
			saveData(msgs)
			msgs = []openprio_pt_position_data.LocationMessage{}
		}

	}
}

func saveData(data []openprio_pt_position_data.LocationMessage) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://167.71.67.173/elastic",
		},
	}

	indexName := "openprio_pt"
	var (
		buf bytes.Buffer
		res *esapi.Response
	)
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err = es.Indices.Exists([]string{"openprio_pt"})
	if err != nil {
		log.Println(err)
	} else {
		res, err = es.Indices.Create("openprio_pt")
		if err != nil {
			log.Fatalf("Cannot create index: %s", err)
		}
	}

	for index, content := range data {
		// Prepare the metadata payload
		//
		//meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%d" } }%s`, uuid.New(), "\n"))
		meta := []byte(fmt.Sprintf(`{  "index" : { "_id" : null } }%s`, "\n"))

		// fmt.Printf("%s", meta) // <-- Uncomment to see the payload

		// Prepare the data payload: encode article to JSON
		//
		type Test struct {
			Time    time.Time                                 `json:"time"`
			Content openprio_pt_position_data.LocationMessage `json:"content"`
		}

		test := Test{}
		test.Content = content
		test.Time = time.Unix(content.Timestamp/1000, (content.Timestamp%1000)*1000000)

		data, err := json.Marshal(test)
		if err != nil {
			log.Println(err)
			log.Fatalf("Cannot encode article %d: %s", index, '\n')
		}

		// Append newline to the data payload
		//
		data = append(data, "\n"...) // <-- Comment out to trigger failure for batch
		// fmt.Printf("%s", data) // <-- Uncomment to see the payload

		// // Uncomment next block to trigger indexing errors -->
		// if a.ID == 11 || a.ID == 101 {
		// 	data = []byte(`{"published" : "INCORRECT"}` + "\n")
		// }
		// // <--------------------------------------------------

		// Append payloads to the buffer (ignoring write errors)
		//
		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)
	}

	res, err = es.Bulk(bytes.NewReader(buf.Bytes()), es.Bulk.WithIndex(indexName))
	if err != nil {
		log.Fatalf("Failure indexing batch")
	}
	// If the whole request failed, print error and mark all documents as failed
	//
	if res.IsError() {
		log.Println("Error")
		var raw map[string]interface{}
		json.NewDecoder(res.Body).Decode(&raw)
		log.Println("%+v", raw)

	}
	res.Body.Close()

	buf.Reset()
	log.Println("Import complete")

}

// type Article struct {
// 	ID        int       `json:"id"`
// 	Title     string    `json:"title"`
// 	Body      string    `json:"body"`
// 	Published time.Time `json:"published"`
// 	Author    Author    `json:"author"`
// }

// type Author struct {
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// }

// var (
// 	_     = fmt.Print
// 	count int
// 	batch int
// )

// func init() {
// 	flag.IntVar(&count, "count", 1000, "Number of documents to generate")
// 	flag.IntVar(&batch, "batch", 255, "Number of documents to send in one batch")
// 	flag.Parse()

// 	rand.Seed(time.Now().UnixNano())
// }

// func main() {
// 	log.SetFlags(0)

// 	type bulkResponse struct {
// 		Errors bool `json:"errors"`
// 		Items  []struct {
// 			Index struct {
// 				ID     string `json:"_id"`
// 				Result string `json:"result"`
// 				Status int    `json:"status"`
// 				Error  struct {
// 					Type   string `json:"type"`
// 					Reason string `json:"reason"`
// 					Cause  struct {
// 						Type   string `json:"type"`
// 						Reason string `json:"reason"`
// 					} `json:"caused_by"`
// 				} `json:"error"`
// 			} `json:"index"`
// 		} `json:"items"`
// 	}

// 	var (
// 		buf bytes.Buffer
// 		res *esapi.Response
// 		err error
// 		raw map[string]interface{}
// 		blk *bulkResponse

// 		articles  []*Article
// 		indexName = "articles"

// 		numItems   int
// 		numErrors  int
// 		numIndexed int
// 		numBatches int
// 		currBatch  int
// 	)

// 	// Create the Elasticsearch client
// 	//

// 	cfg := elasticsearch.Config{
// 		Addresses: []string{
// 			"http://167.71.67.173/elastic",
// 		},
// 	}
// 	es, err := elasticsearch.NewClient(cfg)
// 	if err != nil {
// 		log.Fatalf("Error creating the client: %s", err)
// 	}

// 	// Generate the articles collection
// 	//
// 	names := []string{"Alice", "John", "Mary"}
// 	for i := 1; i < count+1; i++ {
// 		articles = append(articles, &Article{
// 			ID:        i,
// 			Title:     strings.Join([]string{"Title", strconv.Itoa(i)}, " "),
// 			Body:      "Lorem ipsum dolor sit amet...",
// 			Published: time.Now().Round(time.Second).UTC().AddDate(0, 0, i),
// 			Author: Author{
// 				FirstName: names[rand.Intn(len(names))],
// 				LastName:  "Smith",
// 			},
// 		})
// 	}
// 	log.Printf("> Generated %d articles", len(articles))

// 	// Re-create the index
// 	//
// 	if res, err = es.Indices.Delete([]string{indexName}); err != nil {
// 		log.Fatalf("Cannot delete index: %s", err)
// 	}
// 	res, err = es.Indices.Create(indexName)
// 	if err != nil {
// 		log.Fatalf("Cannot create index: %s", err)
// 	}
// 	if res.IsError() {
// 		log.Fatalf("Cannot create index: %s", res)
// 	}

// 	if count%batch == 0 {
// 		numBatches = (count / batch)
// 	} else {
// 		numBatches = (count / batch) + 1
// 	}

// 	start := time.Now().UTC()

// 	// Loop over the collection
// 	//
// 	for i, a := range articles {
// 		numItems++

// 		currBatch = i / batch
// 		if i == count-1 {
// 			currBatch++
// 		}

// 		// Prepare the metadata payload
// 		//
// 		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%d" } }%s`, a.ID, "\n"))
// 		// fmt.Printf("%s", meta) // <-- Uncomment to see the payload

// 		// Prepare the data payload: encode article to JSON
// 		//
// 		data, err := json.Marshal(a)
// 		if err != nil {
// 			log.Fatalf("Cannot encode article %d: %s", a.ID, err)
// 		}

// 		// Append newline to the data payload
// 		//
// 		data = append(data, "\n"...) // <-- Comment out to trigger failure for batch
// 		// fmt.Printf("%s", data) // <-- Uncomment to see the payload

// 		// // Uncomment next block to trigger indexing errors -->
// 		// if a.ID == 11 || a.ID == 101 {
// 		// 	data = []byte(`{"published" : "INCORRECT"}` + "\n")
// 		// }
// 		// // <--------------------------------------------------

// 		// Append payloads to the buffer (ignoring write errors)
// 		//
// 		buf.Grow(len(meta) + len(data))
// 		buf.Write(meta)
// 		buf.Write(data)

// 		// When a threshold is reached, execute the Bulk() request with body from buffer
// 		//
// 		if i > 0 && i%batch == 0 || i == count-1 {
// 			log.Printf("> Batch %-2d of %d", currBatch, numBatches)

// 			res, err = es.Bulk(bytes.NewReader(buf.Bytes()), es.Bulk.WithIndex(indexName))
// 			if err != nil {
// 				log.Fatalf("Failure indexing batch %d: %s", currBatch, err)
// 			}
// 			// If the whole request failed, print error and mark all documents as failed
// 			//
// 			if res.IsError() {
// 				numErrors += numItems
// 				if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
// 					log.Fatalf("Failure to to parse response body: %s", err)
// 				} else {
// 					log.Printf("  Error: [%d] %s: %s",
// 						res.StatusCode,
// 						raw["error"].(map[string]interface{})["type"],
// 						raw["error"].(map[string]interface{})["reason"],
// 					)
// 				}
// 				// A successful response might still contain errors for particular documents...
// 				//
// 			} else {
// 				if err := json.NewDecoder(res.Body).Decode(&blk); err != nil {
// 					log.Fatalf("Failure to to parse response body: %s", err)
// 				} else {
// 					for _, d := range blk.Items {
// 						// ... so for any HTTP status above 201 ...
// 						//
// 						if d.Index.Status > 201 {
// 							// ... increment the error counter ...
// 							//
// 							numErrors++

// 							// ... and print the response status and error information ...
// 							log.Printf("  Error: [%d]: %s: %s: %s: %s",
// 								d.Index.Status,
// 								d.Index.Error.Type,
// 								d.Index.Error.Reason,
// 								d.Index.Error.Cause.Type,
// 								d.Index.Error.Cause.Reason,
// 							)
// 						} else {
// 							// ... otherwise increase the success counter.
// 							//
// 							numIndexed++
// 						}
// 					}
// 				}
// 			}

// 			// Close the response body, to prevent reaching the limit for goroutines or file handles
// 			//
// 			res.Body.Close()

// 			// Reset the buffer and items counter
// 			//
// 			buf.Reset()
// 			numItems = 0
// 		}
// 	}

// 	// Report the results: number of indexed docs, number of errors, duration, indexing rate
// 	//
// 	log.Println(strings.Repeat("=", 80))

// 	dur := time.Since(start)

// 	if numErrors > 0 {
// 		log.Fatalf(
// 			"Indexed [%d] documents with [%d] errors in %s (%.0f docs/sec)",
// 			numIndexed,
// 			numErrors,
// 			dur.Truncate(time.Millisecond),
// 			1000.0/float64(dur/time.Millisecond)*float64(numIndexed),
// 		)
// 	} else {
// 		log.Printf(
// 			"Sucessfuly indexed [%d] documents in %s (%.0f docs/sec)",
// 			numIndexed,
// 			dur.Truncate(time.Millisecond),
// 			1000.0/float64(dur/time.Millisecond)*float64(numIndexed),
// 		)
// 	}
// }
