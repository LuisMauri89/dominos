package advent

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	consume(ctx)
}
func consume(ctx context.Context) {

	r := kafka.NewReader(kafka.ReaderConfig{
		//Brokers: []string{broker1Address, broker2Address, broker3Address}, DATOS A RECIBIR
		Topic:   topic,
		GroupID: "my-group",
	})
	for {

		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("No Encontrado " + err.Error())
		}

		fmt.Println("Recibido: ", string(msg.Value))
	}
}
