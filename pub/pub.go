package pub

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

var (
	clusterID = "test-cluster"
	clientID  = "test-publisher"
	channel   = "orders"
)

func PublishValid() {
	guid := uuid.New()

	connection, err := stan.Connect(clusterID, clientID)

	if err != nil {
		log.Fatal(err)
	}

	pubData := fmt.Sprintf(dataVal, strings.ReplaceAll(guid.String(), "-", ""))

	connection.Publish(channel, []byte(pubData))
}

func PublishInvalid() {
	connection, err := stan.Connect(clusterID, clientID)

	if err != nil {
		log.Fatal(err)
	}

	pubData := dataInval

	connection.Publish(channel, []byte(pubData))
}

var dataInval = "231ewsad"

var dataVal = `{
	"order_uid": "%s",
	"track_number": "WBILMTESTTRACK",
	"entry": "WBIL",
	"delivery": {
	  "name": "Test Testov",
	  "phone": "+9720000000",
	  "zip": "2639809",
	  "city": "Kiryat Mozkin",
	  "address": "Ploshad Mira 15",
	  "region": "Kraiot",
	  "email": "test@gmail.com"
	},
	"payment": {
	  "transaction": "b563feb7b2b84b6test",
	  "request_id": "123",
	  "currency": "USD",
	  "provider": "wbpay",
	  "amount": 1817,
	  "payment_dt": 1637907727,
	  "bank": "alpha",
	  "delivery_cost": 1500,
	  "goods_total": 317,
	  "custom_fee": 123
	},
	"items": [
	  {
		"chrt_id": 9934930,
		"track_number": "WBILMTESTTRACK",
		"price": 453,
		"rid": "ab4219087a764ae0btest",
		"name": "Mascaras",
		"sale": 30,
		"size": "0",
		"total_price": 317,
		"nm_id": 2389212,
		"brand": "Vivienne Sabo",
		"status": 202
	  }
	],
	"locale": "en",
	"internal_signature": "",
	"customer_id": "test",
	"delivery_service": "meest",
	"shardkey": "9",
	"sm_id": 99,
	"date_created": "2021-11-26T06:22:19Z",
	"oof_shard": "1"
  }`
