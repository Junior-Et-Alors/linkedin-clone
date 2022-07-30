package mongo

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Initializer() *options.ClientOptions {

	//client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	client := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	//errInitilizer(client, err)
	return client

}

// func errInitilizer(client *mongo.Client, err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(ctx)

// 	/*
// 	   List databases
// 	*/
// 	databases, err := client.ListDatabaseNames(ctx, bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(databases)
// }
