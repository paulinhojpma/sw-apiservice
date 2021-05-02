package database

import (
	"log"
	// "time"
)

// IDataBase ..
type IDataBase interface {
	connectService(config *OptionsDBClient) error
	Get(value interface{}, key string, collection string) error
	Cadastrar(value interface{}, collection string) (interface{}, error)
	Listar(collection string) ([]map[string]interface{}, error)
	Consultar(nome string, collection string) ([]map[string]interface{}, error)
	Delete(key string, collection string) error
	// GerarID() ([]byte, error)
}

// OptionsCacheClient ..
type OptionsDBClient struct {
	URL    string `json:"url"`
	DBName string `json:"DBName"`
	Driver string
}

// ConfiguraCache
func (o *OptionsDBClient) ConfiguraDatabase() (*IDataBase, error) {
	log.Println("Entrou no configura dataBase")
	var client IDataBase
	switch o.Driver {
	case "mongoDB":
		log.Println("Configura MongoDB")
		mongoDB := &MongoDB{}
		errMongo := mongoDB.connectService(o)
		if errMongo != nil {
			return nil, errMongo
		}
		client = mongoDB
	}
	return &client, nil
}
