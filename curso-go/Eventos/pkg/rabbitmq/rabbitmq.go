package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") //user:senha@enderecorabbitmq:porta/namespace
	if err != nil {
		panic(err) //panic pois sem a conexao, o pgm nao devera rodar.
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	return ch, nil
}

func Consume(ch *amqp.Channel, outmsg chan<- amqp.Delivery) error {
	msgs, err := ch.Consume( //faz o pooling do rabbitMQ
		"minhafila",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		outmsg <- msg //pega cada mensagem na fila do RabbitMQ e joga para fora, no canal p/ ser lido em outra thread
	}
	return nil
}
