package web

//Aqui são executados os testes unitários relativo as funções que acessam o banco de dados
import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"log"

	"sw-sys/api-service/database"
	// "sw-sys/api-service/messaging"
)

func initHandler() *Handler {
	handler := &Handler{}

	handler.Database = initDataBase()
	// DB, errDB := connectDB()
	// if errDB != nil {
	// 	log.Println(errDB)
	// }
	// handler.DB = DB
	// handler.Authorization = initAuthorization()
	// redis, errRedis := connectChache()
	// if errRedis != nil {
	// 	log.Println(errRedis)
	// }
	// handler.Cache = redis
	// iAuth := initAuthJWT()
	// handler.Auth = iAuth
	return handler
}

func initDataBase() *database.IDataBase {
	DBOpt := &database.OptionsDBClient{
		URL:    "mongodb://localhost:27017",
		DBName: "StarWars",
		Driver: "mongoDB",
	}

	DBCliente, err := DBOpt.ConfiguraDatabase()
	if err != nil {
		log.Println("ERRO AO CONECTAR: ", err)
	}
	return DBCliente

}

// func TestCadastrarPlaneta(t *testing.T) {
// 	h := initHandler()
// 	planeta := &Planet{
// 		Nome:      "Alderan",
// 		Clima:     "temperado",
// 		Terreno:   "montanhoso",
// 		Aparicoes: 2,
// 	}

// 	err := Cadastrar(h, planeta)
// 	log.Println(planeta.ID)
// 	if err != nil {
// 		t.Error("Expect nil, got:", err)
// 	}

// }

// func TestGetPlaneta(t *testing.T) {
// 	h := initHandler()

// 	key := "608d8de4d753dcd9a43ea2b4"

// 	planeta, err := Get(h, key)
// 	log.Println(planeta.ID)
// 	log.Println(planeta.Nome)
// 	if err != nil {
// 		t.Error("Expect nil, got:", err)
// 	}

// }

// func TestListarPlanetas(t *testing.T) {
// 	h := initHandler()

// 	planetas, err := Listar(h)
// 	log.Println("Quantidade de planetas: ", len(planetas))
// 	if err != nil {
// 		t.Error("Expect nil, got:", err)
// 	}

// }

// func TestConsultarPlanetas(t *testing.T) {
// 	h := initHandler()
// 	nome := "Alderan"
// 	planetas, err := Consultar(h, nome)
// 	log.Println("Quantidade de planetas: ", len(planetas))
// 	if err != nil {
// 		t.Error("Expect nil, got:", err)
// 	}

// }

// func TestDeletePlaneta(t *testing.T) {
// 	h := initHandler()

// 	key := "608d8de4d753dcd9a43ea2b4"

// 	err := Delete(h, key)
// 	if err != nil {
// 		t.Error("Expect nil, got:", err)
// 	}

// }
