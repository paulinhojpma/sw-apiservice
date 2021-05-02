package web

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	// "sw-sys/api-service/database"
)

// Planet monta a estrututra do plano
type Planet struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nome      string             `json:"nome"`
	Clima     string             `json:"clima"`
	Terreno   string             `json:"terreno"`
	Aparicoes int                `json:"aparicoes"`
}

const (
	PLANET = "planet"
)

// CadastraPlaneta ...
func Cadastrar(h *Handler, p *Planet) error {
	db := *h.Database
	// p.ID, _ = db.GerarID()
	idInte, err := db.Cadastrar(p, PLANET)
	if err != nil {
		return err
	}
	p.ID = idInte.(primitive.ObjectID)

	return nil
}

// ListaPlanetas ...
func Listar(h *Handler) ([]Planet, error) {
	db := *h.Database

	results, err := db.Listar(PLANET)
	if err != nil {
		return nil, err
	}
	planets := make([]Planet, len(results))
	for i, v := range results {
		planet := Planet{
			ID:        v["_id"].(primitive.ObjectID),
			Nome:      v["nome"].(string),
			Clima:     v["clima"].(string),
			Terreno:   v["terreno"].(string),
			Aparicoes: int(v["aparicoes"].(int32)),
		}

		planets[i] = planet

	}
	return planets, nil
}

// Consultar consultar a partir de um parâmetro ...
func Consultar(h *Handler, nome string) ([]Planet, error) {
	db := *h.Database
	results, err := db.Consultar(nome, PLANET)
	if err != nil {
		return nil, err
	}
	planets := make([]Planet, len(results))
	for i, v := range results {
		planet := Planet{
			ID:        v["_id"].(primitive.ObjectID),
			Nome:      v["nome"].(string),
			Clima:     v["clima"].(string),
			Terreno:   v["terreno"].(string),
			Aparicoes: int(v["aparicoes"].(int32)),
		}
		log.Println(i, "- id: ", planet.ID)
		planets[i] = planet
	}
	return planets, nil
}

// Get pega um planeta por ID ...
func Get(h *Handler, ID string) (*Planet, error) {
	db := *h.Database
	planet := &Planet{}
	err := db.Get(planet, ID, PLANET)
	if err != nil {
		return nil, err
	}

	return planet, nil
}

func Delete(h *Handler, ID string) error {
	db := *h.Database
	err := db.Delete(ID, PLANET)
	if err != nil {
		return err
	}
	return nil
}
