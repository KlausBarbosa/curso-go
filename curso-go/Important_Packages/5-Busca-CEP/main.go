package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"-"`
	Gia         string `json:"-"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"-"`
}

func main() {
	http.HandleFunc("/", BuscaCEPHandler)
	//metodo faz o handle/attach entre  /endpoint e com o metodo BuscaCEPHandler
	http.ListenAndServe(":8080", nil)

}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	//Pegando qualquer CEP que esteja passando na query String
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, error := BuscaCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
	//	vai dar um newEncoder usando o w(writer/response) e o Encode vai pegar o resultado da var cep e jogar dentro de w(writer/result)
	//muito utilizado quando quer devolver logo o valor da resposta sem necessidade de jogar em variavel pra trabalhar algum processamento

	//Tambem eh posivel jogar para uma variavel para que seja realizado algum tipo de manipulacao/processamento vide exemplo abaixo
	//result, err := json.Marshal(cep)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//w.Write(result)
}

func BuscaCep(cep string) (*ViaCEP, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var c ViaCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}
	return &c, nil
}
