package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN global variable*/
var MongoCN = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://poluxuser:pQcGVpSgoRUAJm3E@cluster0.uceea.mongodb.net/twittor?retryWrites=true&w=majority")

/*conectarBD function allo connect db */
func conectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Println("Conexion exitosa")
	return client
}

/*ChequeoConnection function for check */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
