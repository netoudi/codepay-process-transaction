package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/mattn/go-sqlite3"
	"github.com/netoudi/codepay-process-transaction/adapter/broker/kafka"
	"github.com/netoudi/codepay-process-transaction/adapter/factory"
	"github.com/netoudi/codepay-process-transaction/adapter/presenter/transaction"
	"github.com/netoudi/codepay-process-transaction/usecase/process_transaction"
	"log"
)

func main() {
	fmt.Println("Hello World!")

	// db
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	// repository
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFactory.CreateTransactionRepository()

	// configMapProducer
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
	}

	kafkaPresenter := transaction.NewTransactionKafkaPresenter()

	// producer
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)

	var mshChan = make(chan *ckafka.Message)

	// configMapConsumer
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "goapp",
	}

	// topic
	topics := []string{"transactions"}

	// consumer
	consumer := kafka.NewKafkaConsumer(configMapConsumer, topics)
	go consumer.Consume(mshChan)

	// usecase
	usecase := process_transaction.NewProcessTransaction(repository, producer, "transactions_result")

	for msg := range mshChan {
		var input process_transaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}
}
