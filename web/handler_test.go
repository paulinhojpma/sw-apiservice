package web

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"bytes"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"testing"
)

func TestCadastraPlanetaServer(t *testing.T) {
	usuS := `{
		"nome": "Tatooine",
		"clima": "arido",
		"terreno": "deserto"
	}`
	// err := json.Unmarshal([]byte(usuS), planet)
	// if err != nil {
	// 	panic(err)

	// }
	r := httptest.NewRequest("POST", "/planetas", bytes.NewReader([]byte(usuS)))
	r.Header.Set("Content-Type", "application/json")
	// r.Header.Set("password", "vicente")

	w := httptest.NewRecorder()

	hWeb := initHandler()
	RouterTest(hWeb).ServeHTTP(w, r)
	bodyResponse, errBody := ioutil.ReadAll(w.Body)
	if errBody != nil {
		t.Error("expect nil, got - ", errBody)
	}
	planetaResponse := &PlanetResponse{}
	errJSON := json.Unmarshal(bodyResponse, planetaResponse)
	if errJSON != nil {
		t.Error("expect nil, got - ", errJSON)
	}
	log.Println("resposta do servidor - ", string(bodyResponse))
	usuS = `{
		"nome": "Alderaan",
		"clima": "temperado",
		"terreno": "montanhoso"
	}`

	r = httptest.NewRequest("POST", "/planetas", bytes.NewReader([]byte(usuS)))
	r.Header.Set("Content-Type", "application/json")
	// r.Header.Set("password", "vicente")

	w = httptest.NewRecorder()

	RouterTest(hWeb).ServeHTTP(w, r)
	bodyResponse, errBody = ioutil.ReadAll(w.Body)
	if errBody != nil {
		t.Error("expect nil, got - ", errBody)
	}
	planetaResponse = &PlanetResponse{}
	errJSON = json.Unmarshal(bodyResponse, planetaResponse)
	if errJSON != nil {
		t.Error("expect nil, got - ", errJSON)
	}
	log.Println("resposta do servidor - ", string(bodyResponse))

	usuS = `{
		"nome": "Hoth",
		"clima": "congelado",
		"terreno": "tundra"
	}`

	r = httptest.NewRequest("POST", "/planetas", bytes.NewReader([]byte(usuS)))
	r.Header.Set("Content-Type", "application/json")
	// r.Header.Set("password", "vicente")

	w = httptest.NewRecorder()

	RouterTest(hWeb).ServeHTTP(w, r)
	bodyResponse, errBody = ioutil.ReadAll(w.Body)
	if errBody != nil {
		t.Error("expect nil, got - ", errBody)
	}
	planetaResponse = &PlanetResponse{}
	errJSON = json.Unmarshal(bodyResponse, planetaResponse)
	if errJSON != nil {
		t.Error("expect nil, got - ", errJSON)
	}
	log.Println("resposta do servidor - ", string(bodyResponse))

}

func TestGetPlanetaServer(t *testing.T) {
	r := httptest.NewRequest("GET", "/planetas", nil)
	r.Header.Set("Content-Type", "application/json")
	// r.Header.Set("password", "vicente")

	w := httptest.NewRecorder()

	hWeb := initHandler()
	RouterTest(hWeb).ServeHTTP(w, r)
	bodyResponse, errBody := ioutil.ReadAll(w.Body)
	if errBody != nil {
		t.Error("expect nil, got - ", errBody)
	}
	planetasResponse := &PlanetsResponse{}
	errJSON := json.Unmarshal(bodyResponse, planetasResponse)
	ID := planetasResponse.Planets[0].ID.Hex()
	r = httptest.NewRequest("GET", fmt.Sprintf("/planetas/%s", ID), nil)
	r.Header.Set("Content-Type", "application/json")
	// r.Header.Set("password", "vicente")

	w = httptest.NewRecorder()

	RouterTest(hWeb).ServeHTTP(w, r)
	bodyResponse, errBody = ioutil.ReadAll(w.Body)
	if errBody != nil {
		t.Error("expect nil, got - ", errBody)
	}
	planetaResponse := &PlanetResponse{}
	errJSON = json.Unmarshal(bodyResponse, planetaResponse)
	if errJSON != nil {
		t.Error("expect nil, got - ", errJSON)
	}
	log.Println("resposta do servidor - ", string(bodyResponse))

}

func TestListarPlanetasServer(t *testing.T) {

	r := httptest.NewRequest("GET", "/planetas", nil)
	r.Header.Set("Content-Type", "application/json")
	// r.Header.Set("password", "vicente")

	w := httptest.NewRecorder()

	hWeb := initHandler()
	RouterTest(hWeb).ServeHTTP(w, r)
	bodyResponse, errBody := ioutil.ReadAll(w.Body)
	if errBody != nil {
		t.Error("expect nil, got - ", errBody)
	}
	planetasResponse := &PlanetsResponse{}
	errJSON := json.Unmarshal(bodyResponse, planetasResponse)
	if errJSON != nil {
		t.Error("expect nil, got - ", errJSON)
	}
	log.Println("resposta do servidor - ", string(bodyResponse))

}

func TestConsultarPlanetasServer(t *testing.T) {
	nome := "Tatooine"
	r := httptest.NewRequest("GET", fmt.Sprintf("/planetas?nome=%s", nome), nil)
	r.Header.Set("Content-Type", "application/json")
	// r.Header.Set("password", "vicente")

	w := httptest.NewRecorder()

	hWeb := initHandler()
	RouterTest(hWeb).ServeHTTP(w, r)
	bodyResponse, errBody := ioutil.ReadAll(w.Body)
	if errBody != nil {
		t.Error("expect nil, got - ", errBody)
	}
	planetasResponse := &PlanetsResponse{}
	errJSON := json.Unmarshal(bodyResponse, planetasResponse)
	if errJSON != nil {
		t.Error("expect nil, got - ", errJSON)
	}
	log.Println("resposta do servidor - ", string(bodyResponse))

}

func TestDeleteServer(t *testing.T) {
	r := httptest.NewRequest("GET", "/planetas", nil)
	r.Header.Set("Content-Type", "application/json")
	// r.Header.Set("password", "vicente")

	w := httptest.NewRecorder()

	hWeb := initHandler()
	RouterTest(hWeb).ServeHTTP(w, r)
	bodyResponse, errBody := ioutil.ReadAll(w.Body)
	if errBody != nil {
		t.Error("expect nil, got - ", errBody)
	}
	planetasResponse := &PlanetsResponse{}
	errJSON := json.Unmarshal(bodyResponse, planetasResponse)
	ID := planetasResponse.Planets[0].ID.Hex()
	r = httptest.NewRequest("DELETE", fmt.Sprintf("/planetas/%s", ID), nil)
	r.Header.Set("Content-Type", "application/json")
	// r.Header.Set("password", "vicente")

	w = httptest.NewRecorder()

	hWeb = initHandler()
	RouterTest(hWeb).ServeHTTP(w, r)
	bodyResponse, errBody = ioutil.ReadAll(w.Body)
	if errBody != nil {
		t.Error("expect nil, got - ", errBody)
	}
	planetaResponse := &PlanetResponse{}
	errJSON = json.Unmarshal(bodyResponse, planetaResponse)
	if errJSON != nil {
		t.Error("expect nil, got - ", errJSON)
	}
	log.Println("resposta do servidor - ", string(bodyResponse))

}
