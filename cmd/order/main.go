package order

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/devfullcycle/go-intensivo-jul/Internal/infra/database"
	"github.com/devfullcycle/go-intensivo-jul/Internal/usecase"
	"github.com/devfullcycle/go-intensivo-jul/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.dbPrimeiroCodigoGO")
	if err != nil {
		panic(err)
	}
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgRabbitMqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitMqChannel) //fica escutando a fila //processo que trava - utiliza o Go pra criar uma thread específica pra ele
	rabbitmqWorker(msgRabbitMqChannel, uc)
	//foi mudado para o método rabbitmqWorker
	//input := usecase.OrderInput{
	//	ID:    "2",
	//	Price: 20.0,
	//	Tax:   2.0,
	//}
	//
	//output, err := uc.Execute(input)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(output)
} //

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting Worker")
	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}
		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println("Mensagem Processada e salva no banco com Sucesso:", output)
	}
}
