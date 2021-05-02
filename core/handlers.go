package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"

	mux "github.com/gorilla/mux"
	// "github.com/pkg/errors"
)

// RequestTest ..
func RequestTest(router *mux.Router, metodo, url string, bodyObj, responseObj interface{}) (int, error) {
	bodyRequestJSON, erroReq := requestEncodeJSON(bodyObj)
	if erroReq != nil {
		return 0, erroReq
	}

	r, err := http.NewRequest(metodo, url, bodyRequestJSON)
	if err != nil {
		return 0, err
	}
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	erroBody := responseDecodeJSON(w.Body, responseObj)
	if erroBody != nil {
		return w.Code, erroReq
	}

	return w.Code, nil
}

// responseDecodeJSON .. usado RequestTest
func responseDecodeJSON(bodyResponse io.Reader, response interface{}) error {
	var body, errBody = ioutil.ReadAll(bodyResponse)
	if errBody != nil {
		return errBody
	}

	errJSON := json.Unmarshal(body, response)
	if errJSON != nil {
		return errJSON
	}

	return nil
}

// requestEncodeJSON .. usado RequestTest
func requestEncodeJSON(objRequest interface{}) (*bytes.Buffer, error) {
	bodyRequestJSON := new(bytes.Buffer)
	encodeJSON, erro := json.Marshal(objRequest)
	if erro != nil {
		return nil, erro
	}
	bodyRequestJSON.Write(encodeJSON)

	return bodyRequestJSON, nil
}

//DecodeBodyJSON ...
func DecodeBodyJSON(r *http.Request, v interface{}) error {
	conteudo, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		return errors.New(ErrorReadAllBuffer)
	}
	// logger.Trace(fmt.Sprintf("Request: %s", string(conteudo)))

	if erro = json.Unmarshal(conteudo, v); erro != nil {
		return errors.New(ErrorJSONUnmarshal)
	}

	return nil
}

//Respond ...
func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		log.Println("Erro encode:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if _, err := io.Copy(w, &buf); err != nil {
		log.Println("Erro copy buf:", err)
	}

	log.Println(r.URL, "status:", status)
}

//RespondErro ...
func RespondErro(logSnet Logger, w http.ResponseWriter, r *http.Request, status int, errMsg *ErrMessage) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(errMsg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

//handleZen ...
func handleZen(w http.ResponseWriter, r *http.Request) {
	data := SuccessMessage{
		Message: "Keep it logically awesome.",
	}
	Respond(w, r, http.StatusOK, data)
	return
}

//handleVersion ...
func handleVersion(w http.ResponseWriter, r *http.Request) {
	data := VersionMessage{
		AppID:          os.Getenv("HEROKU_APP_ID"),
		AppName:        os.Getenv("HEROKU_APP_NAME"),
		ServerID:       os.Getenv("HEROKU_DYNO_ID"),
		CreatedAt:      os.Getenv("HEROKU_RELEASE_CREATED_AT"),
		ReleaseVersion: os.Getenv("HEROKU_RELEASE_VERSION"),
		Commit:         os.Getenv("HEROKU_SLUG_COMMIT"),
		Description:    os.Getenv("HEROKU_SLUG_DESCRIPTION"),
	}
	Respond(w, r, http.StatusOK, data)
	return
}

//handleNotFound ...
func handleNotFound(w http.ResponseWriter, r *http.Request) {

	body := ErrMessage{Message: "URL não encontrada",
		Code: strconv.Itoa(http.StatusNotFound)}
	Respond(w, r, http.StatusNotFound, body)
	return
}
