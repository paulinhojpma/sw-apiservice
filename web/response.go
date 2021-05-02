package web

//DefaultResponse ...
type DefaultResponse struct {
	IDTipoTrans int    `json:"idTipoTrans"`
	IDTransBD   int    `json:"idTransBD,omitempty"`
	Mensagem    string `json:"mensagem,omitempty"`
}

//ResponseBodyJSONDefault ...
type ResponseBodyJSONDefault struct {
	// IDOperation string `json:"idOperation,omitempty"`
	CodResposta int    `json:"codResposta"`
	Mensagem    string `json:"mensagem"`
}

type PlanetResponse struct {
	Planet *Planet `json:"planet"`
	ResponseBodyJSONDefault
}

type PlanetsResponse struct {
	Planets []Planet `json:"planets"`
	ResponseBodyJSONDefault
}

type ErroResponse struct {
	Erro string `json:"erro"`
	ResponseBodyJSONDefault
}
