package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DataItem struct {
	ID   int    `json:"player_id"`
	Nome string `json:"player_name"`
}

type Resposta struct {
	Data []DataItem `json:"data"`
}

func main(){
	res, _ := http.Get("https://technicalowl.herokuapp.com/playerhustlestats?seasontype=regular&season=2023")
	var resposta Resposta
	err := json.NewDecoder(res.Body).Decode(&resposta)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}
	if len(resposta.Data) > 0 {
		primeiroItem := resposta.Data[0]
		fmt.Println("ID:", primeiroItem.ID)
		fmt.Println("Nome:", primeiroItem.Nome)
	} else {
		fmt.Println("Lista de dados está vazia")
	}
}
