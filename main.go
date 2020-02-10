package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	openprio_pt_position_data "openprio_log/openprio"
	openprio_ssm "openprio_log/openprio_ssm"
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
	opts.SetAutoReconnect(true)

	password := os.Getenv("MQTT_PASSWORD")
	if password == "" {
		log.Fatal("MQTT_PASSWORD is not set.")
	}
	opts.SetPassword(password)

	if os.Getenv("ELASTIC_ADDRESS") == "" {
		log.Fatal("No elasticAddress set.")
	}

	return opts
}

func main() {
	log.Println("Start logger.")
	positionCh := make(chan openprio_pt_position_data.LocationMessage)
	ssmCh := make(chan openprio_ssm.ExtendedSSM)
	client := connect()
	client.Subscribe("/prod/pt/position/#", 0, func(client mqtt.Client, msg mqtt.Message) {
		newTest := openprio_pt_position_data.LocationMessage{}
		proto.Unmarshal(msg.Payload(), &newTest)

		positionCh <- newTest
	})
	client.Subscribe("/prod/pt/ssm/#", 0, func(client mqtt.Client, msg mqtt.Message) {
		ssmMsg := openprio_ssm.ExtendedSSM{}
		proto.Unmarshal(msg.Payload(), &ssmMsg)

		fmt.Printf("* [%s] %v\n", msg.Topic(), ssmMsg)
		ssmCh <- ssmMsg
	})

	msgs := []openprio_pt_position_data.LocationMessage{}
	//last_bulk_import := time.time()

	go saveSsmLog(ssmCh)

	for {
		data := <-positionCh
		msgs = append(msgs, data)
		if len(msgs) > 50 {
			log.Println("Save data.")
			saveData(msgs)
			msgs = []openprio_pt_position_data.LocationMessage{}
		}

	}
}

func saveSsmLog(ssmCh chan openprio_ssm.ExtendedSSM) {
	es := getElasticClient()
	indexName := "openprio_ssm"

	_, err := es.Indices.Exists([]string{indexName})
	if err != nil {
		log.Println(err)
	} else {
		_, err = es.Indices.Create(indexName)
		if err != nil {
			log.Fatalf("Cannot create index: %s", err)
		}
	}
	log.Println("Completed creating index.")

	for {
		msg := <-ssmCh
		type ElasticContainerLocation struct {
			Time    time.Time                `json:"time"`
			Content openprio_ssm.ExtendedSSM `json:"content"`
		}
		res := ElasticContainerLocation{Time: time.Now(), Content: msg}
		data, err := json.Marshal(res)
		if err != nil {
			log.Println("Error while serializing ssm.")
		}
		_, err = es.Index(indexName, bytes.NewReader(data))
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}
	}

}

func saveData(data []openprio_pt_position_data.LocationMessage) {
	es := getElasticClient()
	indexName := "openprio_pt"

	var (
		buf bytes.Buffer
		res *esapi.Response
	)
	_, err := es.Indices.Exists([]string{"openprio_pt"})
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
		type ElasticContainerLocation struct {
			Time     time.Time                                 `json:"time"`
			Location string                                    `json:"location_es"`
			Content  openprio_pt_position_data.LocationMessage `json:"content"`
		}

		test := ElasticContainerLocation{}
		test.Content = content
		test.Location = fmt.Sprintf("%f", content.GetPosition().Latitude) + "," + fmt.Sprintf("%f", content.GetPosition().Longitude)
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

func getElasticClient() *elasticsearch.Client {
	elasticAddress := os.Getenv("ELASTIC_ADDRESS")
	cfg := elasticsearch.Config{
		Addresses: []string{
			elasticAddress,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return es
}
