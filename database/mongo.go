package database

import (
	"context"
	"errors"
	"log"

	// "sw-sys/api-service/web"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	DBName string
}

func (m *MongoDB) connectService(config *OptionsDBClient) error {

	clientOptions := options.Client().ApplyURI(config.URL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	m.Client = client
	m.DBName = config.DBName
	log.Println("NOME DB: ", m.DBName)
	return nil
}

func (m *MongoDB) Get(value interface{}, key string, collection string) error {
	objectID, _ := primitive.ObjectIDFromHex(key)
	filter := bson.M{"_id": objectID}
	collectionOBJ := m.Client.Database(m.DBName).Collection(collection)
	err := collectionOBJ.FindOne(context.TODO(), filter).Decode(value)
	if err != nil {
		return err
	}
	return nil

}

func (m *MongoDB) Cadastrar(value interface{}, collection string) (interface{}, error) {

	collectionOBJ := m.Client.Database(m.DBName).Collection(collection)
	insRes, err := collectionOBJ.InsertOne(context.TODO(), value)
	if err != nil {
		return nil, err
	}
	log.Println("Inserido: ", insRes.InsertedID)
	// value = insRes
	return insRes.InsertedID, nil

}

// func (m *MongoDB) GerarID() ([]byte, error) {

// 	return primitive.NewObjectID(), nil

// }
func (m *MongoDB) Listar(collection string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	collectionOBJ := m.Client.Database(m.DBName).Collection(collection)
	cur, err := collectionOBJ.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		obj := make(map[string]interface{})
		// create a value into which the single document can be decoded
		var elem bson.D
		// elem := value
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		hash := elem.Map()

		for k, v := range hash {
			obj[k] = v
		}
		results = append(results, obj)
	}
	return results, nil
}
func (m *MongoDB) Consultar(nome string, collection string) ([]map[string]interface{}, error) {
	filter := bson.D{{"nome", nome}}
	var results []map[string]interface{}
	collectionOBJ := m.Client.Database(m.DBName).Collection(collection)
	cur, err := collectionOBJ.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		obj := make(map[string]interface{})
		// create a value into which the single document can be decoded
		var elem bson.D
		// elem := value
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		hash := elem.Map()

		for k, v := range hash {
			obj[k] = v
		}
		results = append(results, obj)
	}
	return results, nil

}
func (m *MongoDB) Delete(key string, collection string) error {
	objectID, _ := primitive.ObjectIDFromHex(key)
	filter := bson.M{"_id": objectID}
	collectionOBJ := m.Client.Database(m.DBName).Collection(collection)
	deleteResult, err := collectionOBJ.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		return errors.New("Planeta não encontrado")
	}
	return nil

}
